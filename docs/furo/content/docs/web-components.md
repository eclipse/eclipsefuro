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

Furo web components are a wide set of components which covers everything you need to write a web application.
They consume the same types which are defined with furo.

{{< columns >}}
## Data Integration
The transparent data agents are responsible for the communication with the APIs and the adapters for the UI interaction.
{{< mermaid >}}
graph TD
UI[UI elements]  --- agent[Data Agents]
agent-- REST ---API
{{< /mermaid >}}
<--->
## Programmable HTML
Furo FBP is like programmable HTML, no deep javascript knowledge is needed to write an application.
![viz](/viz.png)
The flowbased programming paradigm results in less complex and more flexible code. 
{{< /columns >}}

{{< columns >}}
## Material or UI5
A set of input elements which will work with the furo data structure out of the box, are available for a wide set of types.

It is no problem to write your own components, by using the data adapter.
<--->
## Layouter
{{< /columns >}}

