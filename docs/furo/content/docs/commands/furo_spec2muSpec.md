---
weight: 90
title: "spec2muSpec"
---

## furo spec2muSpec


Updates or creates the µSpecs with the definitions from the specs.


> if you have set the config option `muSpec.forceSync: true` in your furo config,
> the command will always enable the `-d` option.

Because the specs have a higher information density then the µSpecs this step will *loose* information.

This command is useful if you want to switch your source of truth from [spec](/docs/sourceoftruth/#spec-as-source) to [µSpec](/docs/sourceoftruth/#µspec-as-source).

This command is useful if you want to discuss your services and types without having unneeded details.


```
furo spec2muSpec [flags]
```

### Options

```
  -d, --delete   Delete muSpecs which not exist in Spec
  -h, --help     help for spec2muSpec
```

### Options inherited from parent commands

```
      --config string   config file (default is CWD/.furo.yaml)
```



###### Auto generated by spf13/cobra on 23-Oct-2020
