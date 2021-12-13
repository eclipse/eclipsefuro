---
title: Script interaction
weight: 40
---

# Interaction with javascript
The wires can be hooked and triggered from the script part of your component. 
This also comes very handy when you are writing tests. 

## Trigger a wire imperatively

To trigger a wire from the javascript part of your component, call the **_FBPTriggerWire** method.

You can also trigger the wire in the constructor, in this case the wire will be queued until the flow is parsed and ready.
 
```js  {linenos=table,hl_lines=[6,14,22],linenostart=1}
class TriggerSample extends FBP(LitElement) {
    constructor(){
      super();
      this.data = "Test";
      // this wire will be queued
      this._FBPTriggerWire("|--wireName", this.data);
    }
    
    /**
     * _FBPReady triggers when the flow is ready
     */
    _FBPReady(){
      super._FBPReady();
      this._FBPTriggerWire("|--wireName", this.data);
    }
    
    /**
     * To pass data from outside to a wire, use this._FBPTriggerWire()
     * 
     */
    fetchRecord(src){
      this._FBPTriggerWire("|--fetchRequested", src);
    }
}
``` 


## Add a wire hook
To hook on a wire use `this._FBPAddWireHook("--wirename")`. This comes very handy at testing, or if you have to manipulate some 
data, because the component doesnt send it like an other component needs it.

```javascript {linenos=table,hl_lines=[6,7,8],linenostart=1}
class HookSample extends FBP(LitElement) {

    constructor() {
        super();
        
        this._FBPAddWireHook("--pathChanged",(e)=>{
          // On hooks `e` contains the complete event, not only the `e.detail`        
      })
    }
}
```

> The most @-events of the furo base components will fit the ƒ-methods of the corresponding components.
It is like playing domino.
