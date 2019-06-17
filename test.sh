#!/bin/sh
echo "Testing user service..."
cd user
chmod +x ./test.sh
. ./test.sh
cd ..

echo "Testing cinema service..."
cd cinema
chmod +x ./test.sh
. ./test.sh
cd ..

echo "Testing movie service..."
cd movie
chmod +x ./test.sh
. ./test.sh
cd ..

echo "Testing presentation service..."
cd presentation
chmod +x ./test.sh
. ./test.sh
cd ..

echo "Testing reservation service..."
cd reservation
chmod +x ./test.sh
. ./test.sh
cd ..

if [[ "$1" == "debug" ]]; then
    read -n 1 -s -r -p "Press any key to continue"
fi