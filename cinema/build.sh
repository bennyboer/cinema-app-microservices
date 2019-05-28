#!/bin/sh
echo "Compiling cinema protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "Creating cinema executable..."
go build -o cinema-service main.go
cd ..