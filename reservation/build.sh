#!/bin/sh
echo "RESERVATION SERVICE | Initializing environment..."
export GO111MODULE=on

echo "RESERVATION SERVICE | Fetching dependencies..."
go get
go install github.com/micro/protoc-gen-micro
go install github.com/golang/protobuf/protoc-gen-go

echo "RESERVATION SERVICE | Compiling protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "RESERVATION SERVICE | Creating executable..."
go build -o reservation-service main.go