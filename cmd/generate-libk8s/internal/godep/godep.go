package godep

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/datawire/libk8s/cmd/generate-libk8s/internal/gomod"
)

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

func PkgToMod(pkg string) string {
	fmt.Fprintf(os.Stderr, "Guessing module name for package %q...\n", pkg)

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

func ParseDir(dirname string) (gomod.GoMod, error) {
	fmt.Fprintf(os.Stderr, "Parsing Godeps for %q...\n", dirname)

	file, err := os.Open(filepath.Join(dirname, "Godeps", "Godeps.json"))
	if err != nil {
		return gomod.GoMod{}, err
	}
	defer file.Close()

	godepBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return gomod.GoMod{}, fmt.Errorf("read %q: %v", file.Name(), err)
	}

	var godepStruct Godep
	if err := json.Unmarshal(godepBytes, &godepStruct); err != nil {
		return gomod.GoMod{}, fmt.Errorf("parse %q: %v", file.Name(), err)
	}

	var goModStruct gomod.GoMod
	goModStruct.Module.Path = godepStruct.ImportPath
	if godepStruct.GoVersion != "unknown" {
		goModStruct.Go = strings.TrimPrefix(godepStruct.GoVersion, "go")
	}
	for _, dep := range godepStruct.Deps {
		goModStruct.Require = append(goModStruct.Require, gomod.Require{
			Path:    PkgToMod(dep.ImportPath),
			Version: dep.Rev,
		})
	}

	return goModStruct, nil
}
