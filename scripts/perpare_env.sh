#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

projects_array=( "artifacts" "deployment-producer" "executor" \
"migration-producer" "test-infra" "verification-producer" \
"ashandler" "digester" "infra-producer" "report-producer" \
"tfhandler" "xplaceholder")

for i in "${projects_array[@]}"
do
  go get github.com/xplaceholder/$i
  pushd $GOPATH/src/github.com/xplaceholder/$i
    git remote remove origin
    git remote add origin git@github.com:xplaceholder/$i.git
  popd
done
