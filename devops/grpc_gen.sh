#!/usr/bin/env bash

protoc --proto_path=devops/idl --go_out=plugins=grpc:grpc_gen devops/idl/user/info.proto
protoc --proto_path=devops/idl --go_out=plugins=grpc:grpc_gen devops/idl/comment/info.proto