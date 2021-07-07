---
weight: 100
title: "Maps"
---


# Maps
If you want to create an associative map as part of your data definition, 
you can define them as a field type like any other type.

{{< hint info >}}
The *key* of the map can only be a **string**.
{{< /hint >}}

### in µSpec
```yaml
fields:
  properties: 'map<string,furo.u33e.Property>:9 #Add properties you want to expose of reflect on your component.'
```

### in spec
```yaml
fields:
  properties:
    type: map<string,furo.u33e.Property>
    description: Add properties you want to expose of reflect on your component.
    __proto:
      number: 9
      oneof: ""
    __ui: null
    meta:
      default: ""
      hint: ""
      label: furo.u33e.U33eModel.properties.label
      options:
        flags: []
        list: []
      readonly: false
      repeated: false
      typespecific: null
    constraints: {}
```