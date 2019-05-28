#!/bin/sh
echo "Fetching dependencies..."
go get

echo "Compiling user service..."
chmod +x ./user/build.sh
cd user
. ./build.sh
cd ..