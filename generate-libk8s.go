// +build

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

const (
	StartRepo = "cli-runtime"
)

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

var goModTemplate = template.Must(template.
	New("go.mod").
	Funcs(template.FuncMap{
		"trimPrefix": strings.TrimPrefix,
		"pkg2mod":    pkg2mod,
	}).
	Parse(`// Code generated by [{{ .Cmdline }}] (for k8s.io/` + StartRepo + ` commit {{ .Commit }}). DO NOT EDIT.

module github.com/datawire/libk8s

go {{ trimPrefix (index .Godeps 0).GoVersion "go" }}

require (
	k8s.io/` + StartRepo + ` {{ .Commit }}
{{- range $manifest := .Godeps }}
{{- range $dep := $manifest.Deps }}
	{{ pkg2mod $dep.ImportPath }} {{ $dep.Rev }}
{{- end }}
{{- end }}
)
`))

var pinGoTemplate = template.Must(template.
	New("pin.go").
	Parse(`// Code generated by [{{ .Cmdline }}] (for k8s.io/` + StartRepo + ` commit {{ .Commit }}). DO NOT EDIT.

//go:generate {{ .Cmdline }}

package libk8s

import (
{{- range $pkg := .Packages }}
	_ {{ printf "%q" $pkg }}
{{- end }}
)
`))

var versionsGoTemplate = template.Must(template.
	New("versions.go").
	Parse(`// Code generated by [{{ .Cmdline }}] (for k8s.io/` + StartRepo + ` commit {{ .Commit }}). DO NOT EDIT.

package main

var versions = []string{
{{- range $mod := .Versions }}
	{{ printf "%q" $mod }},
{{- end }}
}
`))

type Godep struct {
	ImportPath   string
	GoVersion    string
	GodepVersion string
	Packages     []string
	Deps         []struct {
		ImportPath string
		Rev        string
	}
}

func getGodep(repo, commit string) (Godep, error) {
	url := "https://raw.githubusercontent.com/kubernetes/" + repo + "/" + commit + "/Godeps/Godeps.json"
	resp, err := http.Get(url)
	if err != nil {
		return Godep{}, fmt.Errorf("%v %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Godep{}, fmt.Errorf("%v HTTP %v", url, resp.StatusCode)
	}

	godepBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Godep{}, fmt.Errorf("%v %v", url, err)
	}

	var godepStruct Godep
	if err := json.Unmarshal(godepBytes, &godepStruct); err != nil {
		return Godep{}, fmt.Errorf("%v %v", url, err)
	}

	return godepStruct, nil
}

func pkg2mod(pkg string) string {
	pkgParts := strings.Split(pkg, "/")
	switch {
	case strings.HasPrefix(pkg, "cloud.google.com/go/"):
		return "cloud.google.com/go"
	case strings.HasPrefix(pkg, "github.com/") && len(pkgParts) > 3:
		return strings.Join(pkgParts[:3], "/")
	case strings.HasPrefix(pkg, "golang.org/x/") && len(pkgParts) > 3:
		return strings.Join(pkgParts[:3], "/")
	case strings.HasPrefix(pkg, "gopkg.in/") && len(pkgParts) > 2:
		return strings.Join(pkgParts[:2], "/")
	case strings.HasPrefix(pkg, "k8s.io/") && len(pkgParts) > 2:
		return strings.Join(pkgParts[:2], "/")
	case strings.HasPrefix(pkg, "sigs.k8s.io/") && len(pkgParts) > 2:
		return strings.Join(pkgParts[:2], "/")
	default:
		return pkg
	}
}

