---
weight: 90
title: "Oneof"
---
# Oneof
If you have a message with many optional fields and where at most one field will be set at the same time.


In the client lib, at most one field can be set at the same time. 
Setting any member of the oneof automatically clears all the other members. 



## Using Oneof
Just define a name for the oneof.

*field section of a type µSpec*
```yaml
  fields:
    method: '* string:1 [handler] #The name of the method to call with the wire data. If you want add custom code use source instead of method. /oneof:handler/'
    source: '* string:2 [handler] #Anonyous method to handle the wire. Prefer the use of method. /oneof:handler/'
    
```

*field section of a type spec*
```yaml
fields:
  method:
    type: string
    description: The name of the method to call with the wire data. If you want add custom code use source instead of method. /oneof:handler/
    __proto:
      number: 3
      oneof: handler
    __ui: null
    meta:
      default: ""
      hint: ""
      label: furo.u33e.WireHook.method.label
      options:
        flags: []
        list: []
      readonly: false
      repeated: false
      typespecific: null
    constraints: {}
  source:
    type: string
    description: Anonyous method to handle the wire. Prefer the use of method. /oneof:handler/
    __proto:
      number: 4
      oneof: handler
    __ui: null
    meta:
      default: ""
      hint: ""
      label: furo.u33e.WireHook.source.label
      options:
        flags: []
        list: []
      readonly: false
      repeated: false
      typespecific: null
    constraints: {}

```

The resulting proto would be like:

*proto*
```proto

// Wire hooks to connect internal wires with methods.
message WireHook {  

    // Short description what you will do when this wire was triggered
    string description = 1;

    // The wire to hook on.
    string wire = 2;

    // Registers the hook as first receiver of the wire.
    bool hookBefore = 5;
    oneof handler {

        // The name of the method to call with the wire data. If you want add custom code use source instead of method. /oneof:handler/
        string method = 3;

        // Anonyous method to handle the wire. Prefer the use of method. /oneof:handler/
        string source = 4;
    }
}
```