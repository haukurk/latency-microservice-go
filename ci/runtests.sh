#!/bin/sh
export GOPATH=$HOME/gopath
export PATH=$HOME/gopath/bin:$PATH
./cmd/server/server --config config.json server & \
pid=$$!; \
go test 
