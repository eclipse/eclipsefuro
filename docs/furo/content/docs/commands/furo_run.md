---
weight: 6
title: "run"
---

## furo run

Runs a configured flow.

### Synopsis

Runs a configured flow of furo commands.

To configure a flow of commands just add a "flow" in the flows section of your .furo config.
A flow is just a list of commands which gets executed in order

Example Config:

	[.furo]
	commands:
	  publish_npm: "./scripts/test.sh"
	flows:
	  type:
		- cleanTypeProtoDir
		- muSpec2Spec
		- TypeSpec2Proto
		- publish_npm

Command:

This config will run "cleanTypeProtoDir",  muSpec2Spec"" and "TypeSpec2Proto" in sequence and calling the command publish_npm

Tipp: If you need the types and services in your command, just call furo again. 

Like:
    #!/bin/bash

    # generate the type documentation...
    furo exportAsYaml | simple-generator -t scripts/typedoc.tpl > dist/typedoc.md

 

```
furo run [flags]
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



###### Auto generated by spf13/cobra on 23-Oct-2020