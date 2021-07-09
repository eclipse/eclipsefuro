---
title: Lifecycle
weight: 80
---
# FBP Lifecycle

## _FBPReady()
_FBPReady() is called, when the wires are registered and the components are able to receive wire data.
This is also the earliest point to enable the tracing.

```javascript
class MyComponent extends FBP(LitElement) {  
  // trace all wires
  _FBPReady(){
    super._FBPReady();
    this._FBPTraceWires();
  }
}
window.customElements.define('my-component', MyComponent);

```

## The **--FBPready** *magic* wire
The wire --FBPready is also triggered when your component is ready.
