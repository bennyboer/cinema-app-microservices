#!/bin/sh
echo "USER SERVICE | Compiling protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "USER SERVICE | Creating executable..."
go build -o user-service main.go