---
title: Introduction
type: docs
bookToc: false
description: フロー Furo FIDL, the **F**uro **I**nterface **D**efinition **L**anguage 
---

# フロー Furo FIDL
The **F**uro **I**nterface **D**efinition **L**anguage

{{< columns >}}
## What are FIDLs?
```yaml
# This is a µSpec for a type
- type: 'person.Client (ce) #A generic client.'
  fields:
    id: '*string:1 # The id is required (*).'
    display_name: '-string:2 # The display_name is readonly (-).'
    tags: '[]string:3 # Define some tags ([])'
    living_address: 'person.Adress:4 # The living address.'

```

Furo IDL is a interface definition language with a lot of output formats. 
The outputs can vary from [protobuf](https://developers.google.com/protocol-buffers), UI components, open api, client code, server code, 
documentations and much more. 
There is a short notation form [(µSpecs)](/docs/µSpecs/) and the standard notation form [(specs)](/docs/specs/). You can use the **same** defined types and services on servers and browsers.
<--->

## How do i start?
```yaml
# This is a µSpec for a service
- name: Client
  description: "List clients"
  package: Services.clientservice
  target: clientservice.proto
  services:
    - md: 'Create: POST /clients person.Client , person.ClientEntity # Add a new client.'      
    - md: 'Delete: DELETE /clients/{cid} google.protobuf.Empty, google.protobuf.Empty # Delete a client'
      qp:
        cid: 'string #The query param cid stands for client id.'
```
- look at some example projects
- read the [overview]()
- install the [tools]()

{{< /columns >}}


## Why are there two spec notation formats?
The µ (micro) spec is ideal for rapid prototyping or simple APIs, Furo can translate them to regular
specs and vice versa. They do not have any extension points and are as minimalistic as possible. The µSpecs have a lower information density then the regular specs.

The regular specs are more expressive. You can also extend them with additional information, which you may need in a custom furoc generator to match your needs. 
 
When you already know protobuf, you will find some similarities in the specs.


