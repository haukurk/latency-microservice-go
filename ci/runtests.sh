#!/bin/sh
pwd
./cmd/server/server --config config.json server & \
pid=$$!; \
gobin=$1
$gobin test 
