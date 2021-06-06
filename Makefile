include .env
export

APP_NAME=zephry-one

GOPATH=$(shell go env GOPATH)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

CURRENT_PATH=$(shell pwd)
GO111MODULE=on

run:
	go run cmd/main/main.go

server-run: build-${GOOS}
	./build/${APP_NAME}-${GOOS}.app

build-go:
	go build -o ./build/${APP_NAME}.app cmd/main/main.go
	file ./build/${APP_NAME}.app

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/${APP_NAME}-linux.app cmd/main/main.go
	file ./build/${APP_NAME}-linux.app

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./build/${APP_NAME}-windows.exe cmd/main/main.go
	file ./build/${APP_NAME}-windows.exe

build-all: build-linux build-windows