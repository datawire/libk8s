generated += go.mod
generated += go.sum
generated += pin.go
generated += cmd/fix-go.mod/versions.go

generate: $(addprefix $(CURDIR)/,$(generated))
.PHONY: generate

$(addprefix %/,$(generated)): %/generate-libk8s.go
	go run $< refs/heads/release-1.14

clobber:
	rm -f $(generated)
.PHONY: clobber
