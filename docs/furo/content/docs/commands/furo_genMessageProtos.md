---
weight: 40
title: "genMessageProtos"
---

## furo genMessageProtos

Generate the message protos from the type specs.

The generated proto messages will be generated to the configured path. 

> As you may have noticed the specs have a attribute `targetfile`. Type specs with the same target file value will be generated in to the same file.


{{< hint danger >}}
**Attention:** You can not target Services and Messages to the same file, because for furo, the generating of types/messages and the generation of services are two different tasks.
{{< /hint >}}


*.furo config example*
````yaml
build:
  proto:
    targetDir: "./dist/protos" #Hint: add this to your proto include path
````
		 
		



```
furo genMessageProtos [flags]
```

### Options

```
  -h, --help   help for genMessageProtos
```

### Options inherited from parent commands

```
      --config string   config file (default is CWD/.furo.yaml)
```


