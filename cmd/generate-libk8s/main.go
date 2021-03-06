package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/datawire/libk8s/cmd/generate-libk8s/internal/git"
	"github.com/datawire/libk8s/cmd/generate-libk8s/internal/godep"
	"github.com/datawire/libk8s/cmd/generate-libk8s/internal/golist"
	"github.com/datawire/libk8s/cmd/generate-libk8s/internal/gomod"
)

const StartRepo = "cli-runtime"

var PinnedPackages = []string{
	"k8s.io/cli-runtime/...",
	"k8s.io/client-go/...",
}

var usageTemplate = template.Must(template.
	New("--help").
	Parse(`Usage: {{ .Arg0 }} VERSION\n", arg0)
Overwrite go.mod, to use the specified version of k8s.io/` + StartRepo + `

VERSION must refer to a commit in the
https://github.com/kubernetes/` + StartRepo + `.git repository.  It may either
be a full 40-character commit hash, or a full ref name (starting with
"ref/").

Examples:
    {{ .Arg0 }} refs/heads/release-1.14                   # branch
    {{ .Arg0 }} refs/tags/kubernetes-1.14.6               # tag
    {{ .Arg0 }} 59d96aa1e208aaaaafb5d528dc5720e9e7b5c07e  # commit
`))

var headTemplate = template.Must(template.
	New("head.txt").
	Parse(`// Code generated by [{{ .Cmdline }}] (for k8s.io/` + StartRepo + ` commit {{ .Commit }}). DO NOT EDIT.

`))

var pinGoTemplate = template.Must(template.
	New("pin.go").
	Parse(`package libk8s

import (
{{- range $pkg := .Packages }}
	_ {{ printf "%q" $pkg }}
{{- end }}
)
`))

var versionsGoTemplate = template.Must(template.
	New("versions.go").
	Parse(`package main

var versions = []string{
{{- range $mod := .Versions }}
	{{ printf "%q" $mod }},
{{- end }}
}
`))

