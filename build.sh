#!/bin/sh
echo "Fetching dependencies..."
go get
go install github.com/micro/protoc-gen-micro
go install github.com/golang/protobuf/protoc-gen-go

echo "Compiling user service..."
chmod +x ./user/build.sh
cd user
. ./build.sh
cd ..