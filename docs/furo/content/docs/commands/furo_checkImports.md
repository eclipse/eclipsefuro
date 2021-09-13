---
weight: 20
title: "checkImports"
---

## furo checkImports

This command will check and correct all the needed imports in the specs. These imports are needed for creating proper proto files.

If a import can not be resolved, you will get a error message which tells you the file and the field which is causing the problem.

### How does the import resolving work?
Furo will look up for type definitions as close as possible to your type (most local to least local).
So types which are coming from some installed dependencies and are also defined in your project, will prefer the 
import from your project.

```bash
specs/sample/Sample.type.spec:1:1:Import unknown.Type not found in type sample.Sample on field id
```


{{< hint danger >}}
**Attention:** If you have a typo on a scalar type like `sting` instead of `string`, the current version of furo will not detect this as error.
{{< /hint >}}

```
furo checkImports [flags]
```

### Options

```
  -h, --help   help for checkImports
```

### Options inherited from parent commands

```
      --config string   config file (default is CWD/.furo.yaml)
```

