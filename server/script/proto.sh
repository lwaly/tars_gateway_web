#! /bin/bash
cd $GOPATH/src/github.com/lwaly/tars_gateway_web/server/protocol
protoc --go_out=plugins=tarsrpc:. *.proto
cd $GOPATH/src/github.com/lwaly/tars_gateway_web/server/models
protoc --go_out=plugins=tarsrpc:. *.proto

