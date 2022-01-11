---
title: "Import protos"
weight: 20
---


# How to import proto files to µSpecs or Specs

We offer 2 different ways to import specs from proto. 
- `protoc-gen-furo-specs` which will generate specs 
- `protoc-gen-furo-muspecs` which will generate muspecs

hint:
ENUM not supported

hint: 
additional bindings

<img src="/grpcio-ar21.svg" style="width: 120px;">
<span style="font-size: 80px; color:#999999">➔</span>
<span style="font-size: 50px; color:#333333">フロー µSpecs</span>

## Installation 
`protoc-gen-furo-specs`

```bash
go install github.com/eclipse/eclipsefuro/protoc-gen-furo-specs
```

`protoc-gen-furo-muspecs`

```bash
go install github.com/eclipse/eclipsefuro/protoc-gen-furo-muspecs
```


## Config
There is nothing to configure. 

*buf.gen.yaml*
```yaml
version: v1beta1
plugins:
  - name: furo-muspecs
    out: dist/muspecs
    opt: paths=source_relative
```