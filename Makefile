.PHONY: build_and up down test build install-tools lint deps

build_and:
	docker-compose -f docker-compose.yml build

up:
	docker-compose -f docker-compose.yml up -d --build

down:
	docker-compose -f docker-compose.yml down

test:
	go test ./... -v

build:
	go build ./cmd

install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

lint:
	golangci-lint run ./...

deps:
	go mod download
	go mod tidy