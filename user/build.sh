#!/bin/sh
echo "USER SERVICE | Initializing environment..."
export GO111MODULE=on

echo "USER SERVICE | Fetching dependencies..."
go get
go install github.com/micro/protoc-gen-micro
go install github.com/golang/protobuf/protoc-gen-go

echo "USER SERVICE | Compiling protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "USER SERVICE | Creating executable..."
go build -o user-service main.go