---
title: Introduction
type: docs
bookToc: false
description: フロー Furo provides simply enterprise-flavoured, language-agnostic API development.
---

# フロー Furo
Furo provides simply enterprise-flavoured, language-agnostic API development.
It comes with multiple sources of truth and generates border-crossing type and service definitions.

{{< columns >}}
## Enterprise ready
Open for integrations, easy to expand, simple to automate, adapts to company processes and is technology independent.
<--->
## API first
Furo provides tools for quickly prototyping and building your API’s functionality. 
We believe in the API first architecture principle. API First obviously is in conflict with the bad practice of 
publishing API definition and asking for peer review after the service integration has started. With Furo you will
get early feedback!
{{< /columns >}}

{{< columns >}}
## Language agnostic
Define your types once and use them on any layer from the frontend to the backend.
<--->
## Easy to write
The furo interface definition language ([FIDL](https://fidl.furo.pro)) makes it easy to define your types and services.

```yaml
# This is an example µSpec for a type
- type: 'person.Client #A generic client.'
  fields:
    id: '* string:1 #This field is required (*).'
    name: '- string:2 #This one is readonly (-).'
    tags: '[] string:3 #This one is repeated ([]).'
```
```yaml
# This is an example µSpec for a service
- name: PersonService
  description: service specs for the person api
  package: personservice
  target: personservice.proto
  methods:
    - md: 'List: GET /persons type.Request, type.Response #The List method takes zero or more parameters as input, and returns a type.Response that match the input parameters.'
      qp:
        filter: 'string #Use this to define filters for the list.'
        order_by: 'string #Use this to specify the ordering.'
        page_size: 'uint32 #Set the size of the pages (elements per page).'
        page: 'uint32 #Use this field to specify the page to display.'
        q: 'string #query string.'
```
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
In addition to a set of web components, Furo Web also offers JavaScript modules for integration.
Any web framework can be connected using the JavaScript integration modules. Or simply integrate your favourite 
web component library such as Carbon Web Components, Spectrum Web Components, Momentum UI or many others.
<--->
## Web Components
Furo Web Components provides an enterprise ready set of web components which play seamlessly with Furo. 
Based on web standards and future proved. Compliant with any technology of choice. With minimal footprint it includes all enterprise standards, i18n, theming, etc
[read more](/docs/web-components/)
{{< /columns >}}
