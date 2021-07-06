---
title: "Overview"
weight: 10
# bookFlatSection: false
bookToc: false
# bookHidden: false
# bookCollapseSection: false
# bookComments: true
---


# Overview
A brief overview over the main tasks where furo / furoc / protoc[^1] are used for.

We have seen some nice generators like *furoc-gen-ddl*[^3], which uses simple type and field extensions to generate mysql DDLs. 

A documentation generator *furoc-gen-apidocs*[^2] is in development.

{{< mermaid >}}
graph TD

µSpec(µSpec)
Spec([Spec])
µSpec-- furo -->Spec
Spec-. furo .->µSpec
someProto[Proto] -- protoc-gen-furo-specs*1 --> µSpec

Spec-- furo -->Es6Module
Spec-- furoc-gen-apidocs -->Docs
Spec-- furoc-gen-u33e  -->U33E
U33E-- simple-generator  -->web-components
Spec-. furoc-gen-ddl .->ddl[(DDL)]
Spec-- furo  -->xs[...]
Spec-- furo -->Proto


Proto-- protoc-gen-grpc-gateway  -->Gateway
Proto-- protoc-gen-openapiv2  -->OpenApi
OpenApi-- swagger  -->xo[...]
Proto-- protoc  -->xp[...]

{{< /mermaid >}}

[^1]: Protos are usualy a product of furo, but they can also be a source for the µSpecs (protoc-gen-furo-specs). 

[^2]:  not released yet

[^3]: sadly not open source :-(
