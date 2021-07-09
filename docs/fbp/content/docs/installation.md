---
weight: 2
title: "Installlation"
---

# Installation

First of all you need the npm module **@furo/fbp**.

```bash
npm i -S @furo/fbp
```


## FBP with lit-element
To use FBP with lit, just extend your class.

```javascript
class MyComponent extends FBP(LitElement) {
  
}
window.customElements.define('my-component', MyComponent);
```



## FBP with native web-components
To use furo-fbp with native components, call `this._appendFBP(this.shadowRoot);` to enable fbp.

```javascript
class MyComponent extends FBP(HTMLElement) {

  constructor() {
    super();
    // Create a shadow root to the element.
    this.attachShadow({mode: 'open'});
    this.shadowRoot.appendChild(template.content.cloneNode(true));
    // Append FBP to my-component
    this._appendFBP(this.shadowRoot);
  }
 
}

window.customElements.define('my-component', MyComponent);

```

## FBP within HTML
You can use furo FBP inside your HTML by using `flow-bind`. Make sure that you have 
imported the component that you want to use.

```html

<flow-bind>
  <template>
    <button @-click="--btnClicked">sender</button>
    <div ƒ-remove="--btnClicked">receiver</div>
  </template>
</flow-bind>
```


## FBP with polymer
To use FBP with polymer, just extend your class.
```javascript
class MyComponent extends FBP(PolymerElement) {
  
}
window.customElements.define('my-component', MyComponent);
```

The appender is automatically triggered from FBP via the _attachDom method.
```javascript
// you dont have to write this, its already done in furo-fbp
_attachDom(dom) {
    this._appendFBP(dom);
    super._attachDom(dom);
}
```
