.PHONY: all
all: clean test build

.PHONY: clean
clean:
	rm -rf bin

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	mkdir bin
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/petstore main.go