export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: test dep lint lint_ci fmt build run postclone precommit clean help

test: ## Run unittests
	@go test -v -p=1 -count=1

dep: ## Get the dependencies
	@go get -v -t -d ./...
	@go get -v -u github.com/sqs/goreturns
	@go get -v -u github.com/golangci/golangci-lint/cmd/golangci-lint

lint: ## Lint the files
	@golangci-lint run --config=./ci/.golangci.yml

lint_ci: ## Lint the files as CI
	@golangci-lint run --config=./ci/.golangci_full.yml

fmt: ## Format source code
	@goreturns -w .

build: ## Build a binary file
	@env go build

run: build ## Build and run a binary file
	@./go-algo

postclone: dep ## Post clone actions

precommit: fmt lint_ci ## Precommit actions

clean: ## Remove previous build
	@go clean

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
