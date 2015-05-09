#!/bin/sh
export GOROOT=$1
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
./cmd/server/server --config config.json server & \
pid=$$!; \
go test 
