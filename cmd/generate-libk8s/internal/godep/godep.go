package godep

import (
	"strings"
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
