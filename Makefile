.PHONY: build test

build:
	go build -o ./bin/modular ./main.go

test:
	go test ./...
