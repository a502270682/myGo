SHELL = /bin/bash

all: lint test build

lint: ## Lint the files
	@echo "check lint"
	@echo "Installing golangci-lint"
	@GO111MODULE=on	 go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.0
	@echo "start cache clean"
	@golangci-lint cache clean
	@echo "start lint check"
	@golangci-lint run --timeout=5m --config ./.golangci.yml

test: ## Run unittests
	@echo "check test"
	@go test ./...

dep: ## Get the dependencies
	@echo "check dep"
	@go mod tidy

build: dep ## Build the binary file
	@echo "make build"
	@chmod 755 ./build.sh
	./build.sh ./cmd/server/main.go
