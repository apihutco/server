.PHONY: run build image

export GO111MODULE=on
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=amd64

all: build

run:
	@go run *.go -f ./conf/config.yaml

build: tidy
	@go build -o server -ldflags "-s -w" .

tidy:
	@go mod tidy

image:
	@podman build -t apihut-server:latest .

clean:
	@rm -rf data server