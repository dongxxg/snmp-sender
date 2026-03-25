.PHONY: build run fmt vet test

build:
	go build ./...

run:
	go run .

fmt:
	gofmt -w .
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...
