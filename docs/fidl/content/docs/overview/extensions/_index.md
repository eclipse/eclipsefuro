---
title: "Extensions"
weight: 300
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookComments: true
bookCollapseSection: true
---

# Extensionpoints

In the past, extensions were added by adding a field to a spec, using "__" as prefix. 
This was no problem, as long the complete build chain was js based, everyone could add
his extension to the spec. Adding your custom extensions is still supported, 
but they have to go to the field extensions. Extensionpoints are available for types, 
fields in types, services and methods in services. The __proto and __ui extensions were so broadly used, that 
they become a fix part of the specs.

Extensions can have any structure and can be used in furoc generator plugins or a scripts that consumes `furo exportAsYaml -f`, which 
is nearly the same structure a furoc generator would receive.


{{< hint warning >}}
Extension are only available in the standard spec notation. You can not use them in µSpecs.
{{< /hint >}}

#### Extensions in a service
```yaml
name: FruitService
extensions:
    dummy:
      corename: fruits
      gen: golang
methods:
  ListFruits:
    description: List fruits with pagination.
    extensions:
      otherextension:
        - fast
```

#### Extensions in a type
```yaml
name: Credentials
type: Credentials
description: Credentials type for login.
__proto:
  package: auth
  targetfile: auth.proto
  imports: []
  options:
    go_package: github.com/veith/doit-specs/dist/pb/auth;authpb
    java_multiple_files: "true"
    java_outer_classname: AuthProto
    java_package: com.furo.baseauth
fields:
  password:
    type: string
    description: The password.
    __proto:
      number: 1
      oneof: ""
    __ui: null
    meta:
      default: ""
      hint: ""
      label: auth.Credentials.password.label
      options:
        flags: []
        list: []
      readonly: false
      repeated: false
      typespecific: null
    constraints:
      required:
        is: "true"
        message: password is required
    extensions: #<-- field extensions
      dummyextension:
        key: a field extension
extensions: #<-- type extensions
  dummyextension:
    key: a type extension

```