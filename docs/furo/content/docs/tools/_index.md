---
weight: 30
title: "Additional Tools"
date: 2020-11-11T19:37:46+01:00
bookCollapseSection: true
---
# Useful tools to work with the furo FIDLs
There are a lot of tools, you have at your hand at your fingertips. If you do not want to install them (a lot of work), you can
simply use the docker image [thenorstroem/furo-bec](./BEC/). It contains erery tool you need.

## [furo](https://github.com/theNorstroem/furo/blob/master/doc/furo.md)
The furo is the essential tool to work with furoc-FIDLs, it replaces the @furo/spec npm package.

#### Installation
```bash
brew tap theNorstroem/tap
brew install furo
```

## [@furo/ui-builder](https://github.com/theNorstroem/FuroBaseComponents/tree/master/packages/furo-ui-builder)
The @furo/ui-builder is the tool to build web-components from your specs. In the near future it will be replaced with
**furoc-gen-u33e**.

## [protoc](https://github.com/protocolbuffers/protobuf)
The protocol buffer compiler, protoc, is used to compile .proto files, which contain service and message definitions. 


## [protoc-gen-furo-specs](https://github.com/theNorstroem/protoc-gen-furo-specs)
Protoc plugin to generate furo specs from proto file.

## [protoc-gen-grpc-gateway (v2)](https://grpc-ecosystem.github.io/grpc-gateway/#getting-started)
gRPC-Gateway is a plugin of protoc. It reads a gRPC service definition, and generates a reverse-proxy server which translates a RESTful JSON API into gRPC. This server is generated according to custom options in your gRPC definition.


## [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/)
Generate the go stubs.
## [protoc-gen-go-grpc](https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc)
This tool generates Go language bindings of services in protobuf definition files for gRPC.

## [simple-generator](https://github.com/theNorstroem/simple-generator)
Very simple golang template engine for cli usage, using the golang template package and the template functions from sprig.

## [furoc](https://github.com/theNorstroem/furoc)
The furoc "compiler" is a generator tool similar to protoc.
There are no official plugins available at the moment. But furoc-gen-u33e will come in near future.

### furoc-gen-xxx
Is a placeholder for all furoc plugins that will come.