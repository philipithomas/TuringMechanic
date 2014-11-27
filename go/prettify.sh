#!/bin/sh
set -e
cd src
`git ls-files | grep '.go$' | xargs gofmt -l -w`
cd ..
