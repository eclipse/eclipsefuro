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

    docker run -it --rm -v `pwd`:/specs thenorstroem/furo-bec
    # do your stuff
    # type exit to quit
    exit

### Command mode
This will run furo with the configured flow *build*.

    docker run -it --rm -v `pwd`:/specs thenorstroem/furo-bec build


> TIPP: If your furo commands have to access different directories, do not forget to mount them.


## Installed Tools

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
- buf
- protoc-gen-buf-breaking
- protoc-gen-buf-lint

- [jq](https://stedolan.github.io/jq/)
- [yq](https://mikefarah.gitbook.io/yq/commands/read)

> No furoc generators are installed. Add the needed furoc-gen-XXX to the `.furobecrc` file.

