FROM golang:latest

WORKDIR $GOPATH/src/github.com/sundogrd/content-api
COPY . $GOPATH/src/github.com/sundogrd/content-api

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

RUN go build .

EXPOSE 8086
ENTRYPOINT ["./content-api"]