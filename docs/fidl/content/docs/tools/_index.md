---
weight: 30
title: "Tools"
date: 2020-11-11T19:37:46+01:00
bookCollapseSection: true
---
# Useful tools to work with the furo FIDLs
There are a lot of tools, you have at your hand at your fingertips. If you do not want to install them (a lot of work), you can
simply use the docker image [thenorstroem/furo-bec](/docs/tools/BEC/). It contains erery tool you need.


## [protoc](https://github.com/protocolbuffers/protobuf)
The protocol buffer compiler, protoc, is used to compile .proto files, which contain service and message definitions. 


## [protoc-gen-furo-specs](https://github.com/theNorstroem/protoc-gen-furo-specs)
Protoc plugin to generate furo µSpecs from proto files.

## [protoc-gen-grpc-gateway (v2)](https://grpc-ecosystem.github.io/grpc-gateway/#getting-started)
gRPC-Gateway is a plugin of protoc. It reads a gRPC service definition, and generates a reverse-proxy server which translates a RESTful JSON API into gRPC. This server is generated according to custom options in your gRPC definition.


## [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/)
Generate the go stubs.
## [protoc-gen-go-grpc](https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc)
This tool generates Go language bindings of services in protobuf definition files for gRPC.

## [simple-generator](https://github.com/theNorstroem/simple-generator)
Very simple golang template engine for cli usage, using the golang template package and the template functions from sprig.

## [furoc](https://github.com/eclipse/eclipsefuro)
The furoc "compiler" is a generator tool similar to protoc.
There are no official plugins available at the moment. Furoc can be used to write your custom generators.

