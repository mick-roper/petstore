.PHONY: all
all: clean-cli test-cli build-cli

.PHONY: clean-cli
clean-cli:
	rm -rf bin

.PHONY: test-cli
test-cli:
	go test ./...

.PHONY: build-cli
build-cli:
	mkdir bin
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/cli main.go