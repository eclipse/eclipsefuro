---
title: "Import protos"
weight: 15
---


# How to import proto files to µSpecs or Specs

We offer 2 different ways to import specs from proto. 
- `protoc-gen-furo-specs` which will generate specs 
- `protoc-gen-furo-muspecs` which will generate muspecs

The steps for any of these two generators are the same. 
This document describes the import process for muspecs.


<img src="/protologo.png" style="width: 100px;">
<span style="font-size: 80px; color:#999999">&nbsp;➔&nbsp;</span>
<span style="font-size: 50px; color:#333333">フロー µSpecs</span>



{{< mermaid >}}
graph TD
Protoc(Proto)-- protoc-gen-furo-spec --> Spec
Protoc -- protoc-gen-furo-muspec --> µSpec
µSpec(µSpec)
Spec(Spec)
µSpec-. furo .->Spec
Spec-. furo .->µSpec
Spec-. furo .->Proto
{{< /mermaid >}}


{{< hint danger >}}
**Note:** `ENUM` types and `additional bindings` are not supported at the moment


{{< /hint >}}

## Installation 

To install the latest `protoc-gen-furo-muspecs` generator, just type:

```bash
go install github.com/eclipse/eclipsefuro/protoc-gen-furo-muspecs
```

## Proto Config Options
The only option you can set, is the **exclude** option. 
It will accept a regex which will not generate the matching file names.

Add the plugin to your buf config and define the output directory.

## Define the buf template
*buf.protoimport.yaml*
```yaml
version: v1beta1
plugins:
  - name: furo-muspecs
    out: dist/muspecs    
```

## Example script to import protos from a folder

```bash
#!/usr/bin/env bash
# exit when any command fails
set -e

buf generate --template ./buf.protoimport.yaml --path $(find sourceprotos/ -type d | grep sourceprotos/[^$] | tr '\n' , | sed 's/.$//')
```


Take a look in to the [sample](https://github.com/eclipse/eclipsefuro/tree/main/protoc-gen-furo-muspecs/sample) to see a complete example.