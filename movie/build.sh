#!/bin/sh
echo "MOVIE SERVICE | Compiling protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "MOVIE SERVICE | Creating executable..."
go build -o movie-service main.go