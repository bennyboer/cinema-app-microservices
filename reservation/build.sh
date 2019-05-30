#!/bin/sh
echo "RESERVATION SERVICE | Compiling protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "RESERVATION SERVICE | Creating executable..."
go build -o reservation-service main.go