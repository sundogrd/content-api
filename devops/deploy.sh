#!/usr/bin/env bash


# generate app.json using app.template.json
cp ./data/config/app.template.json ./data/config/app.json
sed -i '' 's#\("ClientID": "\).*#\1'"$2"'",#g' "./data/config/app.json"
sed -i '' 's#\("ClientSecret": "\).*#\1'"$3"'"#g' "./data/config/app.json"

# update the repository
docker pull sundogrd/content-api:$1
if docker ps -a | grep -q sundogrd-content-api; then
    docker rm -f sundogrd-content-api
fi
docker run -d --name sundogrd-content-api -p 9431:8086 sundogrd/content-api:$1
