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

.PHONY: test_quick
test_quick:
	@go test -v ./...

.PHONY: test_sdk_ipc
test_sdk_ipc:
	@NODE_RPC_ADDRESS=127.0.0.1:5000 PRIVATE_KEY=$(PRIVATE_KEY) DIAL_URI=$(DIAL_URI) go test -v -count=1 -timeout 10m -run '^TestIPC' ./sdk/...

.PHONY: test_sdk_streaming
test_sdk_streaming:
	@NODE_RPC_ADDRESS=127.0.0.1:5000 go test -v -count=1 -timeout 10m -run '^TestStreaming' ./sdk/...

.PHONY: test_cli_ipc
test_cli_ipc:
	@NODE_RPC_ADDRESS=127.0.0.1:5000 PRIVATE_KEY=$(PRIVATE_KEY) go test -v -count=1 -timeout 6m -run '^TestIPC' ./cmd/akavecli/...

.PHONY: test_cli_streaming
test_cli_streaming:
	@NODE_RPC_ADDRESS=127.0.0.1:5000 go test -v -count=1 -timeout 6m -run '^TestStreaming' ./cmd/akavecli/...

.PHONY: check
check: # for local usage
	@golangci-lint run ./... 
	@linelint .

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
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2
	@go install github.com/fernandrone/linelint@0.0.6
