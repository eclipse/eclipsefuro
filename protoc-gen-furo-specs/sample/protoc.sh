#!/usr/bin/env bash

cd sample/protos

protoc  \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I/usr/local/include \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I$GOPATH/src/github.com/googleapis/googleapis \
-I./ \
--furo-specs_out=\
Mhelloworld.proto=../helloworld,\
:../targetDir ./helloworld/*.proto
