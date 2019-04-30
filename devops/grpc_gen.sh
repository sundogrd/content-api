#!/usr/bin/env bash

protoc --proto_path=devops/idl --go_out=plugins=grpc:grpc_gen devops/idl/**/*.proto