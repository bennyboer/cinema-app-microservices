#!/bin/sh
echo "MOVIE SERVICE | Initializing environment..."
export GO111MODULE=on

echo "MOVIE SERVICE | Fetching dependencies..."
if [[ "$1" == "docker" ]]; then
    git config --global user.name "bennyboer-machine-user"
    git config --global credential.helper store
    echo https://51faa31d4b9f08c8e56d4fb23fc082a85e617df8:x-oauth-basic@github.com >> ~/.git-credentials
fi
if [[ "$1" == "upgrade" ]]; then
    go get -u
else
    go get
fi
go install github.com/micro/protoc-gen-micro
go install github.com/golang/protobuf/protoc-gen-go

echo "MOVIE SERVICE | Compiling protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "MOVIE SERVICE | Creating executable..."
go build -o movie-service main.go