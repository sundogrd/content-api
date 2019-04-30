.PHONY: start build run lint clean doc dev docker-build docker-image docker-image-staging docker-image-dev start-docker-dev stop-docker-dev

GONAME=api

default: build

start:
	@GIN_MODE=release ./bin/$(GONAME)

build:
	@export GO111MODULE=on && export GOFLAGS=-mod=vendor && go build -o bin/$(GONAME)

clean:
	@go clean && rm -rf ./bin/$(GONAME) && rm -f gin-bin

doc:
	godoc -http=:6060 -index

dev:
	@go build && go run main.go
	# @gin -a 8086 -p 3030 run main.go

init:
	@sh ./devops/grpc_gen.sh

update:
	@git submodule foreach git pull
