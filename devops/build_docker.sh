#!/usr/bin/env bash

# generate app.json using app.template.json
cp ./data/config/app.template.json ./data/config/app.json
sed -i '' 's#\("ClientID": "\).*#\1'"$1"'",#g' "./data/config/app.json"
sed -i '' 's#\("ClientSecret": "\).*#\1'"$2"'"#g' "./data/config/app.json"
