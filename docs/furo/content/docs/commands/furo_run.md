---
weight: 6
title: "run"
---

## furo run

Runs a configured flow.

### Synopsis

Runs a configured flow of furo commands.

A flow is a list of *custom commands* or *furo run commands* which are executed.

To configure a custom command, add them to the commands section of the .furo config file.

Example Config:
```yaml

commands:
  publish_npm: "./scripts/test.sh"
flows:
  default:
    - muSpec2Spec
    - publish_npm
  publish:    
    - publish_npm

```


Tipp: If you need the types and services in your command, just call furo again. 

```bash
#!/bin/bash
# generate the type documentation...
furo exportAsYaml | simple-generator -t scripts/typedoc.tpl > dist/typedoc.md
```
 

```
furo run [name of the flow]
```

### Options

```
  -f, --flow string   A configured flow from the .furo config (default "default")
  -h, --help          help for run
```

### Options inherited from parent commands

```
      --config string   config file (default is CWD/.furo.yaml)
```

