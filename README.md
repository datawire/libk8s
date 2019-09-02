# libk8s - A wrapper around `k8s.io/client-go` [![CircleCI](https://circleci.com/gh/datawire/libk8s.svg)](https://circleci.com/gh/datawire/libk8s)

Right now it just makes dealing with the dependency list a little more
sane.  Maybe in the future [`github.com/datawire/teleproxy/pkg/k8s`][]
will migrate to here.

[`github.com/datawire/teleproxy/pkg/k8s`]: https://godoc.org/github.com/datawire/teleproxy/pkg/k8s

## Using libk8s

 1. Stick `_ "github.com/datawire/libk8s"` in the `imports ()`
    section of one of your `.go` files.
 2. Run `go run github.com/datawire/libk8s/cmd/fix-go.mod` to
    downgrade anything that `k8s.io/client-go` wants to use an older
    version of.
 3. Run `go mod tidy` again to bring everything back in to alignment.

## Hacking on libk8s.

Most of libk8s is generated from `k8s.io/client-go`'s Godep files by
the `go run ./generate-libk8s.go` script (which is called by `make`,
for convenience).  The only "actual" files to edit are:

 - `Makefile`
 - `generate-libk8s.go`
 - `cmd/fix-go.mod/main.go`
 - `README.md`
 - `.gitignore`

I advise that use Go 1.13 when running `make` (even though it's still
in RC2 at the time of this writing).  It does a better job of
resolving commit hashes to tags than 1.12.9 does.
