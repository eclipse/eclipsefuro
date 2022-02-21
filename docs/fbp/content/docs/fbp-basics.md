---
title: Getting Started
weight: 5
---

# Getting Started
Write your applications like you make the concept for your application. 
Normaly you draw a few components and some arrows to connect the components, to describe your intention.
Why you don’t do the same to write your app or component?
Use existing web-components and wire them to build up your application.


## Furo FBP 101

In furo-FBP you connect events from one component to methods of anohter component. 
If you get this simple concept, you have understand already 90% of the things you need to write a program.



## Connecting things

We have a lightswitch `furo-button` and a light bulb `light-bulb` which we want to toggle when the switch is clicked.

When the button fires a **click** event, the action will trigger the **toggle** method of the `light-bulb` because it is 
connected by the **wire** `--lightSwitchClicked`.


<furo-demo-snippet flow>
<template>
  <!-- This button acts as a light switch -->
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>  
  <button at-click="--lightSwitchClicked">i am a lightswitch</button>
</template>
</furo-demo-snippet>


*This doesnt look very impressive in the first moment.
But as you can see, there is no scripting involved and there are no id's assigned to the components.*

{{< hint info >}}
**Hint**  
- Click on demo to see the flow in action.
- Click on source to see the corresponding source.
- Click on flow to see the resulting flow of the source.
{{< /hint >}}
 



## Multiple targets
A wire can have as many targets as you want. 

So the wire `--lightSwitchClicked` can trigger multiple targets if you want.


<furo-demo-snippet flow style="height:400px">
<template>
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <button at-click="--lightSwitchClicked">i am a lightswitch</button> 
  <light-bulb fn-toggle="--noSource"></light-bulb>
</template>
</furo-demo-snippet>

*When you press the button, all elements which are connected to the wire will trigger the defined function. In this example the last light-bulb is not connected.*


{{< hint info >}}
**Name of the wire**

You can name a wire with alphanumeric characters without a "space".

It is a good practice to name the wire by the thing that happened like `--saveClicked` or by
the values it will transport like `--responseData`. The dashes are not needed, but it makes the code easier to read too.

  {{< /hint >}}

{{< hint danger >}}
**Name of the wire**

Sometimes people tend to name the wire by the target action that the wire will trigger (`--saveData`). 
THIS IS NOT a good idea and will result in unreadable flows. 

Assume that your team want to introduce some other features, like playing a animation. 

```html
<!-- bad wire name-->
<some-animation fn-play="--saveData"></some-animation>

<!-- good wire name -->
<some-animation fn-play="--saveClicked"></some-animation>
```   
  {{< /hint >}}



## Multiple sources
A wire is not limited as a point to point connection and can have multiple sources. 

<furo-demo-snippet flow style="height:400px">
<template>
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <button at-click="--lightSwitchClicked">lightswitch 1</button>
  <button at-click="--lightSwitchClicked">lightswitch 2</button>
</template>
</furo-demo-snippet>

*It doesn't matter if you press the first or the second light switch. 
Both of them will trigger the wire`--lightSwitchClicked`, which will invoke the `toggle` method on the light-bulb.*

## Multiple sources and targets 
A wire is not limited as a point to point connection and can have multiple sources and also
have multiple targets.

<furo-demo-snippet source style="height:500px">
<template>
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <button at-click="--lightSwitchClicked">lightswitch</button>
  <button at-click="--lightSwitchClicked">lightswitch 2</button>  
</template>
</furo-demo-snippet>

*If you look at the source of this example, it doesn't look very complex, 
the resulting flow of this example is quite complex.*
  
  
## Triggering multiple wires from one source event
Sometimes you want to trigger multiple wires from one source event.
You can do this by separating them with a comma.
`at-click="--lightSwitchClicked, --blinkerClicked"`

<furo-demo-snippet flow style="height:500px">
<template>
  <light-bulb fn-toggle="--lightSwitchClicked, --intervallPulse"></light-bulb>  
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <button at-click="--lightSwitchClicked,--blinkerClicked">lightswitch</button> 
  <furo-interval-pulse fn-start="--blinkerClicked" fn-stop="--stopBlinkerClicked" at-tick="--intervallPulse" interval="500"></furo-interval-pulse> 
  <button at-click="--stopBlinkerClicked">Stop the blinking</button>
</template>
</furo-demo-snippet>

*When you press the button, it will trigger the `--lightSwitchClicked` and the `--blinkerClicked` wire.*
 
## Receiving multiple wires on a target
You can receive from multiple wires by comma separating them.

<furo-demo-snippet flow style="height:400px">
<template>
<light-bulb fn-toggle="--lightSwitchClicked, --intervallPulse"></light-bulb>  
  <light-bulb fn-toggle="--lightSwitchClicked"></light-bulb>
  <button at-click="--lightSwitchClicked">i am a lightswitch</button>   
<button at-click="--blinkerClicked">blinkswitch</button>
  <furo-interval-pulse fn-start="--blinkerClicked" fn-stop="--stopBlinkerClicked" at-tick="--intervallPulse" interval="500"></furo-interval-pulse>
<button style="color:red" at-click="--stopBlinkerClicked">Stop the blinking</button>
</template>
</furo-demo-snippet>

*The light bulb on the top right will be triggered by the `--lightSwitchClicked` and `--intervallPulse` wire *

