# furoc

The `furoc` "compiler" is a generator tool similar to `protoc`.

### When furoc and when protoc?
- The furo specs can [translate to *.proto](https://github.com/eclipse/eclipsefuro/furo/blob/master/doc/furo_genMessageProtos.md) and [vice versa](https://github.com/theNorstroem/protoc-gen-furo-specs). When you already have protoc plugins, use them.
- The furo specs have a much higher information density then the proto specs have. For generating ui components with furoc-gen-u33e, proto is not enough.
- Furoc does not stop on an incomplete import chain. If you need this, use protoc before furoc.
- Using furoc and protoc in combination gives you a lot of advantages, [furoc plugins are easier to write,...](#writing-your-own-plugins)

## Usage
You can configure the arguments in your .furo file or give the arguments in the cli.

#### running with command arguments:
As soon you pass a Sxxx or Fxxx argument, only these will be built. To generate everything, do not pass a Sxxx or Txxx argument. 

This command will generate **only** service:`reference-search`, type:`form`  and service:`collection-dropdown` components

```shell script
furoc 
-I./pathTto/spec/project
--plugin=furoc-gen-u33e
--u33e_out=
Sreference-search,\
Tform,\
Scollection-dropdown,\
:outputDirectoryForGenU33e
```

#### running with config:
Same rules from "run as command" are applied. 

Furoc will look for a `.furoc` file in the current directory.

```shell script 
furoc
```

**Example .furoc file:**

```yaml
furoc:
  Input:
      - ./
  Commands:
    - OutputDir: dist/u33e
      Plugin: furoc-gen-u33e
      Args:
          - Sreference-search
          - Tform
          - Scollection-dropdown
```




## Command arguments

### `-I./path`
Defines the spec projects to include. At the moment only one spec project include is supported.

### `--plugin=./path/to/bin/furoc-gen-pluginname`
Defines the binary to use. If you do not use this option, the binary in `$PATH` will be used. 

### `--pluginname_out=arg1,arg2,argN:ouputdir`
Defines which plugins furoc should use.

- `pluginname_out` will translate to furo-gen-pluginname
- `arg1,arg2,argN` will be passed as arguments for the plugin. Read on the documentation of the plugin which arguments are valid.
- `:outputdir` the generated files from the plugin will written to this directory

## Writing your own plugins
It is not so difficult to write a plugin for furoc. Look at the sample `furoc-gen-sample` or look at other plugins.

### Interface
Furoc will pass a yaml with the current config of the spec project, the types, the services, the installed types and the installed services.
As an example input for your plugin, look at sample/fullyaml.yaml.

The list will go to `stdin` of your plugin.

Furoc exects a list of files as response. The response goes to `stdout` of your plugin.

You can use the package **reqres** which handles all the stdin stdout encoding decoding stuff for you.  

### Furoc plugins in other languages
At the moment, only serialized go structs are accepted as response. A variant which accepts proto (with the same interface protoc excepts the responses) will come; if requested. 