export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: test dep fmt build clean help

test: ## Run unittests
	@go test -v -p=1 -count=1

dep: ## Get the dependencies
	@go get -v -t -d ./...
	@go get -v -u github.com/sqs/goreturns

fmt: ## Format source code
	@goreturns -w .

build: ## Build a binary file
	@env go build

run: build ## Build and run a binary file
	@./go-algo

clean: ## Remove previous build
	@go clean

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
