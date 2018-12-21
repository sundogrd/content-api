FROM golang:1.11-alpine3.8

ENV appdir /go/src/github.com/sundogrd/content-api
RUN mkdir -p $appdir
WORKDIR $appdir

RUN echo http://mirrors.aliyun.com/alpine/v3.8/main > /etc/apk/repositories \
    && echo http://mirrors.aliyun.com/alpine/v3.8/main >> /etc/apk/repositories \
    && apk add --no-cache git make bash gcc musl-dev \
    # && git config --global http.sslVerify false \
    # && go get -insecure -v github.com/codegangsta/gin
