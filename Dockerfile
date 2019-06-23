FROM golang:latest

WORKDIR $GOPATH/src/github.com/sundogrd/content-api
COPY . $GOPATH/src/github.com/sundogrd/content-api

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

ARG GITHUB_CLIENT_ID
ARG GITHUB_SECRET

RUN go build .
RUN "./devops/build_docker.sh" $GITHUB_CLIENT_ID $GITHUB_SECRET

#RUN ip -4 route list match 0/0 | awk '{print $3 " host.docker.internal"}' >> /etc/hosts

EXPOSE 8086
ENTRYPOINT ["./devops/entrypoint.sh"]