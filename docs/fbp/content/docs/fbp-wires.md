---
title: Wires
weight: 5
---

# Base concept
With FBP you mostly connect the output of a component (from a dom event) with the methods of another element.


## Connecting things

We have a lightswitch `furo-button` and a light bulb `light-bulb` which we want to toggle when the switch is clicked.

When the button fires a **click** event, the action will trigger the **toggle** method of the `light-bulb` because it is 
connected by the **wire** `--lightSwitchClicked`.


<furo-demo-snippet flow>
<template>
  <!-- This button acts as a light switch -->
  <furo-button @-click="--lightSwitchClicked" label="i am a lightswitch"></furo-button>
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>  
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
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <furo-button @-click="--lightSwitchClicked" label="i am a lightswitch"></furo-button> 
  <light-bulb ƒ-toggle="--noSource"></light-bulb>
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
<some-animation ƒ-play="--saveData"></some-animation>

<!-- good wire name -->
<some-animation ƒ-play="--saveClicked"></some-animation>
```   
  {{< /hint >}}



## Multiple sources
A wire is not limited as a point to point connection and can have multiple sources. 

<furo-demo-snippet flow style="height:400px">
<template>
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <furo-button @-click="--lightSwitchClicked" label="i am a light switch"></furo-button> 
  <furo-button @-click="--lightSwitchClicked" label="i am a light switch too"></furo-button>  
 
</template>
</furo-demo-snippet>

*It doesn't matter if you press the first or the second light switch. 
Both of them will trigger the wire`--lightSwitchClicked`, which will invoke the `toggle` method on the light-bulb.*

## Multiple sources and targets 
A wire is not limited as a point to point connection and can have multiple sources and also
have multiple targets.

<furo-demo-snippet source style="height:500px">
<template>
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <furo-button @-click="--lightSwitchClicked" label="i am a lightswitch"></furo-button> 
  <furo-button @-click="--lightSwitchClicked" label="i am a lightswitch too"></furo-button>  
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
</template>
</furo-demo-snippet>

*If you look at the source of this example, it doesn't look very complex, 
the resulting flow of this example is quite complex.*
  
  
## Triggering multiple wires from one source event
Sometimes you want to trigger multiple wires from one source event.
You can do this by separating them with a comma.
`@-click="--lightSwitchClicked, --blinkerClicked"`

<furo-demo-snippet flow style="height:500px">
<template>
  <light-bulb ƒ-toggle="--lightSwitchClicked, --intervallPulse"></light-bulb>  
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <furo-button @-click="--lightSwitchClicked, --blinkerClicked" label="i am a lightswitch"></furo-button>   
  <furo-interval-pulse ƒ-start="--blinkerClicked" ƒ-stop="--stopBlinkerClicked" @-tick="--intervallPulse" interval="500"></furo-interval-pulse>
  <furo-button danger  @-click="--stopBlinkerClicked" label="Stop the blinking"></furo-button>   
</template>
</furo-demo-snippet>

*When you press the button, it will trigger the `--lightSwitchClicked` and the `--blinkerClicked` wire.*
 
## Receiving multiple wires on a target
You can receive from multiple wires by comma separating them.

<furo-demo-snippet flow style="height:400px">
<template>
  <furo-button @-click="--lightSwitchClicked" label="i am a lightswitch"></furo-button>   
  <furo-button @-click="--blinkerClicked" label="i am a blinkswitch"></furo-button>   
  <furo-interval-pulse ƒ-start="--blinkerClicked" ƒ-stop="--stopBlinkerClicked" @-tick="--intervallPulse" interval="500"></furo-interval-pulse>
  <furo-button danger  @-click="--stopBlinkerClicked" label="Stop the blinking"></furo-button>   
  <light-bulb ƒ-toggle="--lightSwitchClicked, --intervallPulse"></light-bulb>  
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>  
</template>
</furo-demo-snippet>

*The light bulb on the top right will be triggered by the `--lightSwitchClicked` and `--intervallPulse` wire *

