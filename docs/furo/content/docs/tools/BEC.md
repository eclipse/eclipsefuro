---
title: "B E C"
weight: 1
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: false
# bookComments: true
---


# フロー furoBEC

#### The furo build essentials container.

This container contains all tools you need to work with a furo spec project.

## Usage
### Bash mode
```bash
docker run -it --rm -v `pwd`:/specs thenorstroem/furo-bec
# do your stuff
# type exit to quit
exit
```

### Command mode
Starting your container with the arguments **build** or **publish** will execute the corresponding flow.

```bash
docker run -it --rm -v `pwd`:/specs thenorstroem/furo-bec build
```

{{< hint Info >}}
**Note:** Only the arguments **build** and **publish** are supported. 
{{< /hint >}}

{{< hint Info >}}
**Tipp:** If your custom commands have to access different directories, do not forget to mount them.
{{< /hint >}}

## Installed Tools
To see the installed versions of the tools, please refer to the [Dockerfile](https://github.com/eclipse/eclipsefuro/blob/main/BEC/Dockerfile) of the version you use.

- golang
- git
- protoc
- protoc-gen-grpc-gateway (v2)
- protoc-gen-openapiv2
- protoc-gen-go
- protoc-gen-go-grpc
- simple-generator
- furo
- furoc
- [buf](https://docs.buf.build/introduction)
- protoc-gen-buf-breaking
- protoc-gen-buf-lint
- [jq](https://stedolan.github.io/jq/)
- [yq](https://mikefarah.gitbook.io/yq/commands/read)

> No furoc generators are installed. Add the needed furoc-gen-XXX to the `.furobecrc` file.


## Customizing the container with .furobecrc
Make settings for your environment or project in this file. 

Maybe you need a `$GOPRIVATE` or other Env vars.

The `.furobecrc` is executed whenever you start the container.

```bash

# change the bash prompt
PS1="フロー myProject#"
GOPRIVATE="$GOPRIVATE,github.com/myprivaterepo"

echo '[url "ssh://git@github.com/"]
        insteadOf = https://github.com/' > ~/.gitconfig
        
```
   
