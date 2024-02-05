# The old school Makefile, following are required targets. The Makefile is written
# to allow building multiple binaries. You are free to add more targets or change
# existing implementations, as long as the semantics are preserved.
#
#   make              - default to 'build' target
#   make test         - run unit test
#   make build        - build local binary targets
#   make container    - build containers
#   make push         - push containers
#   make clean        - clean up targets
#
# The makefile is also responsible to populate project version information.

# Tweak the variables based on your project.

SHELL := /bin/bash
NOW_SHORT := $(shell date +%Y%m%d%H%M)

PROJECT := dingmark
# Target binaries. You can build multiple binaries for a single project.
TARGETS := dingmark

# Container registries.
REGISTRIES ?= ""

# Container image prefix and suffix added to targets.
# The final built images are:
#   $[REGISTRY]$[IMAGE_PREFIX]$[TARGET]$[IMAGE_SUFFIX]:$[VERSION]
# $[REGISTRY] is an item from $[REGISTRIES], $[TARGET] is an item from $[TARGETS].
IMAGE_PREFIX ?= $(strip )
IMAGE_SUFFIX ?= $(strip )

# This repo's root import path (under GOPATH).
ROOT := github.com/alswl/dingmark

# Project main package location (can be multiple ones).
CMD_DIR := ./cmd
CMD_WASM_DIR := ./cmdwasm

# Project output directory.
OUTPUT_DIR := ./bin

# Build directory.
BUILD_DIR := ./build

PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))

# Git commit sha.
COMMIT := $(strip $(shell git rev-parse --short HEAD 2>/dev/null))
COMMIT := $(COMMIT)$(shell [[ -z $$(git status -s) ]] || echo '-dirty')
COMMIT := $(if $(COMMIT),$(COMMIT), $${COMMIT})
COMMIT := $(if $(COMMIT),$(COMMIT),"Unknown")

DEFAULT_BUMP_STAGE := final # final, alpha, beta, candidate
DEFAULT_BUMP_SCOPE := minor # major, minor, patch
DEFAULT_BUMP_DRY_RUN := true # true, false

# Current version of the project.
VERSION_IN_FILE = $(shell cat VERSION)
BUILD_VERSION = $(COMMIT)
GO_MOD_VERSION = $(shell cat go.mod | sha256sum | cut -c-6)
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
VERSION ?= $(VERSION_IN_FILE)-$(BUILD_VERSION)

UT_COVER_PACKAGES := $(shell go list ./pkg/... |grep -Ev 'pkg/clientsets|pkg/dal|pkg/models|pkg/version|pkg/injector')

.PHONY: all
all: fmt test build

.PHONY: check-git-status
check-git-status: ## Check git status
	@test -z "$$(git status --porcelain)" || (echo "Git status is not clean, please commit or stash your changes first" && exit 1)

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: install-dev-tools
install-dev-tools: ## Install dev tools
# 	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
# 	go get -u github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	go install golang.org/x/tools/cmd/stringer@latest
	bash ./hack/install-dev-tools.sh

##@ Build

.PHONY: download
download: ## Run go mod download
	go mod download

.PHONY: fmt
fmt: ## Run format code
	gofmt -w ./pkg/ ./cmd/
	go fmt ./pkg/... ./cmd/...
	go vet ./pkg/... ./cmd/...
	goimports -w ./pkg/ ./cmd/
	golangci-lint run --fix

.PHONY: lint
lint: ## Run lint
	@echo "# gofmt"
	@test $$(gofmt -l . | wc -l) -eq 0

	@echo "# ensure integration test with // +build integration"
	@test $$(find test -name '*_test.go' | wc -l) -eq $$(cat $$(find test -name '*_test.go') | grep -E '// ?go:build integration' | wc -l)

	go mod tidy
	golangci-lint run --timeout 5m
	gofmt -w .

.PHONY: generate-code-mockery
generate-code-mockery: ## Run generate generated unit test code
	# 如果遇到问题
	# Unexpected package creation during export data loading
	# https://github.com/vektra/mockery/pull/435#issuecomment-1134329306
	@echo "# generate mock of interfaces for testing"
	@rm -rf test/mock
	@mkdir -p test/mock
	@(cd . && mockery --all --keeptree --case=underscore --packageprefix=mock --output=./test/mock/)
	# mockery not support 1.18 generic now, temporarily drop zero size golang file
	# https://github.com/vektra/mockery/pull/456
	find test/mock -size 0 -exec rm {} \;

.PHONY: generate-code-enum
generate-code-enum: ## Generate enum String for models
	# go install golang.org/x/tools/cmd/stringer
	@echo generate stringer for enums
	# TODO using ls
	@(cd pkg/models/enums/tasks/priority; go generate)
	@(cd pkg/models/enums/tasks/status; go generate)
	@(cd pkg/models/enums/tasks/subtasksview; go generate)

.PHONY: generate-manual
generate-manual: ## Generate develop docs
	# go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	gomarkdoc --output GO_DOC.md ./...

	rm -rf docs/dingmark/commands
	mkdir -p docs/commands
	go run ./cmd/gendocs/

.PHONY: build
build: ## Build
	go build -v -o $(OUTPUT_DIR)/dingmark-$(GOOS)-$(GOARCH) \
		 -ldflags "-s -w -X $(ROOT)/pkg/version.Version=$(VERSION) -X $(ROOT)/pkg/version.Commit=$(COMMIT) -X $(ROOT)/pkg/version.Package=$(ROOT)" \
		 $(CMD_DIR)/dingmark
	cp $(OUTPUT_DIR)/dingmark-$(GOOS)-$(GOARCH) $(OUTPUT_DIR)/dingmark-$(VERSION)-$(GOOS)-$(GOARCH)

.PHONY: build-wasm
build-wasm: ## Build WASM
	GOOS=js GOARCH=wasm go build -v -o $(OUTPUT_DIR)/dingmark-js-wasm \
		-ldflags "-s -w -X $(ROOT)/pkg/version.Version=$(VERSION) -X $(ROOT)/pkg/version.Commit=$(COMMIT) -X $(ROOT)/pkg/version.Package=$(ROOT)" \
		 $(CMD_WASM_DIR)/dingmark
	 
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js $(PROJECT_DIR)/static/wasm_exec.js
	cp $(OUTPUT_DIR)/dingmark-js-wasm $(OUTPUT_DIR)/dingmark-$(VERSION)-js-wasm

	cp $(OUTPUT_DIR)/dingmark-js-wasm $(PROJECT_DIR)/static/main.wasm

.PHONY: test
test: ## Run unit tests
	@# NOTICE, the test output is using for coverage analytics, did not modify the std out
	@echo "cover package: ${UT_COVER_PACKAGES}"
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go test -v ${UT_COVER_PACKAGES} -coverprofile cover.out -tags=\!integration ./...

.PHONY: integration-test
integration-test: ## Run integration tests
	@echo "cover package: ${IT_COVER_PACKAGES}"
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go test -v ${IT_COVER_PACKAGES} -coverprofile cover.out -tags=integration ./...

.PHONY: clean
clean: ## Clean temp files
	@rm -vrf ${OUTPUT_DIR}/*

.PHONY: bump
STAGE=$(DEFAULT_BUMP_STAGE)
SCOPE=$(DEFAULT_BUMP_SCOPE)
DRY_RUN=$(DEFAULT_BUMP_DRY_RUN)
bump: check-git-status ## Bump version
	(bash ./hack/bump.sh ${STAGE} ${SCOPE} ${DRY_RUN})

