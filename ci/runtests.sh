#!/bin/sh
./cmd/server/server --config config.json server & \
pid=$$!; \
go test 
