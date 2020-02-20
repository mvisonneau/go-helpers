NAME          := go-helpers
FILES         := $(shell find * -type f -name '*.go')
REPOSITORY    := mvisonneau/$(NAME)
.DEFAULT_GOAL := help

export GO111MODULE=on

.PHONY: setup
setup: ## Install required libraries/tools for build tasks
	@command -v goveralls 2>&1 >/dev/null || GO111MODULE=off go get -u -v github.com/mattn/goveralls
	@command -v golint 2>&1 >/dev/null    || GO111MODULE=off go get -u -v golang.org/x/lint/golint
	@command -v cover 2>&1 >/dev/null     || GO111MODULE=off go get -u -v golang.org/x/tools/cmd/cover
	@command -v goimports 2>&1 >/dev/null || GO111MODULE=off go get -u -v golang.org/x/tools/cmd/goimports

.PHONY: fmt
fmt: setup ## Format source code
	gofmt -s -w $(FILES)
	goimports -w $(FILES)

.PHONY: lint
lint: setup ## Run golint, goimports and go vet against the codebase
	golint -set_exit_status .
	go vet ./...
	goimports -d $(FILES) > goimports.out
	@if [ -s goimports.out ]; then cat goimports.out; rm goimports.out; exit 1; else rm goimports.out; fi

.PHONY: test
test: ## Run the tests against the codebase
	go test -v ./...

.PHONY: coverage
coverage: ## Generates coverage report
	rm -rf *.out
	go test -v ./... -coverpkg=./... -coverprofile=coverage.out

.PHONY: show-coverage
show-coverage: ## Display coverage report in browser
	go tool cover -html=coverage.out

.PHONY: sign-drone
sign-drone: ## Sign Drone CI configuration
	drone sign $(REPOSITORY) --save

.PHONY: help
help: ## Displays this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
