.PHONY: dev build tidy image clean
VERSION=$(shell git describe --always)
TIME_NOW=$(shell date +"%Y-%m-%d_%H:%M:%S")
include .env

export GO111MODULE=on
export CGO_ENABLED=1
export GOPROXY=${CUSTOM_GOPROXY}

all: build

dev:
	go run -ldflags "-X github.com/apihutco/server/config.VERSION=$(VERSION) -X github.com/apihutco/server/config.BUILD_TIME=$(TIME_NOW)" *.go -f ./conf/config.yaml

build:
	go build -o ./bin/apihut -ldflags "-s -w -X github.com/apihutco/server/config.VERSION=$(VERSION) -X github.com/apihutco/server/config.BUILD_TIME=$(TIME_NOW)" .

update:
	go get -u -v

image:
	make image_build_dev
	# make image_push_dev

image_build_dev:
	make build
	# 仅构建 dev 镜像
	${CONTAINER_BUILDER} build --push -f deploy/docker/DockerfileDev --platform linux/amd64 -t ${TARGET_IMAGE}:dev -t ${TARGET_IMAGE}:dev-${VERSION} --build-arg BUILDER_IMAGR=${BUILDER_IMAGR} --build-arg RUNNER_IMAGE=${RUNNER_IMAGE} --build-arg CUSTOM_GOPROXY=${CUSTOM_GOPROXY} .

image_push_dev:
	# 仅推送 dev 镜像
	${CONTAINER_BUILDER} push ${TARGET_IMAGE}:dev-${VERSION}
	${CONTAINER_BUILDER} push ${TARGET_IMAGE}:dev

image_publish_dev:
	# 需提前切换好相应的 builder
	${CONTAINER_BUILDER} build -f deploy/docker/DockerfilePublish --platform linux/amd64,linux/arm64 --push -t ${TARGET_IMAGE}:dev -t ${TARGET_IMAGE}:dev-${VERSION} --build-arg BUILDER_IMAGR=${BUILDER_IMAGR} --build-arg RUNNER_IMAGE=${RUNNER_IMAGE} --build-arg CUSTOM_GOPROXY=${CUSTOM_GOPROXY} .

image_publish_prod:
	# 需提前切换好相应的 builder
	${CONTAINER_BUILDER} build -f deploy/docker/DockerfilePublish --platform linux/amd64,linux/arm64 --push -t ${TARGET_IMAGE}:latest -t ${TARGET_IMAGE}:dev-${VERSION} --build-arg BUILDER_IMAGR=${BUILDER_IMAGR} --build-arg RUNNER_IMAGE=${RUNNER_IMAGE} --build-arg CUSTOM_GOPROXY=${CUSTOM_GOPROXY} .

clean:
	@rm -rf data bin