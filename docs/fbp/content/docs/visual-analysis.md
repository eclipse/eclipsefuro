---
title: Visual analysis
weight: 170
---

# Visual Application Analysis
Inspect your application logic visually by using viz.

## Starting a viz.furo.pro session
### Start from **body**
To start a viz session just type `viz()` in your browser console.
This will open and connect your current browser tab to viz.furo.pro where you can see your application flow.
By a double click on a component, viz will request the component from your application and render it.

### Start from a particular component
Instead of clicking down until you have reached the component you are interested in, you can start viz with a reference to your component.

* example with chrome*
- Use the picker to select an element in the page
- it will be highlighted in the element view and a reference is available with **$0**
- start viz with a reference to the element by typing `viz($0)`

### Examples

![viz-from-root](/viz-from-root.png#max) 
*viz started from root*

![viz-from-root](/app-shell.png#max) 
*component: app-shell*

![viz-from-root](/unconnected.png#max) 
*A wire which has no connection to target will be displayed as a red circle*


### How to read the flow graph
If you are familiar with fbp, you should not have any problem to read the graph. 

* The **boxes** represent the used components.
![viz-from-root](/component.png)
*The name of the component is on the top left (furo-app-flow)*
   * Boxes with dashed lines have a comment in the source. Hover on the box to read the comment


* The **blue lines** are the wires. 
  * Hover on them to read the wire name, like `--unauthorized`.

* The **small blue boxes** with an **at-** are the catched events. 
  * Hover on them to read the used name and more.

* The **small green boxes** with an **fn-** are the triggerer for the methods of the component.

* The **small black boxes** are attributes without a value assignment.
  * These are often boolean flags like `hidden` or `readonly` which are setted.

* The **small orange boxes** are string attributes of the component which are setted. 
  * Hover on them to read the setted string.

* The **orange dots** are indicating a wire from nowhere or a wire which was triggered from the source (like this._FBPTriggerWire("--dataReceived",data)) or from outside (like --pageEntered from furo-pages).
  * If you trigger a wire from the source and use the prefix `|--` on the wire, the dot will turn **green**

* The **red dots** are indicating a wire which goes nowhere or a wire which is cathced in the source (like this._FBPAddWireHook("--wireName",(e)=>{ ... });

#### Keyboard shortcuts
* **f** on the buttons toggles the fullscreen mode. Press "esc" to get back.
* **ctrl v** or **cmd v** renders the clipboard content. Do not forget to allow your browser to accept the clipboard content.
* **arrow-left**, **arrow-right** ◀, ▶ re renders the last pasted content.
* **Backspace** removes the current view.

#### Mouse controls
* **scroll down** zooms the flow in.
* **scroll up** zooms  out.
* **moving the mouse with mousedown** pans the flow.

#### Touch controls

* pinch in zooms the flow in.
* pinch out zooms the flow out.
* paning (with 2 fingers) pans the flow.
