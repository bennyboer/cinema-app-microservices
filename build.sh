#!/bin/sh
echo "Fetching dependencies..."
go get
go install github.com/micro/protoc-gen-micro
go install github.com/golang/protobuf/protoc-gen-go

echo "Building user service..."
chmod +x ./user/build.sh
cd user
. ./build.sh
cd ..

echo "Building movie service..."
chmod +x ./movie/build.sh
cd movie
. ./build.sh
cd ..

if [[ "$1" == "debug" ]]; then
    read -n 1 -s -r -p "Press any key to continue"
fi