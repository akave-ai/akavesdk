SHELL := /bin/bash
TZ=UTC

# set local envs if any.
ifneq (,$(wildcard .env))
  include .env
  export
endif

.PHONY: build
build:
	@go build -o bin/akavecli ./cmd/akavecli/

.PHONY: test
test:
	@go test -v -race -count=1 ./...

.PHONY: testsum
testsum:
	@gotestsum --format testname -- -race -count=1 ./...

.PHONY: test_quick
test_quick:
	@go test -v ./...

.PHONY: test_sdk
test_sdk:
	@NODE_RPC_ADDRESS=127.0.0.1:5000 PRIVATE_KEY=$(PRIVATE_KEY) DIAL_URI=$(DIAL_URI) go test -v -count=1 -timeout 10m -run '^TestIPC' ./sdk/...

.PHONY: test_cli
test_cli:
	@NODE_RPC_ADDRESS=127.0.0.1:5000 PRIVATE_KEY=$(PRIVATE_KEY) go test -v -count=1 -timeout 6m -run '^TestExternal' ./cmd/akavecli/...

.PHONY: check
check: # for local usage
	@golangci-lint run ./... 
	@linelint .
	@go mod tidy -diff

.PHONY: lint_fix
lint_fix: # for local usage
	@golangci-lint run --fix ./...
	@linelint -a .

.PHONY: fix
fix: lint_fix

.PHONY: clean
clean: # for local usage
	@rm -rf bin/*

.PHONY: install_tools
install_tools: # for local usage
	@go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.5.0
	@go install github.com/fernandrone/linelint@0.0.6
	@go install gotest.tools/gotestsum@v1.13.0
