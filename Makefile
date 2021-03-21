BUILD_DIR=bin
GO    := CGO_ENABLED=0 GOOS=linux GOPROXY=off go
PKG    = git.balhau.net/monitor
pkgs   = $(shell $(GO) list $(PKG)/... | grep -v /vendor/)

cyan = "\e[36m"
green = "\e[32m"
red ="\e[31m"


all: format build

format:
	@echo $(cyan)
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

build:
	@echo $(green)
	@echo ">> Creating build dir"
	@mkdir -p $(BUILD_DIR)
	@echo ">> building binaries"
	@$(GO) vet $(PKG)/...
	@$(GO) build -o $(BUILD_DIR) -mod=vendor $(PKG)/cmd/hello
	@$(GO) build -o $(BUILD_DIR) -mod=vendor $(PKG)/cmd/dnsspy

clean:
	@echo $(red)
	@echo ">> cleaning project"
	@rm -rf $(BUILD_DIR)

.PHONY: all format build