#!/bin/sh
pwd
echo $1
export GOROOT=$1
export GOPATH=/home/travis/gopath/
./cmd/server/server --config config.json server & \
pid=$$!; \
$1/bin/go test 
