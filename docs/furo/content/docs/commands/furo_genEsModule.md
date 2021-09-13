---
weight: 60
title: "genEsModule"
---

## furo genEsModule

This command will generate es6 spec module to use in your client projects.

The `@furo/data` components needs them to validate user input against the defined constraints, calculates the deltas to send to the server in a PATCH request,
forward them to the input components like the ones from `@furo/ui5` with data binding, apply error messages from the server to the corresponding field, resolve the correct component to display or edit the data fields, ...


{{< hint info >}}
**Hint:** Publish the env file with npm, to make them installable with proper versioning for the web projects.

{{< /hint >}}


```
furo genEsModule [flags]
```

### Options

```
  -h, --help   help for genEsModule
```

### Options inherited from parent commands

```
      --config string   config file (default is CWD/.furo.yaml)
```

