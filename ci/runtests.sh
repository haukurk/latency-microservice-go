#!/bin/sh
pwd
./cmd/server/server --config config.json server & \
pid=$$!; \
/home/travis/gopath/bin/go test 
