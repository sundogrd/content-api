#!/usr/bin/env bash

# update repo although it will not be used
git pull

# update the repository
docker pull sundogrd/content-api:$1-$2
if docker ps -a | grep -q sundogrd-content-api; then
    docker rm -f sundogrd-content-api
fi
docker run -d --name sundogrd-content-api -p 9431:8086 sundogrd/content-api:$1-$2
