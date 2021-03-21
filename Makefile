BUILD_DIR=bin
GO    := CGO_ENABLED=0 GOOS=linux GOPROXY=off go
GOARM := GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 GOOS=linux GOPROXY=off go
PKG    = git.balhau.net/monitor
pkgs   = $(shell $(GO) list $(PKG)/... | grep -v /vendor/)

cyan = "\e[36m"
green = "\e[32m"
red ="\e[31m"


all: format build

arm: format buildarm

vendor:
	@echo $(green)
	@echo ">> Updating vendor"
	@$(GO) mod vendor

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

buildarm:
	@echo $(green)
	@echo ">> Creating build dir"
	@mkdir -p $(BUILD_DIR)
	@echo ">> building binaries"
	@$(GOARM) vet $(PKG)/...
	@$(GOARM) build -o $(BUILD_DIR) -mod=vendor $(PKG)/cmd/hello
	@$(GOARM) build -o $(BUILD_DIR) -mod=vendor $(PKG)/cmd/dnsspy


clean:
	@echo $(red)
	@echo ">> cleaning project"
	@rm -rf $(BUILD_DIR)

.PHONY: all format build vendor