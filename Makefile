GO    := CGO_ENABLED=0 GOOS=linux GOPROXY=off go
PKG    = git.balhau.net/monitor
pkgs   = $(shell $(GO) list $(PKG)/... | grep -v /vendor/)

all: format build

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

build:
	@echo ">> building binaries"
	@$(GO) vet $(PKG)/...
	@$(GO) build -mod=vendor $(PKG)/cmd/hello

.PHONY: all format build