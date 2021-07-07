---
title: "Types"
weight: 10
---

# Anatomy of a type spec
A type spec consists of 2 main sections (fields and __proto) and some properties on the root.

## name
This property is deprecated. Fill in the same as you fill in on the property `type` for compatibility with some old generators.

## type
The type of the field without the package name. Write it in CamelCase with a capital letter at the beginning.

## __proto
The proto section defines some properties to generate the proto files.

```yaml
__proto:
    package: auth
    targetfile: auth.proto
    imports: []
    options:
        go_package: github.com/veith/doit-specs/dist/pb/auth;authpb
        java_multiple_files: "true"
        java_outer_classname: AuthProto
        java_package: com.furo.baseauth
```
### Field `package` *string*
This is the desired proto package name. 

In proto this is optional, in specs this field is **not optional**

> You can add an optional package specifier to a .proto file to prevent name clashes between protocol message types.

[Learn more about packages in protobuf](https://developers.google.com/protocol-buffers/docs/proto3#packages)

### Field `targetfile` *string*
The name of the proto file. Multiple types can write to the same file, if they belong to the same package.
The file will be generated in a folder according to the package.

A package `company.groups` with a protofile `green.proto` and a config target `dist/proto` will be generated to 

`dist/proto/company/groups/green.proto`

### Field `imports` *[]string*
Add imports by hand only if you do not work with Furo chain. Furo will check and fix imports for you.

This can be done with the command `furo checkImports`. Imports that can not be found would be reported.

Imports that are not needed anymore, are removed too.

```yaml
    imports:
        - google/type/date.proto
        - google/type/timeofday.proto
```

### field options *map<string, string>*
When needed, you can add options for your protos. This can be something like the following:

```yaml
    options:
        go_package: github.com/veith/doit-specs/dist/pb/auth;authpb
        java_multiple_files: "true"
        java_outer_classname: AuthProto
        java_package: com.furo.baseauth
```


{{< hint danger >}}
**Potential pitfall:**
The value of an option is a string, so write values which should have true as value like this:

`java_multiple_files: "true"`

When you write true not as string, the value will not work.
{{< /hint >}}



## Fields *map<string, Field>*
The most important part of a type spec are the fields. The fields section contains a map with fields. A field itself
has the properties `type`, `description`, `meta`, `constraints`, `__ui` and `__proto`.

```yaml
    password:
        type: string
        description: The password.
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: auth.Credentials.password.label
            placeholder: xx.xx.xx.xx
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: "true"
                message: password is required
```

### Field `type` *string*
The type of the field. This should be one of the types that you have defined or installed.

### Field `description` *string*
It is a good practice to give a good description of the type. This description will go to the generated protos and other generates.

### Field `meta` *Meta*
In the meta field you can set additional information for field in your type. Only the field *meta.repeated* does impact the proto. The options that you set here
are thought for the backend and the client. 
This properties are domain specific and can give instructions for *generators*, *validators* or *displaying* the field.

```yaml
 meta:
    default: "1234"
    hint: "look at the post-it on your monitor or below the keyboard"
    label: auth.Credentials.password.label
    placeholder: xx.xx.xx.xx
    options:
        flags: []
        list: []
    readonly: false
    repeated: false
    typespecific: null
```
#### `default` *string*
The default value for the field when you create a new object. Keep in mind that this is a string and should be parsed by your implementation.

#### `hint` *string*
This property gives you a hint message. On the furo client libs, this property goes through the translation engine first [optional]. The `furo-data-xxx-input` components will display 
this value below the field, when you focus it.

#### `label` *string*
This property labels the field. On the furo client libs, this property goes through the translation engine first [optional]. The `furo-data-xxx-input` components will display
this value as placeholder (as long the field is empty) and or label (as soon you have some value) 

#### `placeholder` *string*
This property labels the field. On the furo client libs, this property goes through the translation engine first [optional]. The `furo-data-ui5-xxx` components will display
this value as placeholder (as long the field is empty).

#### `options.flags` *[]string*

#### `options.list`  *[]Anything*

#### `readonly` *bool*
Define the field as readonly. The furo client libs will not send this field on a request by default.

#### `repeated` *bool*
Define the field as repeated. Keep in mind that not all combinations are possible. 
As an exapmle, if you set *oneof* in __proto, repeated must be set to false. 

#### typespecific: *Anything*
**Deprecated**

This is something like a extension point for fields in types. Use the extensionpoint on the field please.

### Field `constraints` *map<string, Constraint>*
Static constraints are defined on a per field basis. The client libs would use this information to mark the fields in a form
which are not valid. This does not mean that you should not check the data on the server side. 

{{< hint warning >}}
**Side note:**

The constraint may be overwritten by a server response (with `meta.fieldname.constraints:{...}`). 
{{< /hint >}}

```yaml
    constraints:
      required: #<-- constraint
        is: "true" #<-- value as string
        message: password is required #<-- message to write on constraint violation
      max:
        is: "100"
        message: not more then 100
      min:
        is: "10"
        message: at least 10
      step:
        is: "5"
        message: incrase by 5 only
```

You have to parse the *is* field according to your type when you want to use the spec directly. The client libs sets the constraints to the according input fields where they are expected as string. So nothing is to do there. If you write extended validators for the client, you have to parse the *is* value too.

### Field `__proto:` *Fieldproto*
Define the field id (proto number) and set a oneof group if needed. 

[The oneof property in detail](/docs/overview/oneof/) 

```yaml
    __proto:
        number: 1
        oneof: ""
```

### Field `__ui:` *Uiprops*
In this property you will define ui relevant attributes of a field. This information is used by some generators.
The idea behind is that you can give hints for your generators. 

The generator @furo/ui-builder use the component property to 
generate the input for the field with an explicit component, if this property is not set, it will look for the best matching input component. 
For a string i.e. a text input will be selected. In the example below we know that we have bigger texts and request a textarea to be used. 

```yaml
  __ui: {
    component: furo-data-textarea-input
    flags: 
      - full
    no_init: false  
```

It depends on your generator what you have to fill in the properties. When your generator requires concrete component names, then you have to 
write it so. When your generator can handle your "intention" then a flag *big* or *lot* would be all you have to set.

The property `no_init` tells the generator to not build something for this input.

