.PHONY: start build run lint clean doc dev docker-build docker-image docker-image-staging docker-image-dev start-docker-dev stop-docker-dev

GONAME=api

default: build

start:
	@GIN_MODE=release ./bin/$(GONAME)

build:
	@export GOPROXY=https://goproxy.cn && export GO111MODULE=on && go build -o bin/$(GONAME)

clean:
	@go clean && rm -rf ./bin/$(GONAME) && rm -f gin-bin

doc:
	godoc -http=:6060 -index

dev:
	@export GOPROXY=https://goproxy.cn && export GO111MODULE=on && go run main.go
	# @gin -a 8086 -p 3030 run main.go

init:
	@sh ./devops/grpc_gen.sh

update:
	@git submodule foreach git pull --rebase --allow-unrelated-histories

# generate_c:
#	protoc --proto_path=devops/idl --go_out=plugins=grpc:grpc_gen devops/idl/user/info.proto \
#     protoc --proto_path=devops/idl --go_out=plugins=grpc:grpc_gen devops/idl/comment/info.proto \
#     protoc --proto_path=devops/idl --go_out=plugins=grpc:grpc_gen devops/idl/content/info.proto