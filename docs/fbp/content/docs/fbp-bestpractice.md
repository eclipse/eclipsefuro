---
date: 2017-04-17T14:02:04+02:00
title: Best Practice
weight: 60
---
# Best Practice
The best practices are a result from the feedback of our users, thanks to them.

Feel free to give us your feedback.
 
## Wire and event naming
 In fact you can use any string to name wires, but most of the users are using two dashes in front of the wire name and for the name they use camelCase notation.
 
For the events they use the dashed-case notation, because you can not write `at-camelCase`.

```html
<furo-button at-click="--buttonClicked, ^^fired-event"></furo-button>
```
With this notation they can see the difference between a **--wire** to an **event** they fire directly.


## Use declarative wire names, don't be imperative

When you use declarative names, it would be easier to read and modify a wired program.

**bad example**
```html
<furo-button at-click="--closeView"></furo-button>
<my-view fn-close="--closeView"></my-view>
<data-component fn-save="--closeView"></data-component>
```

**good example**
```html
<furo-button at-click="--closeButtonClicked"></furo-button>
<my-view fn-close="--closeButtonClicked"></my-view>
<data-component fn-save="--closeButtonClicked"></data-component>
```

*It is a subtile but relevant difference between this two examles.*

## Use *event delegation*
When you use a set of components, you don't have to wire every single component to the appropriate target. Use *event delegation* whenever possible. It is faster and easier to read.

 

 
**without event delegation**

```html
    <div >
      <mole-hole key="a" at-closed="--moleClosed" at-continue="--continue" at-miss="--missed" at-whack="--whacked"></mole-hole>
      <mole-hole key="s" at-closed="--moleClosed" at-continue="--continue" at-miss="--missed" at-whack="--whacked"></mole-hole>
      <mole-hole key="d" at-closed="--moleClosed" at-continue="--continue" at-miss="--missed" at-whack="--whacked"></mole-hole>
      <mole-hole key="f" at-closed="--moleClosed" at-continue="--continue" at-miss="--missed" at-whack="--whacked"></mole-hole>
      <mole-hole key="g" at-closed="--moleClosed" at-continue="--continue" at-miss="--missed" at-whack="--whacked"></mole-hole>
      <mole-hole key="w" at-closed="--moleClosed" at-continue="--continue" at-miss="--missed" at-whack="--whacked"></mole-hole>
    </div>
```
![without event delegation](/withoutEventDelegation.png)


**with event delegation**
```html
    <div at-closed="--moleClosed" at-continue="--continue" at-miss="--missed" at-whack="--whacked">
      <mole-hole key="a"></mole-hole>
      <mole-hole key="s"></mole-hole>
      <mole-hole key="d"></mole-hole>
      <mole-hole key="f"></mole-hole>
      <mole-hole key="g"></mole-hole>
      <mole-hole key="w"></mole-hole>
    </div>
```
![with event delegation](/eventDelegation.png)



