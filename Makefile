BINARY_NAME:=template


# =================================== DEFAULT =================================== #

default: all

## default: Runs build and test
.PHONY: default
all: build test

# =================================== HELPERS =================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Template - You can run the CLI with "go run main.go"'
	@echo ''
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Commands:'
	@sed -n 's/^## //p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
	@echo ''
	@echo 'Extra:'
	@sed -n 's/^### //p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'




## install: Install dependencies
.PHONY: install
install: install/go install/npm

### install/npm: Install npm dependencies
.PHONY: install/npm
install/npm:
	npm i

### install/go: Install go dependencies
.PHONY: install/go
install/go:
	go get ./...




## links: Shows the project links
.PHONY: links
links:
	@echo 'Links:'
	@echo ' Github Repository:      https://github.com/caffeine-addictt/template'
	@echo ' Official Documentation: https://github.com/caffeine-addictt/template/blob/main/docs/index.md'




## issue: Where to create an issue
.PHONY: issue
issue:
	@echo 'Create an issue at:'
	@echo ' https://github.com/caffeine-addictt/template/issues/new'




## docs: Shows simple development documentation
.PHONY: docs
docs:
	@echo 'Development documentation'
	@echo ''
	@echo 'Prerequisites:'
	@echo ' 1. Go 1.23.0 or later'
	@echo ' 3. NPM 10.8.2 or later'
	@echo ' 4. Node 22.7.0 or later'
	@echo ''
	@echo 'Steps to run the CLI:'
	@echo ' 1. Run the CLI with "go run main.go"'
	@echo ''
	@echo 'Learn more at https://github.com/caffeine-addictt/template/blob/main/CONTRIBUTING.md'

# =================================== DEVELOPMENT =================================== #

## build: Builds Python, Go binaries and Docker image
.PHONY: build
build:
	npx cross-env GOARCH=amd64 GOOS=linux   go build -ldflags="-s -w" -o ./bin/$(BINARY_NAME)-linux main.go
	npx cross-env GOARCH=amd64 GOOS=darwin  go build -ldflags="-s -w" -o ./bin/$(BINARY_NAME)-darwin main.go
	npx cross-env GOARCH=amd64 GOOS=windows go build -ldflags="-s -w" -o ./bin/$(BINARY_NAME)-windows.exe main.go




## test: Runs tests
.PHONY: test
test:
	go mod tidy
	go mod verify
	go vet ./...
	go run github.com/securego/gosec/v2/cmd/gosec@latest -quiet ./...
	go run github.com/go-critic/go-critic/cmd/gocritic@latest check -enableAll ./...
	go run github.com/google/osv-scanner/cmd/osv-scanner@latest -r .
	go test -v -race ./...

# =================================== QUALITY ================================== #

## lint: Lint code
.PHONY: lint
lint: lint/go lint/npm

### lint/go: Lint Go code
.PHONY: lint/go
lint/go:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

### lint/npm: Lint NPM code
.PHONY: lint/npm
lint/npm:
	npm run lint




## format: Format code
.PHONY: format
format: format/go format/npm

### format/go: Format Go code
.PHONY: format/go
format/go:
	go fmt ./...
	go mod tidy -v
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run --fix

### format/npm: Format NPM code
.PHONY: format/npm
format/npm:
	npm run lint:fix




## tidy: Clean up code artifacts
.PHONY: tidy
tidy:
	go clean ./...
	${RM_CMD} bin




## clean: Remove node_modules
.PHONY: clean
clean: tidy
	${RM_CMD} node_modules
