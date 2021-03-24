IMAGE_NAME=pets-api
IMAGE_VERSION ?= 0

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
	GOOS=linux GOARCH=amd64 go build -o bin/pets_api main.go

.PHONY: docker-build
docker-build:
	docker build -t $IMAGE_NAME:latest .

docker-tag:
	docker tag $IMAGE_NAME:dev $IMAGE_NAME:$IMAGE_VERSION