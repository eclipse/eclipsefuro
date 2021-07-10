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
You can use furo FBP inside your HTML by using `furo-fbp`. Make sure that you have 
imported the components that you want to use. 

You can import furo-fbp from the CDN or npm. Note the **type="module"**

**CDN**

`<script type="module" src="https://cdn.jsdelivr.net/npm/@furo/fbp@5.8.1/assets/furo-fbp.js"></script>`

**npm**

`<script type="module" src="/node_modules/@furo/fbp/src/furo-fbp.js"></script>`


```html
<head>
    <script type="module" src="https://cdn.jsdelivr.net/npm/@furo/fbp@5.8.1/assets/furo-fbp.js"></script>
</head>
<body>
  <!-- use the component as many times you want, even before the definition-->
  <language-sample></language-sample>
  <hey-component></hey-component>
  
  
  <!-- define the component -->
  <furo-fbp name="language-sample">
    <template>
      <button @-click="--playClicked">play</button>
      <button @-click="--pauseClicked">pause</button>
      <audio ƒ-play="--playClicked" 
             ƒ-pause="--pauseClicked" 
             src="https://upload.wikimedia.org/wikipedia/commons/9/92/German_alphabet-2.ogg"></audio>
      
      <!-- use other components that you have defined -->
      <hey-component></hey-component>
    </template>
  </furo-fbp>
  
  
  <!-- define the component -->
  <furo-fbp name="hey-component">
    <template>
      <div>Hej</div>
    </template>
  </furo-fbp>
</body>

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
