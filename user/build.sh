#!/bin/sh

echo "Compiling protocol buffers..."
cd ./proto
chmod +x ./build.sh
./build.sh
cd ..

echo "Creating executable..."
go build -o user-service main.go