func mergeGoMod(a, b gomod.GoMod) (gomod.GoMod, error) {
	var ret gomod.GoMod
	ret.Module.Path = "merged"
	if a.Go == b.Go {
		ret.Go = a.Go
	} else if a.Go == "" {
		ret.Go = b.Go
	} else if b.Go == "" {
		ret.Go = a.Go
	} else if strings.HasPrefix(a.Go, "1.") && strings.HasPrefix(b.Go, "1.") {
		aMinor, err := strconv.ParseInt(strings.TrimPrefix(a.Go, "1."), 10, 0)
		if err != nil {
			return gomod.GoMod{}, fmt.Errorf("cannot merge `go.mod`s: invalid Go version: %q", a.Go)
		}
		bMinor, err := strconv.ParseInt(strings.TrimPrefix(b.Go, "1."), 10, 0)
		if err != nil {
			return gomod.GoMod{}, fmt.Errorf("cannot merge `go.mod`s: invalid Go version: %q", b.Go)
		}
		if aMinor > bMinor {
			ret.Go = a.Go
		} else {
			ret.Go = b.Go
		}
	} else {
		return gomod.GoMod{}, fmt.Errorf("cannot merge `go.mod`s: invalid Go versions: %q, %q", a.Go, b.Go)
	}

	ret.Require = append(ret.Require, a.Require...)
	ret.Require = append(ret.Require, b.Require...)
	ret.Exclude = append(ret.Exclude, a.Exclude...)
	ret.Exclude = append(ret.Exclude, b.Exclude...)

	var ret_Replace []gomod.Replace
	ret_Replace = append(ret_Replace, a.Replace...)
	ret_Replace = append(ret_Replace, b.Replace...)
	replaces := make(map[string]string)
	for _, replace := range ret_Replace {
		if val, set := replaces[replace.New.Path]; set {
			if val != replace.New.Version {
				return gomod.GoMod{}, fmt.Errorf("cannot merge `go.mod`s: conflicting replace %q: %q != %q", replace.New.Path, val, replace.New.Version)
			}
		} else {
			ret.Replace = append(ret.Replace, replace)
			replaces[replace.New.Path] = replace.New.Version
		}
	}

	return ret, nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func stage0(commit string) (gomod.GoMod, error) {
	var mergedGoMod gomod.GoMod
	loadedMods := make(map[string]struct{})
	var loadMod func(modDir string) error
	loadMod = func(modDir string) error {
		if _, set := loadedMods[modDir]; set {
			return nil
		}
		var goMod gomod.GoMod
		var err error
		switch {
		case exists(filepath.Join(modDir, "go.mod")):
			goMod, err = gomod.ParseDir(modDir)
		case exists(filepath.Join(modDir, "Godeps", "Godeps.json")):
			goMod, err = godep.ParseDir(modDir)
		default:
			err = fmt.Errorf("does not appear to use go.mod or Godep.json: %q", modDir)
		}
		if err != nil {
			return err
		}
		mergedGoMod, err = mergeGoMod(mergedGoMod, goMod)
		if err != nil {
			return err
		}
		loadedMods[modDir] = struct{}{}

		pkgs, err := golist.ListPackages(modDir, "-deps", "./...")
		if err != nil {
			return fmt.Errorf("go list -deps ./...: %q: %v", modDir, err)
		}
		for _, pkg := range pkgs {
			if pkg.Name == "main" {
				continue
			}
			if pkg.Module == nil {
				continue
			}
			if pkg.Module.Main {
				continue
			}
			if !strings.HasPrefix(pkg.ImportPath, "k8s.io/") {
				continue
			}

			// force the full sources to be extracted
			mod := pkg.Module
			if mod.Replace != nil {
				mod = mod.Replace
			}
			if _, err := gomod.GetDir(mod.Path, mod.Version); err != nil {
				return err
			}

			if pkg.Module.Path == "k8s.io/klog" || pkg.Module.Path == "k8s.io/utils" || pkg.Module.Path == "k8s.io/gengo" {
				if _, err := os.Stat(filepath.Join(pkg.Module.Dir, "go.mod")); err != nil {
					continue
				}
			}

			// recurse(+memoize)
			if err := loadMod(pkg.Module.Dir); err != nil {
				return err
			}
		}
		return nil
	}

	mergedGoMod = gomod.GoMod{
		Require: []gomod.Require{
			{Path: "k8s.io/" + StartRepo, Version: commit},
		},
	}
	startModDir, err := gomod.GetDir("k8s.io/"+StartRepo, commit)
	if err != nil {
		return gomod.GoMod{}, err
	}
	if err := loadMod(startModDir); err != nil {
		return gomod.GoMod{}, err
	}
	return mergedGoMod, nil
}

func stage1(mergedGoMod gomod.GoMod) ([]golist.Package, error) {
	tmpdir, err := ioutil.TempDir("", "tmp.generate-libk8s.stage1.")
	if err != nil {
		return nil, fmt.Errorf("tempdir: %v", err)
	}
	defer os.RemoveAll(tmpdir)

	goMod, err := os.OpenFile(filepath.Join(tmpdir, "go.mod"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	err = gomod.Write(goMod, mergedGoMod)
	goMod.Close()
	if err != nil {
		return nil, fmt.Errorf("write go.mod: %v", err)
	}

	pkgs, err := golist.ListPackages(tmpdir, append([]string{"-deps"}, PinnedPackages...)...)
	if err != nil {
		return nil, fmt.Errorf("go list: %v", err)
	}

	return pkgs, nil
}

func stage2(arg0 string, version string, commit string, mergedGoMod gomod.GoMod, allPkgs []golist.Package) error {
	// Do this in "." instead of os.TempDir ($TMPDIR or /tmp) so
	// that we don't have to worry about cross-filesystem moves
	// and can just use os.Rename.
	tmpdir, err := ioutil.TempDir(".", "tmp.generate-libk8s.stage2.")
	if err != nil {
		return fmt.Errorf("tempdir: %v", err)
	}
	defer os.RemoveAll(tmpdir)

	var pinPkgs []golist.Package
	for _, pkg := range allPkgs {
		if pkg.Standard {
			continue
		}
		if pkg.Name == "main" {
			continue
		}
		if len(pkg.GoFiles) == 0 {
			continue
		}
		if strings.Contains(pkg.ImportPath, "/internal/") || strings.HasSuffix(pkg.ImportPath, "/internal") {
			continue
		}
		pinPkgs = append(pinPkgs, pkg)
	}

	pinPkgNames := make([]string, 0, len(pinPkgs))
	pinModNames := make([]string, 0)
	pinModVersions := make(map[string]string)
	for _, pkg := range pinPkgs {
		mod := pkg.Module
		if mod.Replace != nil {
			mod = mod.Replace
		}
		pinPkgNames = append(pinPkgNames, pkg.ImportPath)
		if _, set := pinModVersions[mod.Path]; !set {
			pinModNames = append(pinModNames, mod.Path)
			pinModVersions[mod.Path] = mod.Version
		}
	}
	sort.Strings(pinPkgNames)
	sort.Strings(pinModNames)

	finalGoMod := gomod.GoMod{
		Module: gomod.Module{
			Path: "github.com/datawire/libk8s",
		},
		Go: mergedGoMod.Go,
	}
	for _, modName := range pinModNames {
		finalGoMod.Require = append(finalGoMod.Require, gomod.Require{
			Path:    modName,
			Version: pinModVersions[modName],
		})
	}

	// go.mod
	fileGoMod, err := os.OpenFile(filepath.Join(tmpdir, "go.mod"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	err = headTemplate.Execute(fileGoMod, map[string]interface{}{
		"Cmdline": fmt.Sprintf("%s %q", arg0, version),
		"Commit":  commit,
	})
	if err == nil {
		err = gomod.Write(fileGoMod, finalGoMod)
	}
	fileGoMod.Close()
	if err != nil {
		return fmt.Errorf("write go.mod: %v", err)
	}

	// pin.go
	filePinGo, err := os.OpenFile(filepath.Join(tmpdir, "pin.go"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	err = headTemplate.Execute(filePinGo, map[string]interface{}{
		"Cmdline": fmt.Sprintf("%s %q", arg0, version),
		"Commit":  commit,
	})
	if err == nil {
		err = pinGoTemplate.Execute(filePinGo, map[string]interface{}{
			"Packages": pinPkgNames,
		})
	}
	filePinGo.Close()
	if err != nil {
		return fmt.Errorf("write pin.go: %v", err)
	}

	// clean up
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = tmpdir
	cmd.Stderr = os.Stderr
	fmt.Fprintf(os.Stderr, "%s$ %s\n", cmd.Dir, strings.Join(cmd.Args, " "))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go mod tidy: %v", err)
	}

	// version.go (do this after `go mod tidy`)
	verBytes, err := golist.List(tmpdir, "-deps", `-f={{ if not .Standard }}{{ if not .Module.Main }}{{ .Module.Path }}@{{ .Module.Version }}{{ end }}{{ end}}`, ".")
	if err != nil {
		return fmt.Errorf("go list: %v", err)
	}
	versionSet := make(map[string]struct{})
	for _, line := range strings.Split(string(verBytes), "\n") {
		if line == "" {
			continue
		}
		versionSet[line] = struct{}{}
	}
	versions := make([]string, 0, len(versionSet))
	for ver := range versionSet {
		versions = append(versions, ver)
	}
	sort.Strings(versions)

	fileVersionsGo, err := os.OpenFile(filepath.Join(tmpdir, "versions.go"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	err = headTemplate.Execute(fileVersionsGo, map[string]interface{}{
		"Cmdline": fmt.Sprintf("%s %q", arg0, version),
		"Commit":  commit,
	})
	if err == nil {
		err = versionsGoTemplate.Execute(fileVersionsGo, map[string]interface{}{
			"Versions": versions,
		})
	}
	fileVersionsGo.Close()
	if err != nil {
		return fmt.Errorf("write versions.go: %v", err)
	}

	// move everything in to place
	if err := os.Rename(filepath.Join(tmpdir, "go.mod"), "go.mod"); err != nil {
		return err
	}
	if err := os.Rename(filepath.Join(tmpdir, "go.sum"), "go.sum"); err != nil {
		return err
	}
	if err := os.Rename(filepath.Join(tmpdir, "pin.go"), "pin.go"); err != nil {
		return err
	}
	if err := os.Rename(filepath.Join(tmpdir, "versions.go"), "cmd/fix-go.mod/versions.go"); err != nil {
		return err
	}

	return nil
}

func Main(arg0, version string) error {
	commit, err := git.ResolveCommit("https://github.com/kubernetes/"+StartRepo+".git", version)
	if err != nil {
		return err
	}
	mergedGoMod, err := stage0(commit)
	if err != nil {
		return err
	}
	pkgs, err := stage1(mergedGoMod)
	if err != nil {
		return err
	}
	return stage2(arg0, version, commit, mergedGoMod, pkgs)
}

func main() {
	arg0 := "go run ./cmd/generate-libk8s"
	if len(os.Args) != 2 {
		_ = usageTemplate.Execute(os.Stderr, map[string]interface{}{
			"Arg0": arg0,
		})
		os.Exit(2)
	}
	if err := Main(arg0, os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
