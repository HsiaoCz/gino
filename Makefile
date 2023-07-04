run: build
	@./bin/gino

build:
	@go build -o bin/gino cmd/main.go

test:
	@go test -v ./...

.PHONY: run build test

all:run

help:
	echo "run:run ./bin/gino"
	echo "build: go build -o bin/gino"
	echo "test: go test -v ./..."