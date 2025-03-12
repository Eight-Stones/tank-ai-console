BIN_EXAMPLE:= "./bin/example"
DOCKER_EXAMPLE_IMG="example:develop"
BIN_MIGRATOR= "./bin/migrator"
DOCKER_MIGRATOR_IMG="migrator:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

deps-proto-utils:
	sh scripts/deps.sh

install-deps: deps-proto-utils

build-example:
	CGO_ENABLED=0 GOOS=linux \
	go build -v -o $(BIN_EXAMPLE) \
	-ldflags "$(LDFLAGS)" \
	./cmd/example

run-example: build
	$(BIN_EXAMPLE) -config ./config/example.toml

build-migrator:
	CGO_ENABLED=0 GOOS=linux \
	go build -v -o $(BIN_MIGRATOR) \
	-ldflags "$(LDFLAGS)" \
	./cmd/migrator

build-img-migrator:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_MIGRATOR_IMG) \
		-f docker/migrator/Dockerfile .

run-img-migrator: build-img-example
	docker run $(DOCKER_MIGRATOR_IMG)

version: build
	$(BIN_EXAMPLE) version

test:
	go test -race ./internal/... ./pkg/...

lint:
	golangci-lint run ./... --config .golangci.yml


arch-lint:
	go-arch-lint check

arch-graph:
	go-arch-lint graph --type di

goose-get:
	go install github.com/pressly/goose/v3/cmd/goose@latest

up:
	docker-compose -f docker-compose.yaml up -d --force-recreate

down:
	docker-compose -f docker-compose.yaml down


# buf command section
buf-build:
	buf build

buf-gen:
	buf generate

buf-lint:
	buf lint

.PHONY: build run build-img run-img version test lint