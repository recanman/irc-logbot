# Makefile

# Variables
GO_VERSION?= $(shell go version | sed 's/go version \([0-9]\+\.[0-9]\+\.[0-9]*\).*/\1/')
OS := linux
ARCHS := amd64 arm64
BUILD_DIR :=./build
BIN_NAME := ircbot

# Default target
all: build-all

.PHONY: build-clean build-all clean

# Build for all architectures
build-all: build-amd64 build-arm64

# Individual build targets for each architecture
build-amd64:
	@echo Building for amd64
	@mkdir -p $(BUILD_DIR)/amd64
	GOOS=$(OS) GOARCH=amd64 go build -o $(BUILD_DIR)/amd64/$(BIN_NAME) cmd/ircbot.go

build-arm64:
	@echo Building for arm64
	@mkdir -p $(BUILD_DIR)/arm64
	GOOS=$(OS) GOARCH=arm64 go build -o $(BUILD_DIR)/arm64/$(BIN_NAME) cmd/ircbot.go

build-ppc64le:
	@echo Building for ppc64le
	@mkdir -p $(BUILD_DIR)/ppc64le
	GOOS=$(OS) GOARCH=ppc64le go build -o $(BUILD_DIR)/ppc64le/$(BIN_NAME) cmd/ircbot.go

build-s390x:
	@echo Building for s390x
	@mkdir -p $(BUILD_DIR)/s390x
	GOOS=$(OS) GOARCH=s390x go build -o $(BUILD_DIR)/s390x/$(BIN_NAME) cmd/ircbot.go

# Clean up build artifacts
clean:
	rm -rf $(BUILD_DIR)
