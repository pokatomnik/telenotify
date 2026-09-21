APP_NAME := telenotify
CMD_DIR := ./cmd/telenotify
DIST_DIR := dist

GO ?= go
GOOS ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)
CGO_ENABLED ?= 0
VERSION ?= dev
LDFLAGS := -s -w -X main.version=$(VERSION)
GOFLAGS := -trimpath
GO_BUILD_FLAGS := $(GOFLAGS) -ldflags '$(LDFLAGS)'

TARGETS := \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64 \
	windows/amd64 \
	windows/arm64

.PHONY: all lint test generate build build-all install clean

all: lint test build

lint:
	$(GO) fmt ./...
	$(GO) vet ./...

test:
	$(GO) test ./...

generate:
	$(GO) generate ./...

build:
	@mkdir -p $(DIST_DIR)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(GO_BUILD_FLAGS) -o $(DIST_DIR)/$(APP_NAME) $(CMD_DIR)

build-all:
	@mkdir -p $(DIST_DIR)
	@set -e; \
	for target in $(TARGETS); do \
		os=$${target%/*}; \
		arch=$${target#*/}; \
		extension=; \
		if [ "$$os" = "windows" ]; then extension=.exe; fi; \
		output=$(DIST_DIR)/$(APP_NAME)-$$os-$$arch$$extension; \
		echo "Building $$output"; \
		CGO_ENABLED=$(CGO_ENABLED) GOOS=$$os GOARCH=$$arch $(GO) build $(GO_BUILD_FLAGS) -o $$output $(CMD_DIR); \
	done

install:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) install $(GO_BUILD_FLAGS) $(CMD_DIR)

clean:
	rm -rf $(DIST_DIR)
