---
title: "furo web"
weight: 10
# bookFlatSection: false
bookToc: false
bookHidden: true
# bookCollapseSection: false
# bookComments: true
---

# フロー Furo Web Components

Furo Web Components provides an enterprise ready set of web components which play seamlessly with Furo. 
Based on web standards and future proved. Compliant with any technology of choice. 
With minimal footprint it includes all enterprise standards, i18n, theming and much more.

The furo web components are a wide set of components which covers everything you need to write a web application.
They consume the same types which are defined with furo.

{{< columns >}}
## [Data Integration](https://components.furo.pro/?t=furo-data)
The transparent data agents are responsible for the communication with the APIs and the adapters for the UI interaction.
{{< mermaid >}}
graph TD
UI[UI elements]-- HTML ---agent[Data Agents]
agent-- REST ---API
{{< /mermaid >}}
<--->

## Programmable HTML
Furo FBP is like programmable HTML, no deep javascript knowledge is needed to write an application.
![viz](/viz.png)
The flowbased programming paradigm results in less complex and more flexible code. 
{{< /columns >}}

{{< columns >}}
## [Material](https://components.furo.pro/?t=furo-data-input) or [SAP UI5](https://components.furo.pro/?t=furo-ui5)
A set of input elements which will work with the furo data structure out of the box, are available for a wide set of types.

It is no problem to write your own components, by using the data adapter.
<--->
## Layouter
Ready to use layouters for [forms](https://components.furo.pro/?t=demo-FuroFormLayouter-0), [grids](https://components.furo.pro/?t=demo-FuroZGrid-0) and [more](https://components.furo.pro/?t=furo-layout)
{{< /columns >}}
