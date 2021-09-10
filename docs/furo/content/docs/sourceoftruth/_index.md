---
weight: 40
title: "Source of Truth"
# bookCollapseSection: true
---

# Source of Truth
The different specification formats can be used as a source or used a sink (receiver) or can be both of them.

{{< mermaid >}}
graph LR
µSpec --> Spec
Spec --> Proto
Proto --> µSpec
Spec --> µSpec
{{< /mermaid >}}


You should choose **one** source of truth and stick to it when possible. *But it is always possible to change your decision.*

### Sources AND Sink
- Spec
- µSpec
- Proto

### Sinks only
- Es6Module (by design)
- Docs
- U33E
- Gateway
- OpenApi


## Spec as source
{{< mermaid >}}
graph TD
Spec --> Proto
Spec --> µSpec
Spec --> Es6Module
Spec --> Docs
Spec --> U33E
Spec --> ...
Proto --> Gateway
Proto --> OpenApi
{{< /mermaid >}}

{{< columns >}}
### Pros
- only one kind of spec to work with
- run your flow at any time with same results (idempotency)
- guaranteed interoperability from backend to client
<--->
### Cons
- design phase is harder
- µSpec only for discussions and not for "design"
{{< /columns >}}

{{< hint warning >}}
Recommended when you only have to maintain an existing project.
{{< /hint >}}


## µSpec as source
{{< mermaid >}}
graph TD
Spec --> Proto
µSpec -->Spec
Spec --> Es6Module
Spec --> Docs
Spec --> U33E
Spec --> ...
Proto --> Gateway
Proto --> OpenApi
{{< /mermaid >}}

{{< columns >}}
### Pros
- fastest variant to design new stuff
- simplest notation
- covers >90% of the cases
- idempotent
- guaranteed interoperability from backend to client
<--->
### Cons
- working in specs for some edge cases required

{{< /columns >}}


{{< hint warning >}}
Recommended when you have a fresh project or a lot of changes.

Easiest method when you know protobuf/grpc. 
{{< /hint >}}


### Proto as source
{{< mermaid >}}
graph TD
Proto --> µSpec
µSpec -->Spec
Spec --> Es6Module
Spec --> Docs
Spec --> U33E
Spec --> ...
Proto --> Gateway
Proto --> OpenApi
{{< /mermaid >}}

{{< columns >}}
### Pros
- good solution when **everything** is already defined in proto
- covers >80% of the cases
- idempotent

  <--->
### Cons
- destructive step Proto --> µSpec
- working in specs for some cases required
- writing REST service definitions by "hand"  
- interoperability from backend to client is not guaranteed because some steps are not under control of furo anymore.
{{< /columns >}}

{{< hint warning >}}
Recommended when you have a huge portfolio of protos and want to bring them to the web (without any changes).
{{< /hint >}}


## Advanced Setup
This is a extended µSpec as source variant.

{{< mermaid >}}
graph TD
SomeProto --> µSpec
µSpec --> Spec
Spec --> Proto
Spec --> Es6Module
Spec --> Docs
Spec --> U33E
Spec --> ...
Proto --> Gateway
Proto --> OpenApi
{{< /mermaid >}}

{{< columns >}}
### Pros
- implement external changes very fast
- designing new stuff still fast
- guaranteed interoperability from backend to client

<--->
### Cons
- hard setup for the SomeProto part
- bigger skill set needed
- partially destructive step
{{< /columns >}}


{{< hint warning >}}
Recommended when you have to import protos to your system. Use "µSpec as source" first and switch to this variant when 
you have to.
{{< /hint >}}

## Multiple Sources of Truth
The you know what you do mode.


{{< hint danger >}}
When you have to migrate a project from different spec formats, you have to use this setup. This is **not recomended** and
should still be an exception and not the default.
{{< /hint >}}


{{< mermaid >}}
graph TD
SomeProto --> µSpec
OtherFormats --> Spec
µSpec --> Spec
Spec --> Proto
Spec --> µSpec
Spec --> Es6Module
Spec --> Docs
Spec --> U33E
Spec --> ...
Proto --> Gateway
Proto --> OpenApi
Proto --> µSpec

{{< /mermaid >}}

{{< columns >}}
### Pros
- Do what ever you ever want.
<--->
### Cons
- Dangerous
- More then one flow required
- Not idempotent
- very hard setup

{{< /columns >}}


{{< hint danger >}}
**NOT Recommended!** 

Do this only when you have to do, try to switch to another setup as fast as you can.
{{< /hint >}}

