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
Bash mode

    docker run -it --rm -v `pwd`:/specs thenorstroem/furo-bec
    # do your stuff
    # type exit to quit
    exit

Command mode

    docker run -it --rm -v `pwd`:/specs thenorstroem/furo-bec build


> TIPP: If your custom commands have to access different directories, do not forget to mount them.


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

> No furoc generators are installed. Add the needed furoc-gen-XXX to the `.furobecrc` file.

