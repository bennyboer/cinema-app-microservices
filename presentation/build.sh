#!/bin/sh
echo "PRESENTATION SERVICE | Initializing environment..."
export GO111MODULE=on

echo "PRESENTATION SERVICE | Fetching dependencies..."
if [[ "$1" == "upgrade" ]]; then
    go get -u
else
    go get
fi
go install github.com/micro/protoc-gen-micro
go install github.com/golang/protobuf/protoc-gen-go

echo "PRESENTATION SERVICE | Compiling protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "PRESENTATION SERVICE | Creating executable..."
go build -o presentation-service main.go