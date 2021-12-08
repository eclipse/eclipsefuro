---
title: "Overview"
weight: 5
# bookFlatSection: false
bookToc: false
# bookHidden: false
# bookCollapseSection: false
# bookComments: true
---


# Overview
Furo can produce and consume proto files. This enables you to use all the existing protoc generators to generate the 
output that fits your needs. Outputs can be documentations, boilerplate code and even applications. 

{{< mermaid >}}
graph TD

µSpec(µSpec)
Spec([Spec])
µSpec-- furo -->Spec
Spec-. furo .->µSpec


Spec-- furoc-gen-XXX  -->x[...]
Spec-- furo -->Es6Module
Spec-- furo -->Proto


Proto-- protoc-gen-grpc-gateway  -->Gateway
Proto-- protoc-gen-openapiv2  -->OpenApi
OpenApi-- swagger  -->xo[...]
Proto-- protoc  -->xp[...]
Proto -. protoc-gen-furo-specs .->µSpec

classDef green fill:#9f6,stroke:#333,stroke-width:2px;
class Spec,µSpec green
{{< /mermaid >}}

