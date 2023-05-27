GO_DIR ?= $(shell pwd)
GOOS ?= $(shell go env GOOS || echo linux)
GOARCH ?= $(shell go env GOARCH || echo amd64)
CGO_ENABLED ?= 0
DOCKER_IMAGE ?= ses-school-3-currency
TAG ?= $(shell git describe --tags --abbrev=0 || "not-set")
REMOVE_CONTAINERS ?= ON

init: ## init packages
	mkdir -p build &&\
    rm -rf build/* 

build: init ## build binary file
	GOOS=${GOOS} CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} \
	go build -ldflags "-X 'main.appVersion=$(TAG)-$$(date -u +%Y%m%d%H%M)'" -o "$(GO_DIR)/build/bin/currencies" cmd/currency_service/main.go

docker-image: ## build docker image
	REMOVE_CONTAINERS=${REMOVE_CONTAINERS} DOCKER_IMAGE=${DOCKER_IMAGE} ./scripts/remove_docker_containers.sh
	docker rmi ${DOCKER_IMAGE}:${TAG} -f || true ;\
	docker build -f "${GO_DIR}/docker/app/Dockerfile" -t ${DOCKER_IMAGE}:${TAG} ${GO_DIR}

test: ## test application with race
	go test -v ./...

coverage: ## test coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -html coverage.out

.PHONY: init
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help