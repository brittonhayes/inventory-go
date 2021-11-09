GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

.PHONY: init
# Install dependencies
init:
	go install entgo.io/ent/cmd/ent@latest
	go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@latest

.PHONY: install
# Install server and cli
install:
	go install ./cmd/inventory
	go install ./cmd/inventoryctl

.PHONY: build
# build binary
build:
	mkdir -p bin/
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/inventory
	go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./cmd/inventoryctl

.PHONY: migrate
# run database migrations
migrate:
	go run cmd/inventoryctl/main.go migrate

.PHONY: run
# run the gRPC and API service
run:
	go run cmd/inventory/main.go

.PHONY: generate
# run go generators
generate:
	go generate ./...

.PHONY: test
# run test suite
test:
	go test -v ./... -cover

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
