#!/bin/sh
echo "Fetching dependencies..."
go get

echo "Compiling user service..."
cd ./user
chmod +x ./build.sh
./build.sh
cd ..

read -n 1 -s -r -p "Press any key to continue"