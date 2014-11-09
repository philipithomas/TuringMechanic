#!/bin/sh
set -e

cd go
./build.sh
cd ..

# cd site
# ./build.sh
# cd ..

# if branch master - deploy

echo "Turing Mechanic Build Complete"
