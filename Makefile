SHELL := /bin/bash
.PHONY: help

help: ## Display this help message
	@echo -e "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' | column -c2 -t -s :)"

docker: ## Start the container
	docker-compose up -d

docker-go: ## Start the container
	docker-compose up -d --force-recreate --no-deps --build adopet-go

graphql: ## Graphql generator
	gqlgen --verbose --config=pkg/graph/gqlgen.yml

setup: ## Project settings
	go install github.com/rafaelsq/wtc@latest

lint: ## Lint the project
	golangci-lint run ./...

start-dev: ## Initializes development watches
start-dev:
	@wtc