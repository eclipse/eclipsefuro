---
weight: 30
title: "exportAsYaml"
---

## furo exportAsYaml

Exports all specs for types and services and also the current config in one big yaml file to stdout.

This comes very handy in custom generators which are ok with the yaml file structure.

Feel free to add custom sections in the config to use them in your own scripts.

```yaml

installedServices: [...] # The list of installed services
installedTypes: [...] # The list of installed types
services: [...] # The list of services from your project
types: [...] # The list of types  from your project
config:
   module: mod # custom config section
   myopts: # custom config section
      remoteDir: "path/to/somewhere"
      otherCustomSetting: true

```

### running with the `-f` option
You will notice that the output structure of running `furo exportAsYaml` does have a different structure. Your definitions will be included in the following strurcture:

```yaml
        specdir: dependencies/github.com/theNorstroem/furoBaseSpecs/specs
        path: furo
        filename: big_decimal.type.spec
        typespec: ... # the spec is available under this property
```

```
furo exportAsYaml [flags]
```

### Options

```
  -f, --full   Include the ast info like filenames and more.
  -h, --help   help for exportAsYaml
```

### Options inherited from parent commands

```
      --config string   config file (default is CWD/.furo.yaml)
```

