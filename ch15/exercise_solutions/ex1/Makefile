.DEFAULT_GOAL := build

.PHONY: fmt vet build clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test -bench ... -benchmem -cover

build: test
	go build

clean:
	go clean

fuzz:
	go test -trimpath -fuzz=Fuzz_parser

race:
	go test -race -trimpath
