#!/bin/sh
echo "Compiling protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "Creating executable..."
go build -o user-service main.go
cd ..