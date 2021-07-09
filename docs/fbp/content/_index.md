---
title: Furo FBP
type: docs
bookToc: false
description: フロー Furo FBP
---

# フロー Furo Flow Based Programming
enables you to write your web applications following [the flow based programming paradigm](https://en.wikipedia.org/wiki/Flow-based_programming).



{{< columns >}}
## Fully Declarative - No "Code" needed

Instead of writing hundreds of lines of code with HTML element selectors and attaching EventListeners to them to 
write another HTML element selector for calling a simple method, simply express your intention and string them together.

<furo-demo-snippet flow style="height: 150px">
<template>
  <!-- This button acts as a light switch -->
  <furo-button @-click="--lightSwitchClicked" label="i am a lightswitch"></furo-button>
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>  
</template>
</furo-demo-snippet>

<--->
## Visual application analysis
Inspect your application logic visually.

![viz](viz.png)

*currently in closed beta, will be opened soon*
{{< /columns >}}

{{< columns >}}
## Web components best friend
With this type of programming, you can add or remove different components 
and chain them together to get different results or functionalities.

And the visual representation makes it easier for a wide range of players to build applications without having to write multiple lines of code.
<--->
## Coexistence of flows and code - <small>at any time</small>

There is **no**  FBP only, furo-FBP comes with the bridges (hooks and triggers)
to solve the problems in the place where it makes the most sense.  

## Framework agnostic
Furo-FBP works anywhere you use HTML, with any framework or none at all.

{{< /columns >}}



{{< columns >}}
## Reduced complexity
UIs are far too complex to write and maintain them imperatively. 
The same methodology is used by some of the most successful game engines to cope with complexity.

The flow based approach is different, but the result is similar to when a programmer 
tells the machine what he wants to do, and the computer takes those instructions 
and executes them just as it would with text-based code.

<--->
## Reduced development time
In 1994 Morrison published a book describing FBP, and providing **empirical evidence** that FBP led to **reduced development times** ([Advances in Dataflow Programming Languages](https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.99.7265)).

{{< /columns >}}
