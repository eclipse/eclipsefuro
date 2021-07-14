---
weight: 2
title: "Installlation"
---

# Installation

You can use furo FBP directly in your HTML documents or within your own web-components.


## FBP and eclipsefuro-web in HTML without a *installation*
You can use furo-fbp and the furo-web-components without a direct installation by using the
precompiled variant of the furo web components.

This is the simplest way to get up and running. 

[Read more about that here](precompiled.md).

## Install the npm module

To work with lit or native web components you need the npm module **@furo/fbp**.

```bash
npm i -S @furo/fbp
```

## FBP with lit-element
To use FBP with lit, just extend your class.

```javascript {linenos=table,hl_lines=[1],linenostart=1}
class MyComponent extends FBP(LitElement) {
  
}
window.customElements.define('my-component', MyComponent);
```



## FBP with native web-components
To use furo-fbp with native components, call `this._appendFBP(this.shadowRoot);` to enable fbp.

```javascript {linenos=table,hl_lines=[8,9],linenostart=1}
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


## FBP with polymer
To use FBP with polymer, just extend your class.
```javascript {linenos=table,hl_lines=[1 ],linenostart=1}
class MyComponent extends FBP(PolymerElement) {
  
}
window.customElements.define('my-component', MyComponent);
```
