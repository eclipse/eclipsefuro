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

Each spec project must have a configuration file. The furo cli tool will look up for *.furo* in the current working directory or can be called with the `--config` option.

You can generate a config file with a inital project structure by calling `furo init`. 

As you can see, the config file is a regular yaml file and is easy to edit in any text editor or IDE.

*Example of a configuration file*
```yaml
furo: "1.28.5" # the minimal version of furo
module: "github.com/yourname/sample-specs"
version: "v0.0.1" # This is the version for your spec project
specDir: "./specs" # The spec files are generated to this folder
specFormat: "yaml" # set to yaml or json
dependencies:
  - "git@github.com:theNorstroem/furoBaseSpecs.git v1.27.1" # The importer looks for all **/*.type.spec files recursive The importer looks for all **/*.service.spec files recursive
dependenciesDir: dependencies # Your installed dependencies from other spec project are saved in this directory
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
labelPrefix: "" # prefixe for the enum, label, placeholder text keys.
commands: #camelCase is not allowed, command scripts can only be executed from a flow
  gen_transcoder: "./scripts/gprcgateway/generate.sh" # shell script to generate a half grpc gateway
  buf_generate: "./scripts/buf_generate.sh"
  buf_braking: "./scripts/buf_breaking.sh"
  clean_dist : "./scripts/cleanUpDist.sh" # Deletes the dist folder
flows:
  default: #we choose µSpec as source https://fidl.furo.pro/docs/sourceoftruth/#%C2%B5spec-as-source
    - deprecated
    - muSpec2Spec
    - clean_dist
    - checkImports
    - genMessageProtos
    - genServiceProtos
    - buf_generate
    - gen_transcoder
    - genEsModule
build:
  proto:
    targetDir: "dist/protos/" # Hint: add this to your proto include path
  esModule:
    targetFile: "dist/env.js" # The environment file to register in eclipsefuro-web
dist:
  files: # enter a list of files and directories which should be installed by other projects when they install your spec projec
    - dist/protos
    - specs
    - README.md
    - .furo
extensions: # Add configurations for your private builders and generators here.
  gen_transcoder:
    additional_imports: # if you use any any type, you need to import them for the gateway
      - github.com/yourname/sample-specs/dist/pb/sample
```