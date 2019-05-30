#!/bin/sh
echo "PRESENTATION SERVICE | Compiling protocol buffers..."
chmod +x ./proto/build.sh
cd proto
. ./build.sh
cd ..

echo "PRESENTATION SERVICE | Creating executable..."
go build -o presentation-service main.go