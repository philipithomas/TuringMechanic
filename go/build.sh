#!/bin/sh
set -e

export GOPATH=`pwd`
./dependencies.sh
cd src
for f in */ ; do
   cd $f
   if [ `git ls-files * | wc -l` -gt 0 ]; then
       echo "Checking formatting in ${f}"
       count=`git ls-files | grep '.go$' | xargs gofmt -l -s | wc -l`
       if [ $count -gt 0 ]; then
           echo "Some files are not formatted. run prettify.sh\n"
           exit 1
       fi
       go vet .
       echo "Building ${f}"
       go test -race -v
       go install -race -v
   fi
   cd ..
done
