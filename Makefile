generated += go.mod
generated += go.sum
generated += pin.go
generated += cmd/fix-go.mod/versions.go

generate: $(addprefix $(CURDIR)/,$(generated))
.PHONY: generate

$(addprefix %/,$(generated)): %/cmd/generate-libk8s FORCE
	rm -f go.mod
	go mod init github.com/datawire/libk8s
	go run $< refs/heads/release-1.15

clobber:
	rm -f $(generated)
.PHONY: clobber

.PHONY: FORCE
