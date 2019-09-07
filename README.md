# libk8s - A wrapper around `k8s.io/client-go`

[![GoDoc](https://godoc.org/github.com/datawire/libk8s?status.svg)](https://godoc.org/github.com/datawire/libk8s)
[![Go Report Card](https://goreportcard.com/badge/github.com/datawire/libk8s)](https://goreportcard.com/report/github.com/datawire/libk8s)
[![CircleCI](https://circleci.com/gh/datawire/libk8s.svg)](https://circleci.com/gh/datawire/libk8s)

Right now libk8s just makes dealing with crazy dependency list of
`k8s.io/client-go` a little bit easier.  Maybe in the future
[`github.com/datawire/teleproxy/pkg/k8s`][] will migrate to here.

[`github.com/datawire/teleproxy/pkg/k8s`]: https://godoc.org/github.com/datawire/teleproxy/pkg/k8s

## Using libk8s

### Adding libk8s

 1. Stick `_ "github.com/datawire/libk8s"` in the `imports ()`
    section of one of your `.go` files.
 2. Run `go run github.com/datawire/libk8s/cmd/fix-go.mod` to
    downgrade anything that `k8s.io/client-go` wants to use an older
    version of.
 3. Run `go mod tidy` to bring everything back in to alignment.

### Upgrading/downgrading libk8s

 1. Run `go mod edit -require=github.com/datawire/libk8s@THE_VERSION_THAT_YOU_WANT`
 2. Run `go run github.com/datawire/libk8s/cmd/fix-go.mod`
 3. Run `go mod tidy`

## Hacking on libk8s

Most of libk8s is generated from `k8s.io/client-go`'s Godep files by
the `go run ./generate-libk8s.go` script (which is called by `make`,
for convenience).  The only "actual" files to edit are:

 - `Makefile`
 - `generate-libk8s.go`
 - `cmd/fix-go.mod/main.go`
 - `README.md`
 - `.gitignore`

I advise that you use Go 1.13 when running `make`.  It does a better
job of resolving commit hashes to tags than 1.12.9 does.
