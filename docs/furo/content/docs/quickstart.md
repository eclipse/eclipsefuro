---
title: "Quickstart Guide"
weight: 6
---


# Quickstart Guide
The easiest way to try out furo is by using `furo init` with the furo-BEC container.
The container brings all the additional tools you need to generate the grpc stubs.

> If you already have an environment for proto and grpc development, you can [install](/docs/installation/) the furo cli 
> localy and use it directly.


In this guide we will setup a furo spec project with the `furo init` command. 
The furo cli will then create the needed files to have a working project with a sample µType and µService definition.  

### Steps

```bash
# 1 Create the project folder
mkdir sample-spec

# 2 Switch in to the project folder
cd sample-spec

# 3 Start the container 
docker run -it --rm -v `pwd`:/specs thenorstroem/furo-bec

# 4 Run furo init
furo init

# 5 Enter your repository name
github.com/yourname/sample-specs

# 6 Install the dependencies
furo install

# Edit your muspecs and or specs

# 7 Start the default flow
furo

# Commit your changes
```


### Project structure after running furo init
The project have a
- [.furoc](/docs/configuration/) file which contains the project configuration
- muspec folder with some examples
- empty specs folder
- some example scripts to use in your flows

*initial project*
```
.
├── buf.gen.yaml
├── buf.yaml
├── go.mod
├── muspecs
│   └── sample
│       ├── Sample.services.yaml
│       └── sample.types.yaml
├── scripts
│   ├── buf_breaking.sh
│   ├── buf_generate.sh
│   ├── cleanUpDist.sh
│   └── gprcgateway
│       ├── autoregister.go.tpl
│       └── generate.sh
└── specs

```