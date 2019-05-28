#!/bin/sh
echo "Compiling user protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "Creating user executable..."
go build -o user-service main.go
cd ..