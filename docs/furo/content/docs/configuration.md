---
title: "Project Configuration"
weight: 10
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: false
# bookComments: true
---


# Configuration

Each spec project must have a *.furo* configuration file.

A fictive example of a furo config for a spec project.
```yaml
furo: "1.20.0" #Minimal furo version
specDir: "./specs"
specFormat: "yaml" #set to yaml or json
dependencies: #do not write every type/message again and again. Install them.
  - "git@github.com:theNorstroem/furoBaseSpecs.git v1.11.8" # The importer looks for all **/*.type.spec files recursive The importer looks for all **/*.service.spec files recursive
dependenciesDir: dependencies #directory where the dependencies get installed to.
muSpec: 
  types:
    - "./muspecs/**/*types.yaml"
    - "./muspecs/*types.yaml" #Use this if you do not put your specs in folders
  services:
    - "./muspecs/**/*services.yaml"
    - "./muspecs/*services.yaml" #Use this if you do not put your specs in folders
  goPackageBase: "github.com/yourname/appname-specs/dist/pb/" #this is used to prefix the go package option
  javaPackagePrefix: "com.yourname.something"
  dir: "muspecs"
commands: #camelCase is not allowed, command scripts can only be executed from a flow
  gen_grpc_gateway: "./scripts/grpcgateway/gateway.sh"
  exec_protoc : "./scripts/protoc_command.sh"
flows:
  default: #we choose µSpec as source https://fidl.furo.pro/docs/sourceoftruth/#%C2%B5spec-as-source
    - muSpec2Spec #Updates the specs from the µSpecs
    - genMessageProtos #Generates the protos from the type specs
    - genServiceProtos #Generates the protos from the services specs
    - exec_protoc #Custom script
    - gen_grpc_gateway #Custom script
    - genEsModule #Generates specs to use in clients
build: #build config, define the targets here
  proto:
    targetDir: "dist/protos/" #Hint: add this to your proto include path
  esModule:
    targetFile: "dist/env.js" #env module for the furo client libs 
  bundledservice:
    targetFile: "dist/allservices/all-services.proto"
    package: "allservices"
    options:
      go_package: "github.com/yourname/appname-specs/dist/pb/allservices;allservicespb"
      java_multiple_files: true
      java_outer_classname: "AllServices"
      java_package: "com.yourname.something.allservices"
dist: # this is for furo install called on other projects, 
  files: # enter a list of files and directories which should be installed by other projects
    - dist/protos
    - specs
    - README.md
    - .furo
```