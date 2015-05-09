#!/bin/sh
pwd
./cmd/server/server --config config.json server & \
pid=$$!; \
$1 test 
