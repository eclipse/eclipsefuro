---
title: Custom Events
weight: 30
---
# Fireing custom events

Web Components use events to communicate state changes up the DOM tree to parent elements. 

Furo also provides a notation for events which allows you to specify events in a declarative manner.
This is very useful when you want to trigger a event with a more specific name then the originating event has.

*On the first view, it does not make sense to *rename* events. Take a look at the example below to get a better understanding.
**controller-component**

<furo-demo-snippet no-demo source style="height:200px">
<template>
   <!-- The inner part of the controller is not accessible from outside -->
    <button @-click="^^playClicked">play</button>
    <button @-click="^^pauseClicked">pause</button>
    <button @-click="^^stopClicked">stop</button> 
</template>
</furo-demo-snippet>

**my-player**

<furo-demo-snippet no-demo flow style="height:200px">
<template>
   <!-- The inner part of the controller is not accessible from outside -->
    <controller-component @-play-clicked="--playClicked" @-pause-clicked="--pauseClicked"></controller-component>
    <music-player ƒ-play="--playClicked" ƒ-pause="--pauseClicked"></music-player>
</template>
</furo-demo-snippet>

*Imagine a simple controller component with some buttons.*
*Each of them will dispatch a simple `click`.*
*Using @-click on the controller inside of the my-player can not distinguish which button was pressed.*

[learn more about events...](https://developer.mozilla.org/en-US/docs/Web/Events)

## Non bubbling events 
Non bubbling events will, as the name says, not bubble and stop at the next dom parent.

To fire a non-bubbling-event use **^event-name**.

<furo-demo-snippet flow no-demo description="*non bubbling example">
<template>
<my-button @-click="--searchClicked">Search</my-button>
<!-- when my-searcher fires the response event, the data-received event will be fired -->
<my-searcher url="https://www.googleapis.com/youtube/v3/search"
ƒ-search="--searchClicked"
@-response="^data-received">                   
</my-searcher>
</template>
</furo-demo-snippet>



## Bubbling events
To fire a bubbling-event use **^^event-name**. Bubbling is useful if you want or have to use the event in a parent component. It is a good practice to document the bubbling events from the child components. 


<furo-demo-snippet flow no-demo description="*non bubbling example">
<template>
<my-button @-click="--searchClicked">Search</my-button>
<!-- when my-searcher fires the response event, the general-error event will be fired -->
<my-searcher url="https://www.googleapis.com/youtube/v3/search"
ƒ-search="--searchClicked"
@-error="^^general-error">                   
</my-searcher>
</template>
</furo-demo-snippet>

*the* ***general-error*** *event will bubble.* 

## Non bubbling host events
With **-^** you can dispatch an event, which is available on the host only, but does not bubble. This is useful when you want 
to mimic the blur event (which does not bubble) on the outside of your component.

## Sending host data with events
Sometimes you want to send some values with your event, when the default **event.detail** is not useful. 
You can send any host property with your event by giving the property name in brackets like  ^^some-event(**propertyName**) .

**bubbling event with custom data**
```html 
   <paper-button @-click="^^some-event(_privateProperty)"> check </paper-button> 
```
*The click event sends usually a number for the amount of clicks with a certain time distance. So it will send 1 for a click, 2 for a doubleClick, 3 for a trippleClick,...*


## Sending multiple events from a single source
You can also send multiple events from a single source. 
```html 
   <paper-button @-click="^^some-event(_privateProperty),^other-event,--checkTapped"> check </paper-button> 
```
*When the button is tapped,* ***some-event*** *and* ***other-event*** *will be fired and the wire* ***--checkTapped*** *will be triggered.* 

## Stop propagation
To stop the event propagation to parent elements, add a **:STOP** to the event wires `@-error="--errorOccured, :STOP"`. 
The wires in this event-chain will be triggered. But the propagation will be stopped.

## Prevent Default
Prevent default can be achieved by using **:PREVENTDEFAULT**. 
