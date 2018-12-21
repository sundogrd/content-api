FROM alpine:3.8

ENV appdir /app
ARG port=8086

RUN mkdir -p $appdir
WORKDIR $appdir

ADD Makefile .
ADD ./bin ./bin
ADD ./data/config/app.json ./data/config/app.json

RUN echo http://mirrors.aliyun.com/alpine/v3.8/main > /etc/apk/repositories \
    && echo http://mirrors.aliyun.com/alpine/v3.8/main >> /etc/apk/repositories \
    && apk add --no-cache make

EXPOSE $port
