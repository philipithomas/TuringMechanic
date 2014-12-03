#!/bin/sh
set -e

export GOPATH=`pwd`
./dependencies.sh
cd src
for f in */ ; do
   cd $f
   if `ls -l *.go | wc -l`>0; then
       echo "Checking formatting in ${f}"
       count=`git ls-files | grep '.go$' | xargs gofmt -l | wc -l`
       if [ $count -gt 0 ]; then
           echo "Some files are not formatted. run prettify.sh\n"
           exit 1
       fi
       echo "Building ${f}"
       go test -v
       go install -v
   fi
   cd ..
done
