BUILD_DIR=bin
GO    := CGO_ENABLED=0 GOOS=linux GOPROXY=off go
PKG    = git.balhau.net/monitor
pkgs   = $(shell $(GO) list $(PKG)/... | grep -v /vendor/)

all: format build

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

build:
	@echo ">> Creating build dir"
	@mkdir -p $(BUILD_DIR)
	@echo ">> building binaries"
	@$(GO) vet $(PKG)/...
	@$(GO) build -o $(BUILD_DIR) -mod=vendor $(PKG)/cmd/hello

clean:
	@echo ">> cleaning project"
	@rm -rf $(BUILD_DIR)

.PHONY: all format build