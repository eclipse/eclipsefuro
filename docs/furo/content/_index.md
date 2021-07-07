---
title: Introduction
type: docs
bookToc: false
description: フロー Furo, the utility to work with Furo FIDLs 
---

# フロー Furo
Furo is the main utility to work with furo specs ([FIDLs](https://fidl.furo.pro) ).

Furo contains helpful generators, converters, sanitizer for the furo specs.


{{< columns >}}
## Enterprise ready
<--->
## API first
{{< /columns >}}

{{< columns >}}
## Language agnostic
<--->
## Easy to write
{{< /columns >}}


{{< columns >}}
## Stay In Sync

The different specification formats can be used as a source or a sink or both of them. Some have a higher information 
density, furo can keep them in sync for you.

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
Furo comes with proto as default output, but is not limited to proto. With furoc you can write your own output formats.
<--->
## Configurable Flows
The commands of furo can be configured as **flows** (chain of commands). You can also add your custom commands and 
add them at any point to your flow. 

**furo run build**
{{< mermaid >}}
graph LR
µSpec2Spec --> genMessageProtos
genMessageProtos --> Protoc(protoc -I...)
Protoc --> Publish(publish)
{{< /mermaid >}}

{{< /columns >}}


{{< columns >}}
## Browser integration
<--->
## web-components
{{< /columns >}}
