---
title: "Enums"
weight: 50
---

# Anatomy of a enum µSpec
The µType specs are regular yaml files.

The "enum-object" contains 4 fields. The fields `enum`, `values` are mandatory, the field `target` and `alias` are optional.

You can have as many enum definitions per file as you want. It makes sense that you put enums in a file, that belongs togehter.

It is not possible to mix enums and types in the same µSpec file at the moment.



*File: muspec/sample.enums.yaml*
```yaml
- enum: 'helloworld.Corpus #Description for a enum sample with aliases'
  values:
    UNKNOWN: 0
    STARTED: 1
    RUNNING: 1
    COMPLETE: 2
  target: enums.proto
  alias: true # this is needed when you want to allow aliases
```

As you can see, the Corpus enum's first constant maps to zero: 
every enum definition must contain a constant that maps to zero as its first element. 
This is because for the compatibility with protobuf.

## The enum definition line

```yaml
- enum: 'sample.Corpus #Description for a enum sample with aliases'
  !___!  !____||_____! !__________________________________________!
    |       |     |          |
    |       |   "type" name  |
    |    package             |
    |                    description (recomended) begins with a #
    |                  
    | 
    | 
field name to declare a enum


```

It is a good practice to give a good description. 

## Values

In the "values-object", you define the values of the enum.


```yaml
values:
  UNKNOWN: 0
  STARTED: 1
  !_____! !_! 
    |      |   
    |      |   
 constant  |    
           |     
         Value

```

## Alias
Alias is a boolean field to define if you want to allow aliases. 

[Read more about aliases here.](https://developers.google.com/protocol-buffers/docs/proto3#enum)

## Target
In the field target, you can define the target proto file for the enum definition.
