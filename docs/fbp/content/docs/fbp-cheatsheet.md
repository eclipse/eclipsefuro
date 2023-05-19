---
title: Cheat Sheet
weight: 7
---

# Cheat Sheet
Furo fbp only knows a few of keywords to control the application flow.

- [**at-**](/docs/fbp-cheatsheet/#adding-event-listeners)
- [**fn-**](/docs/fbp-cheatsheet/#execute-exposed-methods)
- [**set-**](/docs/fbp-cheatsheet/#set-member-variables-of-a-component)
- [**at-fnret-**](/docs/fbp-cheatsheet/#receiving-the-return-values-from-fn--calls)
- [**^**](/docs/fbp-cheatsheet/#non-bubbling-events)
- [**^^**](/docs/fbp-cheatsheet/#bubbling-events)
- [**-^**](/docs/fbp-cheatsheet/#non-bubbling-host-events)
- [**(( ))**](/docs/fbp-cheatsheet/#setting-member-variables-of-the-host-component)


## Adding event listeners
The **at-** keyword let you bind a wire to an event.
This can be any browser event or custom event.

`at-click="--lightSwitchClicked"` Assign a wire which should be triggered when the event happens.
By default the `event.detail` property is put as data on the wire.

To put the event root on the wire, use `at-click="--lightSwitchClicked(*)"`.

## Execute exposed methods

The **fn-** keyword let you trigger methods with the data which was put on the wire as argument.

If a method needs multiple arguments you have to put a array on the wire data, otherwise only the first argument is filled.

<furo-demo-snippet flow>
<template>
  <light-bulb fn-toggle="--lightSwitchClicked" fn-set-color="--newColor"></light-bulb>
  <!-- This button acts as a light switch -->
  <button at-click="--lightSwitchClicked">i am a lightswitch</button>
  <input type="color"  at-input="--newColor(*.target.value)">
</template>
</furo-demo-snippet>

> The classic notation for **at-** is **@-**. The classic notation for **fn-** is **ƒ-**.


## Set member variables of a component
The **set-** keyword let you set member variables (not the attributes) of element.

`<span set-inner-text="--labelReceived"> initial label </span>` Will set the attribute `innerText` on the span,
as soon the wire `--labelReceived` is triggered.

> The classic notation for **set-** is **ƒ-.**.

## Setting member variables of the host component
The **(( ))** keyword let you store data from events on member variables of the host component.

`<my-component at-event-name="((color))">` This will store the data of the event `event-name` on the
host member variable `color`.

## Emitting events
### Non bubbling events
Non bubbling events will, as the name says, not bubble and stop at the next dom parent.
To fire a non-bubbling-event use **^event-name**.

### Bubbling events
To fire a bubbling-event use **^^event-name**.
Bubbling is useful if you want or have to use the event in a parent component.


### Non bubbling host events
With **-^** you can dispatch an event, which is available on the host only, but does not bubble.


<furo-demo-snippet flow>
<template>
  <light-bulb on="" fn-toggle="--lightSwitchClicked" fn-set-color="--colorSetClicked"></light-bulb>
  <button at-click="--lightSwitchClicked">i am a lightswitch</button>
  <!-- we put the host member variable color on the wire, --colorSetClicked(color) -->
  <button at-click="--colorSetClicked(color)">setcolor</button>
  <!-- We catch a deep property of the input event and emit this value as new event color-changed.
    We store the value on the host variable color.-->
  <input type="color" at-input="^color-changed(*.target.value)" at-color-changed="((color))">
</template>
</furo-demo-snippet>




## Receiving the return values from fn- calls
The **at-fnret-** keyword let you wire the response of a method call.
The data which is put on the wire is the response of the method.

## Stop propagation
The keyword **:STOP**  is used to stop the event propagation to parent elements.

## Prevent default
The keyword **:PREVENTDEFAULT** will set the prevent default on the event.
