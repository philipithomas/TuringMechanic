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

# TODO - if branch master - deploy

echo "Turing Mechanic Build Complete"
