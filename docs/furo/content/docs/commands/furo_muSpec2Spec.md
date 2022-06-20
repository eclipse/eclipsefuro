---
weight: 10
title: "muSpec2Spec"
---

## furo muSpec2Spec

Updates or creates the specs with the definitions from the µSpecs.

Renaming and restructuring your spec folder is ok, furo will always update the correct file. Also a reordering of the fields in your spec,
to get nicer protos of for other reasons, is also ok. Furo works with a internal AST and is not interested in order or file names.


{{< hint warning >}}
**Attention:** 
if you have set the config option `muSpec.forceSync: true` in your furo config,
the command will always enable the `-d` option.
{{< /hint >}}


Because the specs have a higher information density then the µSpecs, furo will fill in calculated default 
values on the first time it creates a spec file. Later calls of this command will not touch the prefilled fields of the spec.

### Config

```yaml
muSpec:
  types: # define a set of globs which matches your type definitions
    - "./muspecs/**/*types.yaml"
    - "./muspecs/*types.yaml"
  services: # define a set of globs which matches your service definitions
    - "./muspecs/**/*services.yaml"
    - "./muspecs/*services.yaml"
  goPackageBase: "github.com/yourname/sample-specs/dist/pb/" # this is used to prefix the go package option
  javaPackagePrefix: "com.example.tutorial."
  dir: "muspecs" # the folder where you save the µSpecs
  forceSync: true # This will delete specs which are deleted in muSpec, this is very useful during prototyping  
  forceLabels: true # This will overwrite the label and placeholder texts during muSpec2Spec command
  requestTypeSuffix: "Request" # Suffix for the generated request type specs
labelPrefix: "" # prefix for the enum , label, placeholder text keys
```



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



