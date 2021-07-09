---
title: Compatibility
weight: 500
---
# Working with sub optimal components

Some components that you may want or have to use are not dispatching events.
Therefore, furo FBP has some features to work with them too.

## Wireing responses from method calls

When the method that you have wired only returns data that you want to use, 
you can wire the response with **@-ƒ-methodname**.


<furo-demo-snippet no-demo flow style="height:200px">
<template>
   <!-- we put the value of number on the wire --calcClicked -->
   <furo-button @-click="--calcClicked(number)"> calculate sqrt </furo-button>
   <!-- The response of the calculate method is dispatched on @-ƒ-calculate -->    
   <square-root ƒ-calculate="--calcClicked" @-ƒ-calculate="--calculatedSqrRoot"></square-root>
  <display-result ƒ-show="--calculatedSqrRoot"></display-result>
</template>
</furo-demo-snippet>

*The response of the `calculate(n)` method is avaliable on the wire `--calculatedSqrRoot`.*

## Spread arguments
When a receiver mehtod accepts multiple arguments or is a spread operator and the data on the wire is spreadable, 
furo FBP will handle this for you correctly.



