---
title: Compatibility
weight: 20
---
# Working with suboptimal components

Some components that you may want or have to use are not dispatching events or receiving data by methods.
Therefore, furo FBP has some features to work with them too.


## Setting a member value
To set a value of a component you use can be done with the **set-property**.

<furo-demo-snippet  flow style="height:150px">
<template>
  <button at-click="--bntClicked">A</button>
  <span set-inner-text="--bntClicked">click counter</span>
</template>
</furo-demo-snippet>

{{< hint info >}}
**Note, a property is not an attribute.**

You have to know what you do. This is a direct manipulation of a component and maybe there is some reason
why it does not expose the property.
 
{{< /hint >}}

## Wireing responses from method calls

When the method that you have wired only returns data that you want to use, 
you can wire the response with **at-fnret-methodname**.


<furo-demo-snippet no-demo flow style="height:200px">
<template>
   <!-- we put the value of number on the wire --calcClicked -->
   <furo-button at-click="--calcClicked(number)"> calculate sqrt </furo-button>
   <!-- The response of the calculate method is dispatched on at-fnret-calculate -->    
   <square-root fn-calculate="--calcClicked" at-fnret-calculate="--calculatedSqrRoot"></square-root>
  <display-result fn-show="--calculatedSqrRoot"></display-result>
</template>
</furo-demo-snippet>

*The response of the `calculate(n)` method is avaliable on the wire `--calculatedSqrRoot`.*

## Spread arguments
When a receiver mehtod accepts multiple arguments or is a spread operator and the data on the wire is spreadable, 
furo FBP will handle this for you correctly.



