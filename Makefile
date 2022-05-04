NAME=Password-Manager
PORT=10000
DOCKER_NAME=password-manager
MAIN_PATH=./cmd/main.go
TARGET_PATH=$(GOPATH)/bin
TARGET=${TARGET_PATH}/${NAME}
.DEFAULT_GOAL := help
export PM_PORT=${PORT}

write:
	@echo ${PORT}

dependencies: ## Update dependencies
	go mod vendor

build: ## Make build of the project
	GOARCH=amd64 GOOS=darwin go build -o ${TARGET} ${MAIN_PATH}

run: ## Run the project
	${TARGET_PATH}/${NAME}

dockerBuild: ## Create an image in docker
	docker build -t ${DOCKER_NAME} ./

dockerRun: ## Run container
	docker run -p ${PORT}:${PORT} --name="${DOCKER_NAME}" ${DOCKER_NAME}

dockerStart: dockerBuild dockerRun ## Create an image in docker and run a container

dockerStop: ## Delete an image and container with name "password-manager"
	docker stop ${DOCKER_NAME}; docker rm ${DOCKER_NAME}; docker rmi -f ${DOCKER_NAME}

up: dependencies build run ## Update dependencies, build the project and run it

clean:
	go clean
	rm ${TARGET}

help: ## Display this help screen
	@grep -E -h '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