func resolveCommit(version string) (string, error) {
	switch {
	case len(version) == 40 && func() bool { _, err := hex.DecodeString(version); return err == nil }():
		return version, nil
	case strings.HasPrefix(version, "refs/"):
		url := "https://github.com/kubernetes/" + StartRepo + ".git/info/refs?service=git-upload-pack"
		resp, err := http.Get(url)
		if err != nil {
			return "", fmt.Errorf("%v %v", url, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("%v HTTP %v", url, resp.StatusCode)
		}

		refsBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("%v %v", url, err)
		}

		for _, line := range strings.Split(string(refsBytes), "\n") {
			if len(line) < 46 || line[44] != ' ' {
				continue
			}
			if line[45:] == version {
				return line[4:44], nil
			}
		}

		return "", fmt.Errorf("did not find a ref named %q in .git/info/refs", version)
	default:
		return "", fmt.Errorf("version %q doesn't look like either a commit hash or a ref name", version)
	}
}

func Main(arg0, version string) error {
	commit, err := resolveCommit(version)
	if err != nil {
		return err
	}

	godep, err := getGodep(StartRepo, commit)
	if err != nil {
		return err
	}

	godeps := []Godep{godep}
	for _, dep := range godep.Deps {
		if dep.ImportPath == "k8s.io/client-go" || strings.HasPrefix(dep.ImportPath, "k8s.io/client-go/") {
			cgGodep, err := getGodep("client-go", dep.Rev)
			if err != nil {
				return err
			}
			godeps = append(godeps, cgGodep)
		}
	}

	// Do this in "." instead of os.TempDir ($TMPDIR or /tmp) so
	// that we don't have to worry about cross-filesystem moves
	// and can just use os.Rename.
	tmpdir, err := ioutil.TempDir(".", "tmp."+filepath.Base(strings.TrimSuffix(arg0, ".go")))
	if err != nil {
		return fmt.Errorf("tempdir: %v", err)
	}
	defer os.RemoveAll(tmpdir)

	goMod, err := os.OpenFile(filepath.Join(tmpdir, "go.mod"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	err = goModTemplate.Execute(goMod, map[string]interface{}{
		"Cmdline": fmt.Sprintf("%s %q", arg0, version),
		"Commit":  commit,
		"Godeps":  godeps,
	})
	goMod.Close()
	if err != nil {
		return fmt.Errorf("write go.mod: %v", err)
	}

	cmd := exec.Command("go", "list", "-deps", `-f={{ if and (not .Standard) (ne .Name "main") (gt (len .GoFiles) 0) }}{{ .ImportPath }}{{ end }}`, "k8s.io/cli-runtime/...", "k8s.io/client-go/...")
	cmd.Dir = tmpdir
	cmd.Stderr = os.Stderr
	allPkgsBytes, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("go list: %v", err)
	}
	var allPkgs []string
	for _, line := range strings.Split(string(allPkgsBytes), "\n") {
		if line == "" {
			continue
		}
		if strings.Contains(line, "/internal/") || strings.HasSuffix(line, "/internal") {
			continue
		}
		allPkgs = append(allPkgs, line)
	}
	sort.Strings(allPkgs)

	pin, err := os.OpenFile(filepath.Join(tmpdir, "pin.go"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	err = pinGoTemplate.Execute(pin, map[string]interface{}{
		"Cmdline":  fmt.Sprintf("%s %q", arg0, version),
		"Commit":   commit,
		"Packages": allPkgs,
	})
	pin.Close()
	if err != nil {
		return fmt.Errorf("write pin.go: %v", err)
	}

	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = tmpdir
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go mod tidy: %v", err)
	}

	cmd = exec.Command("go", "list", "-deps", `-f={{ if not .Standard }}{{ if not .Module.Main }}{{ .Module.Path }}@{{ .Module.Version }}{{ end }}{{ end}}`, ".")
	cmd.Dir = tmpdir
	cmd.Stderr = os.Stderr
	verBytes, err := cmd.Output()
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

	versionsGo, err := os.OpenFile(filepath.Join(tmpdir, "versions.go"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	err = versionsGoTemplate.Execute(versionsGo, map[string]interface{}{
		"Cmdline":  fmt.Sprintf("%s %q", arg0, version),
		"Commit":   commit,
		"Versions": versions,
	})
	versionsGo.Close()
	if err != nil {
		return fmt.Errorf("write versions.go: %v", err)
	}

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

func main() {
	arg0 := "go run ./generate-libk8s.go"
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
