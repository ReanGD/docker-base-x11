export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: test bench dep lint lint_ci fmt build run postclone precommit clean help

test: ## Run unittests
	@env go test -v -p=1 -count=1 ./...

bench: ## Run benchmarks
	@env go test -bench=. -benchtime=10s -benchmem ./...

dep: ## Get the dependencies
	@env go get -v -t -d ./...
	@env go get -v -u github.com/sqs/goreturns
	@env go get -v -u github.com/golangci/golangci-lint/cmd/golangci-lint

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

precommit: fmt lint_ci test ## Precommit actions

clean: ## Remove previous build
	@env go clean

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
