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
