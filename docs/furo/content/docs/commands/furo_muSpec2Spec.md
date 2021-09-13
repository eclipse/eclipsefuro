---
weight: 10
title: "muSpec2Spec"
---

## furo muSpec2Spec

Updates or creates the specs with the definitions from the µSpecs.

Renaming and restructuring your spec folder is ok, furo will always update the correct file.

> if you have set the config option `muSpec.forceSync: true` in your furo config,
> the command will always enable the `-d` option.

Because the specs have a higher information density then the µSpecs, furo will fill in calculated default 
values on the first time it creates a spec file. Later calls of this command will not touch the prefilled fields of the spec.

### Options

```
  -d, --delete   Delete specs which are not defined in muspecs
  -h, --help     help for muSpec2Spec
      --overwrite-spec-options   Overwrite the proto options section in the spec files
```

### Options inherited from parent commands

```
      --config string   config file (default is CWD/.furo)
```



