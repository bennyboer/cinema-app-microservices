#!/bin/sh
echo "Initializing environment..."
export GO111MODULE=on

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

echo "Building presentation service..."
chmod +x ./presentation/build.sh
cd presentation
. ./build.sh
cd ..

echo "Building reservation service..."
chmod +x ./reservation/build.sh
cd reservation
. ./build.sh
cd ..

if [[ "$1" == "debug" ]]; then
    read -n 1 -s -r -p "Press any key to continue"
fi