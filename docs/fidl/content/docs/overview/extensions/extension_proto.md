---
weight: 200
title: "Protobuf Extension"
---

# The Protobuf Extension

The protobuf extension is meanwhile a fix part of the specs and Furo.  

In this, meanwhile built in, extesnion you define things for the protos.


#### __proto extension in a type
The proto extension in a type defines the package, target file, imports (*.proto) and options

```yaml
__proto:
    package: auth
    targetfile: auth.proto
    imports: []
    options:
        go_package: github.com/veith/doit-specs/dist/pb/auth;authpb
        java_multiple_files: "true"
        java_outer_classname: AuthProto
        java_package: com.furo.baseauth
```


#### __proto extension in a field of a type
The proto extension in a field let you define the field id an can set a oneof group.

```yaml
fields:
  id:
    type: string
    description: The identifier.
    __proto:
        number: 1
        oneof: ""
    __ui: null
```


#### __proto extension in a service
The proto extension in a service defines the package, target file, imports (*.proto) and options

```yaml
__proto:
  package: Services.auth
  targetfile: auth.proto
  imports:
    - google/api/annotations.proto
    - Services/auth/reqmsgs.proto
    - google/protobuf/empty.proto
    - auth/auth.proto
  options:
    go_package: github.com/veith/doit-specs/dist/pb/Services/auth;authpb
    java_multiple_files: "true"
    java_outer_classname: AuthProto
    java_package: com.furo.baseServices.auth
```