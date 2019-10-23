package gomod

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

type Module struct {
	Path    string
	Version string
}

type GoMod struct {
	Module  Module
	Go      string
	Require []Require
	Exclude []Module
	Replace []Replace
}

type Require struct {
	Path     string
	Version  string
	Indirect bool
}

type Replace struct {
	Old Module
	New Module
}

type cacheKey struct {
	name    string
	version string
}

var cache = map[cacheKey]string{}

func encode(in string) string {
	out := new(strings.Builder)
	for _, b := range []byte(in) {
		switch {
		case 'A' <= b && b <= 'Z':
			out.WriteByte(b - 'A' + 'a')
		case ('a' <= b && b <= 'z') || ('0' <= b && b <= '9'):
			out.WriteByte(b)
		default:
			fmt.Fprintf(out, "-%02x", b)
		}
	}
	return out.String()
}

func GetDir(modname, version string) (string, error) {
	if cacheValue, cached := cache[cacheKey{modname, version}]; cached {
		return cacheValue, nil
	}

	tmpdir, err := ioutil.TempDir("", "tmp.generate-libk8s.Getdir."+encode(modname)+"."+encode(version)+".")
	if err != nil {
		return "", fmt.Errorf("tempdir: %v", err)
	}
	defer os.RemoveAll(tmpdir)

	goMod, err := os.OpenFile(filepath.Join(tmpdir, "go.mod"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return "", err
	}
	// The k8s.io/klog is a special hack to make sure that we
	// don't get an old commit in the middle of a module rename.
	if _, err := fmt.Fprintln(goMod, `module tmp
require k8s.io/klog v0.1.0
require`, modname, version); err != nil {
		return "", fmt.Errorf("write go.mod: %v", err)
	}

	// force a resolve+download+extract
	cmd := exec.Command("go", "list", modname+"/...")
	cmd.Dir = tmpdir
	cmd.Stderr = os.Stderr
	fmt.Fprintf(os.Stderr, "%s$ %s\n", cmd.Dir, strings.Join(cmd.Args, " "))
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("go list %s@%s/...: %v", modname, version, err)
	}

	// get the directory that it was extracted to
	cmd = exec.Command("go", "list", "-m", "-f={{ .Dir }}", modname)
	cmd.Dir = tmpdir
	cmd.Stderr = os.Stderr
	fmt.Fprintf(os.Stderr, "%s$ %s\n", cmd.Dir, strings.Join(cmd.Args, " "))
	outputBytes, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("go list -m %s@%s: %v", modname, version, err)
	}

	ret := strings.TrimSpace(string(outputBytes))
	cache[cacheKey{modname, version}] = ret
	return ret, nil
}

func ParseDir(dirname string) (GoMod, error) {
	cmd := exec.Command("go", "mod", "edit", "-json")
	cmd.Dir = dirname
	cmd.Stderr = os.Stderr
	fmt.Fprintf(os.Stderr, "%s$ %s\n", cmd.Dir, strings.Join(cmd.Args, " "))
	goModBytes, err := cmd.Output()
	if err != nil {
		return GoMod{}, fmt.Errorf("go mod edit -json: %q: %v", dirname, err)
	}

	var goModStruct GoMod
	if err := json.Unmarshal(goModBytes, &goModStruct); err != nil {
		return GoMod{}, fmt.Errorf("parse json: %q: %v", dirname, err)
	}

	return goModStruct, nil
}

var tmpl = template.Must(template.
	New("go.mod").
	Parse(`module {{ .Module.Path }}

go {{ .Go }}{{ if gt (len .Require) 0 }}

require ({{ range $require := .Require }}
	{{ $require.Path }} {{ $require.Version }}{{ end }}
){{ end }}{{ if gt (len .Exclude) 0 }}

exclude ({{ range $exclude := .Exclude }}
	{{ $exclude.Path }} {{ $exclude.Version }}{{ end }}
){{ end }}{{ if gt (len .Replace) 0 }}

replace ({{ range $replace := .Replace }}
	{{ $replace.Old.Path }} {{ $replace.Old.Version }} => {{ $replace.New.Path }} {{ $replace.New.Version }}{{ end }}
){{ end }}
`))

func Write(w io.Writer, dat GoMod) error {
	return tmpl.Execute(w, dat)
}
