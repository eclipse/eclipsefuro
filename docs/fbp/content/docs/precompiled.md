---
weight: 2
title: "Precompiled"
bookHidden: true
---

# @furo/precompiled

The [@furo/precompiled](https://github.com/theNorstroem/precompiled-furo-web-components) package is for those who want to use the furo componetnt directly in HTML.
There is no build step needed. It is a convenient way to use the components to prototype some ideas or just play around
with FBP without a complex installation procedure. When you know that your ideas work, transfer them 1:1 to a web component,
so others can install, use and extend them.

Look at the [sample](https://github.com/theNorstroem/precompiled-furo-web-components) folder in the repo to see an example of what is possible.

> **Note** This package was created to teach furo FBP and for creating our demos,
> so we can use and show our components in a HUGO generated page.
Some of the files are very big at the moment, because they are not optimized for file size yet.
 
This documentation uses @furo/precompiled at any place where you can see a `[demo/source/flow]` panel.
All that was needed was to load the precompiled scripts via CDN.

- [open demo](https://kgjfw.csb.app/)
- [open sandbox](https://codesandbox.io/s/hardcore-maxwell-kgjfw?from-embed)
 
```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <meta http-equiv="X-UA-Compatible" content="ie=edge"/>
  <!-- This module contains the light bulb -->
  <script type="module" src="https://cdn.jsdelivr.net/npm/@furo/precompiled@2.0.0-rc.16/dist/doc-helper.js"></script>
  <script>
    import(
      "https://cdn.jsdelivr.net/npm/@furo/precompiled@2.0.0-rc.16/dist/DOMFBP.js"
      ).then(() => {
      // activate FBP on body
      const fbphandle = new DOMFBP(document.body);
      // enable tracing
      fbphandle._FBPTraceWires();
    });
  </script>

  <title>Static Template</title>
</head>
<body>
  <h1>This is a static template, there is no bundler or bundling involved!</h1>
  
  <light-bulb ƒ-toggle="--lightSwitchClicked"></light-bulb>
  <button @-click="--lightSwitchClicked">i am a lightswitch</button>
</body>
</html>
``` 

## Compatibility Table
This compatibility list is only meant for the precompiled components.

![compat](/compat.png)

