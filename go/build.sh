#!/bin/sh
set -e

export GOPATH=`pwd`
cd src
for f in */ ; do
   cd $f
   echo "Building ${f}"
   go test -v
   go install -v
   echo "Done"
   cd ..
done
