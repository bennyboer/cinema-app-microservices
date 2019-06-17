#!/bin/sh
echo "Linting user service..."
cd user
chmod +x ./lint.sh
. ./lint.sh
cd ..

echo "Linting cinema service..."
cd cinema
chmod +x ./lint.sh
. ./lint.sh
cd ..

echo "Linting movie service..."
cd movie
chmod +x ./lint.sh
. ./lint.sh
cd ..

echo "Linting presentation service..."
cd presentation
chmod +x ./lint.sh
. ./lint.sh
cd ..

echo "Linting reservation service..."
cd reservation
chmod +x ./lint.sh
. ./lint.sh
cd ..

if [[ "$1" == "debug" ]]; then
    read -n 1 -s -r -p "Press any key to continue"
fi