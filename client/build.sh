#!/bin/sh
echo "CLIENT | Initializing environment..."
export GO111MODULE=on

echo "CLIENT | Fetching dependencies..."
if [[ "$1" == "upgrade" ]]; then
    go get -u
else
    go get
fi
go install github.com/micro/protoc-gen-micro
go install github.com/golang/protobuf/protoc-gen-go

echo "CLIENT | Creating executable..."
go build -o client main.go