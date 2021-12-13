---
title: Lifecycle
weight: 80
---
# FBP Lifecycle

## _FBPReady()
_FBPReady() is called, when the wires are registered and the components are able to receive wire data.


```javascript
class MyComponent extends FBP(LitElement) {
  /**
   * flow is ready lifecycle method
   */
  _FBPReady() {
    super._FBPReady()
    this._FBPTraceWires()
  }
}
window.customElements.define('my-component', MyComponent);

```

## The **|--FBPready** *magic* wire
The wire `|--FBPready` is triggered when your component is ready.

This can be used as a start trigger for your component. 
