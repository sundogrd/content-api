FROM golang:latest

WORKDIR $GOPATH/src/github.com/sundogrd/content-api
COPY . $GOPATH/src/github.com/sundogrd/content-api

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ARG GITHUB_CLIENT_ID
ARG GITHUB_SECRET

RUN go build .
RUN "./devops/build_docker.sh"

EXPOSE 8086
ENTRYPOINT ["./content-api"]