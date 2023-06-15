GOOS ?= linux
GOARCH ?= $(shell go env GOARCH)
GOFLAGS ?=""

.PHONY: build
build: build_server build_cli

.PHONY: build_server
build_server:
	@echo "Building assistant server..."
	@GOFLAGS=$(GOFLAGS) BUILD_PLATFORMS=$(GOOS)/$(GOARCH) go build -o bin/assistant ./cmd/server

.PHONY: build_cli
build_cli:
	@echo "Building assistant CLI..."
	@GOFLAGS=$(GOFLAGS) BUILD_PLATFORMS=$(GOOS)/$(GOARCH) go build -o bin/acli ./cmd/cli
