#!/bin/sh
echo "Building user service..."
chmod +x ./user/build.sh
cd user
. ./build.sh "$1"
cd ..

echo "Building cinema service..."
chmod +x ./cinema/build.sh
cd cinema
. ./build.sh "$1"
cd ..

echo "Building movie service..."
chmod +x ./movie/build.sh
cd movie
. ./build.sh "$1"
cd ..

echo "Building presentation service..."
chmod +x ./presentation/build.sh
cd presentation
. ./build.sh "$1"
cd ..

echo "Building reservation service..."
chmod +x ./reservation/build.sh
cd reservation
. ./build.sh "$1"
cd ..

echo "Building client..."
chmod +x ./client/build.sh
cd client
. ./build.sh "$1"
cd ..

if [[ "$1" == "debug" ]]; then
    read -n 1 -s -r -p "Press any key to continue"
fi