---
title: "Quickstart Guide"
weight: 6
---


# Quickstart Guide
The easiest way to try out furo is by using `furo init` with the furo-BEC container.
The container brings all the additional tools you need to generate the grpc stubs.


In this guide we will setup a furo spec project with the `furo init` command. 
The furo cli will then create the needed files to have a working project with a sample µType and µService definition.  

### Steps to setup a spec project

```bash
# 1 Create the project folder
mkdir sample-spec

# 2 Switch in to the project folder
cd sample-spec

# 3 Start the container 
docker run -it --rm -v `pwd`:/specs -v ~/.ssh:/root/.ssh thenorstroem/furo-bec

# 4 Run furo init
furo init

# 5 Enter your repository name
github.com/yourname/sample-specs

# 6 Install the dependencies to other spec projects
furo install

# 7 Edit your muspecs and or specs

# 8 Start the default flow
furo

# Commit your changes
```

### Detailed explanation of the steps

#### Step 1 
This folder must be under version control with git to make your specs installable by other spec projects.

#### Step 3 Start the container
If you already have an environment for proto and grpc development, you can [install](/docs/installation/) and use the furo cli
and all the other required commands locally. The [furo-BEC](/docs/tools/BEC/) comes with a set of needed or useful tools.

#### Step 4 + 5 Run furo init
This command will create a [project structure](/docs/quickstart/#project-structure-after-running-furo-init) with good default settings to begin with.
All you have to enter is the repository name for your project.

> If there is already a `.furo` file in the folder, the init command will abort. 

#### Step 6 Install the dependencies to other spec projects
A spec project can use definitions from other spec projects. The created example will use the FuroBaseSpecs which comes with a lot of
additional types to use (the google WellKnownTypes , google.protobuf.Any and many more).

#### Step 8 Start the default flow
By running `furo` the configured default flow will be executed. This command is similar to [`furo run default`](/docs/commands/furo_run/).

The example is configured to do the following flow steps:

- `deprecated` : Check for used types which are declared as deprecated.
- `muSpec2Spec` : Create / Update the specs using the muSpecs as input.
- `clean_dist` : Configured command which starts a script to delete the dist folder.
- `checkImports` : Check for types which are not imported, add the imports when furo can resolve it.
- `genMessageProtos` : Create the protos for the messages.
- `genServiceProtos` : Create the protos for the services.
- `buf_generate` : Configured command which starts the buf generator.
- `gen_transcoder` : Configured command which produces a go package with a [grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway/) server mux. 
- `genEsModule` : Generates a es module with the specs which is used by [eclipsefuro-web](/docs/web-components/).

### Project structure after running furo init
The initialized project will have a
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