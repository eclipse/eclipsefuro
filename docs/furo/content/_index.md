---
title: Introduction
type: docs
bookToc: false
description: フロー Furo, the utility to work with Furo FIDLs 
---

# フロー Furo Furo
Furo furo is the main utility to work with furo specs ([FIDLs](https://fidl.furo.pro) ).

Furo furo contains helpful generators, converters, sanitizer for the furo specs.


{{< columns >}}
## Stay In Sync

The different specification formats can be used as a source or a sink or both of them. Some have a higher information 
density, furo furo can keep them in sync for you.

{{< mermaid >}}
graph LR
µSpec --> Spec
Spec --> Proto
Proto --> µSpec
Spec --> µSpec
{{< /mermaid >}}


<--->

## Multiple Sources Of Truth Supported
You have the choice from multiple definition formats as your source of truth.
This can be *.proto or one of the FIDLs (µSpec, spec). 

[Read more here](/docs/sourceoftruth/), to find out what fits best for you.

{{< /columns >}}

{{< columns >}}
## Extendable

Furo FIDLs comes with a lot of extension points. You can add your extensions at type level, field level and services. 
The extensions can be used by *your* generators or scripts.
<--->
## Configurable Flows
The commands of furo can be configured as **flows** (chain of commands). You can also add your custom commands and 
add them at any point to your flow. 

**flow build**
{{< mermaid >}}
graph LR
µSpec2Spec --> genMessageProtos
genMessageProtos --> Protoc(protoc -I...)
Protoc --> Publish(publish)
{{< /mermaid >}}

{{< /columns >}}