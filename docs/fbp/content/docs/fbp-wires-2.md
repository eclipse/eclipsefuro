---
title: Wire Data
weight: 8
---


# Data on wires
Wires are not limited to triggering something, they also transport information.

{{< hint info >}}
**Default**

By default the **event.detail** is passed to the target. 

  {{< /hint >}}

## Passing useful data to target
By default the **event.detail** is passed to the function you wire. 

<furo-demo-snippet demo style="height:550px">
<template>
  <furo-color-input label="choose a color"  @-value-changed="--newColor"></furo-color-input>
  <hr />
  <light-bulb ƒ-toggle="--lightSwitchClicked, --intervallPulse" ƒ-set-color="--newColor"></light-bulb>  
  <furo-button @-click="--lightSwitchClicked" label="i am a lightswitch"></furo-button>   
  <furo-button @-click="--blinkerClicked" label="i am a blinkswitch"></furo-button>   
  <furo-interval-pulse ƒ-start="--blinkerClicked" ƒ-stop="--stopBlinkerClicked" @-tick="--intervallPulse" interval="500"></furo-interval-pulse>
  <furo-button danger  @-click="--stopBlinkerClicked" label="Stop the blinking"></furo-button>   
  <light-bulb ƒ-toggle="--lightSwitchClicked, --intervallPulse" ƒ-set-color="--newColor"></light-bulb>
    
</template>
</furo-demo-snippet>
  
But sometimes you want an event as trigger and another property then event.detail as payload. 
You have several ways to accomplish this task.
 

## Send the raw *event* instead of *event.detail*

With an asterix as argument at the producer `@-event="--wireName(*)"` you will send the raw event.

## Use a sub property of the event
If the data on the wire is an object you can use a sub property.

With   `ƒ-something="--wireName(*.page.2.title)"` you will send the property title of the index 2 of the property page (which is an array in this case).


<furo-horizontal-flex>
<furo-empty-spacer>
</furo-empty-spacer>
<a href="../fbp-events/">Fireing events</a>
</furo-horizontal-flex>
