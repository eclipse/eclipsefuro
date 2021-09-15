---
title: "Working with Dependencies"
weight: 19
---

# Working with Dependencies
You can install another spec in to your project, to use their types.

{{< hint warning >}}
**Note:** A installable dependency must be reachable by your `git` command.
{{< /hint >}}


## Install a spec project
To install a dependency add them to the `dependencies` section of the config.
It is a list with git repositories followed by a version tag or branch name.

By running `furo install` the dependency will be added to the configured directory in your project. It is up to you, if you want to add this folder to your VCS.


```yaml
dependencies:
  - "git@github.com:theNorstroem/furoBaseSpecs.git v1.27.1" # We use this to have the well known types available.
dependenciesDir: dependencies # Your installed dependencies from other spec project are saved in this directory
```

Installed dependencies are saved with their full path in to your `dependencyDir`. 

Like `dependencies/github.com/theNorstroem/furoBaseSpecs/...`


{{< hint info>}}
**Tipp:** If you work with `buf`, do not forget to add the folder to the roots of your buf config.

```yaml
build:
  roots:
    - dist/protos
    - dependencies/github.com/theNorstroem/furoBaseSpecs/dist/proto
```
{{< /hint >}}

{{< hint danger >}}
**Attention:** Dependencies of dependencies are not resolved. You will have to add them to your project too. 

To identify them is not so difficult. `furo checkImports` can give you a hint for the types that you use directly and 
protoc or buf will also tell you if you have forgotten something.
{{< /hint >}}

## Make your project installable
To make your project installable by other projects, add the `dist` section in to your project. 
You have to define, which files should be available in the projects that installs your project.

It depends on your use case, which files you will provide. At least you have to provide the **.furo** file and the **specs** directory.
Providing the generated **protos** is also strongly recommended, so the consumer of your spec can add them easily in their buf config (if grpc is a target for you).

```yaml
dist:
  files: # enter a list of files and directories which should be installed by other projects when they install your spec projec
    - dist/protos
    - specs
    - README.md
    - .furo
```