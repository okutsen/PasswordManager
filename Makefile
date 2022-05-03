include local.env

NAME=Password-Manager
DOCKER_NAME=password-manager
PORT=10000
MAIN_PATH=./cmd/main.go
TARGET_PATH=$(GOPATH)/bin
TARGET=${TARGET_PATH}/${NAME}

.PHONY: dependencies
dependencies:
	go mod vendor

.PHONY: build
build:
	GOARCH=amd64 GOOS=darwin go build -o ${TARGET} ${MAIN_PATH}

.PHONY: run
run:
	${TARGET_PATH}/${NAME}

.PHONY: dockerBuild
dockerBuild:
	docker build -t ${DOCKER_NAME} ./

.PHONY: dockerRun
dockerRun:
	docker run -p ${PORT}:${PORT} --name="${DOCKER_NAME}" ${DOCKER_NAME}

.PHONY: dockerStart
dockerStart: dockerBuild dockerRun

.PHONY: dockerStop
dockerStop:
	docker stop ${DOCKER_NAME}; docker rm ${DOCKER_NAME}; docker rmi -f ${DOCKER_NAME}

.PHONY: up
up: dependencies build run

clean:
	go clean
	rm ${TARGET}