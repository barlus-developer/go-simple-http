SHELL := /bin/sh
.DEFAULT_GOAL := run

.PHONY: run test build fmt

run:
	go run ./cmd/server

test:
	go test ./...

build:
	go build ./cmd/server

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './.git/*')
