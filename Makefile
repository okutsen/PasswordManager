include .env

NAME="Password-Manager"
TARGET_PATH=$(GOPATH)/bin
TARGET=${TARGET_PATH}/${NAME}

.PHONY: dependencies
dependencies:
	go mod vendor

.PHONY: build
build:
	GOARCH=amd64 GOOS=darwin go build -o ${TARGET}

.PHONY: run
run:
	${TARGET_PATH}/${NAME}

.PHONY: up
up: dependencies build run

clean:
	go clean
	rm ${TARGET}