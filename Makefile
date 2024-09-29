BINARY_NAME:=waku
TAG ?= latest

ifeq ($(OS),Windows_NT)
RM_CMD:=rd /s /q
NULL:=/dev/nul
EXT:=.exe
else
RM_CMD:=rm -rf
NULL:=/dev/null
EXT=
endif


# =================================== DEFAULT =================================== #

default: all

## default: Runs build and test
.PHONY: default
all: build test

# =================================== HELPERS =================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Waku - You can run the CLI with "go run main.go"'
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
install:
	go get ./...

# =================================== DEVELOPMENT =================================== #

## docs: Runs Documentation
.PHONY: docs
docs:
	mkdocs serve -f www/mkdocs.yml -a 0.0.0.0:8000




## build: Builds Go binary
.PHONY: build
build:
	go build -ldflags="-s -w" -o $(BINARY_NAME)$(EXT) main.go

## build/docs: Builds documentation
build/docs:
	mkdocs build -f www/mkdocs.yml

### build/docker: Builds Docker image
build/docker:
	docker build . -t $(BINARY_NAME):$(TAG)




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




## bench: Run benchmarks
bench:
	go test -v -bench=. -benchmem ./...




## bump: Quickly bump Waku version (X.X.X, X.X.X-rc.X, etc.)
bump:
	@VERSION=$(version) ; \
	if [ -z "$$VERSION" ]; then \
		echo "Usage: make bump version=x.x.x"; \
		exit 1; \
	fi; \
	./scripts/version_bump.sh "$$VERSION"

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
	prettier --cache --check .




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
	prettier --cache --write .




## tidy: Clean up code artifacts
.PHONY: tidy
tidy:
	go clean ./...
	${RM_CMD} $(BINARY_NAME)$(EXT)




## clean: Remove node_modules
.PHONY: clean
clean: tidy
	${RM_CMD} node_modules
