#!/bin/sh
echo "Compiling protocol buffers..."
chmod +x /proto/build.sh
. ./proto/build.sh

echo "Creating executable..."
go build -o user-service main.go