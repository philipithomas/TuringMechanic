#!/bin/sh
set -e

# Golang unit tests and compiling
cd go
./build.sh
cd ..

# Jekyll building
cd site
./build.sh
cd ..

echo "Turing Mechanic Build Complete"
