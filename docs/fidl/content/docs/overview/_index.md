---
weight: 1
title: "Overview"

---

# Overview

This guide describes how to use the specs, including *.service.spec, *.type.spec file syntax.

{{< hint info >}}
We will use µSpec notation when it makes no difference or it is easier to use.

The standard notation is then used when the µSpec notation does not cover the feature.

For a better understanding and when possible the **µSpec**, **spec** and the resulting **proto** will be shown.
{{< /hint >}}

This is a reference guide – for a step by step example, see the tutorials or sample projects.

The pages [Defining a type](/docs/overview/define_type/) , [Defining a service](/docs/overview/define_type/) and [Style Guide](/docs/overview/style_guide/) should give you enough information for a brief overview. 
Most of the other pages are detailed topics and can be read later or by interest or need.

If you have specified additional commands and flows in your furo config, make sure you install the corresponding dependencies too.


## What is Generated From Your specs?

### protos
First of all, **protos** and therfore [all you can generate with protos](https://developers.google.com/protocol-buffers/docs/proto3#whats_generated_from_your_proto).

This is done with the following commands.

- `furo genMessageProtos` - Generate the message protos from the type specs.
- `furo genServiceProtos` - Generate service protos from the specs

### client environment (es6)
The client types as es6 module which can be consumed by [@furo/furo-data*](https://components.furo.pro/?t=furo-data) and various web components of furo.
This module allows the usage of the same types on the backend side and in the browser.

- `furo genEsModule` - generate es6 spec module

### Validators, DB Shemas, Custom Documentation,...
With [furoc](https://github.com/theNorstroem/furoc), which has a lot of similarities with protoc, you can easyli write your 
custom generators. The main benefit is that you have a much higher information density then you have with the protos alone (to be fair, you can have the same information density with protos too, but this is not so trivial).
The input format for a furoc generator is a yaml structure with your services and types. The output format is the same like in protoc. 

### User Interface Components
Yes, you have read it correctly. You can generate web-components that you can use in your web projects. At the moment you can
do that with [@furo/ui-builder](https://github.com/theNorstroem/FuroBaseComponents/tree/master/packages/furo-ui-builder).
A [furoc](https://github.com/theNorstroem/furoc) based version is in development and comming soon.

### REST APIs *indirect*
This is done by using the protos with [protoc-gen-grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

```bash
 protoc -I . --grpc-gateway_out ./gen/go \
     --grpc-gateway_opt logtostderr=true \
     --grpc-gateway_opt paths=source_relative \
     --grpc-gateway_opt grpc_api_configuration=path/to/config.yaml \
     your/service/v1/your_service.proto
```

### Open Api Specifications *indirect*
[Open Api aka swagger](https://swagger.io/) can be used for various things.

Generating Open Api Specifications is done by using the generated protos with [protoc-gen-openapiv2](https://github.com/grpc-ecosystem/grpc-gateway).

```bash
protoc -I . --openapiv2_out ./gen/openapiv2 --openapiv2_opt logtostderr=true your/service/v1/your_service.proto
```

### Various Clients and Servers *indirect*
By using the swagger files on [https://editor.swagger.io/](https://editor.swagger.io/) you can generate server and client
code for different languages and architectures. You do not have to write the backend as grpc service to work with the furo client framework.


## Importing proto Messages
You can import your existing proto messages with the [protoc-gen-furo-specs](https://github.com/theNorstroem/protoc-gen-furo-specs) protoc plugin.
After running Furo, you should receive the same proto file (proto3).

## furoc
[Furoc](https://github.com/theNorstroem/furoc) is the compiler/transpiler for the specs.

Furoc will pass a yaml structure with the current config of the spec project, the types, the services, the installed types (dependencies) and the installed services (dependencies) to the generators.


## Related Documents

- [Anatomy of a µType spec](/docs/µSpecs/types/)
- [Anatomy of a µService spec](/docs/µSpecs/services/)
- [Anatomy of a type spec](/docs/specs/types/)
- [Anatomy of a service spec](/docs/specs/services/)

## Good to Read

### [API Design Guide from google
The [API Design Guide](https://cloud.google.com/apis/design) from google gives you a good guideline for designing your APIs.

### Protocol Buffers
This page gives you a good overview on [Protocol Buffers](https://developers.google.com/protocol-buffers).

### gRPC
This site [grpc.io](https://grpc.io/) gives you a good entry point to grpc itself.