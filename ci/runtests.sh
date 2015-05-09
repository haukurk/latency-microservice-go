#!/bin/sh
pwd
echo $1
export $GOROOT=$2
./cmd/server/server --config config.json server & \
pid=$$!; \
$1/bin/go test 
