---
weight: 30
title: "Default Values"
---
# Fields And Default Values
 
The client lib will fill the "proto" default values for the fields when they are not transmited. This means i.e. numeric types will get a 0, strings a empty string. 
The default value that you can specify in the specs has nothing to do with proto directly. Proto 3 does not know anything about default values.

This comes very handy when you create an instance of a type on the client side and pass the data object to a form.

## Static Default Values
The standard spec let you define a default value. This values are entered as **string** and must be parsed by them who use them.
The client framework, does the parsing transparently for you. 

```yaml
fields:
  description:
    type: string
    description: Describe the fruit
    __proto:
      number: 3
      oneof: ""
    __ui: null
    meta:
      default: "This is the default" #<-- default values goes here!
      hint: ""
      label: fruit.Fruit.description.label
      options:
        flags: []
        list: []
      readonly: false
      repeated: false
      typespecific: null
    constraints:
      required:
        is: "true"
        message: description is required
```

### Json Object Example
It is recomended to enter the default values as json string. The javascript clients will not parse them, 
when you use yaml notation.

```yaml
fields:
  fruit:
    type: fruit.Fruit    
    meta: #the following default value is a string too, it just look like a json object.
      default: |
        {"key":"value"}        

```


## Runtime Default Values
This topic belongs more to the client libs and is here only for completenes.

You can set default value for a type on the client side during runtime with a server response via the *meta* field.

```json
{
  "data": {
    "id": "1",
    "scalar_string": "this is a scalar string"
  },
  "links": [],
  "meta": {
    "fields": {
      "data.scalar_string": {
        "meta": {
          "label": "scalar_string string label setted via response meta",
          "readonly": false,
          "default": "runtime default value"
        }
      }
    }
  }
}

```


{{< hint info >}}
There are more patterns to feed the clien with default values. It is also possible to build a custom method which creates a 
default object on the server side and feed this to the data object on the client side to create a new object. Hint: by doing this do not forget to set the HATEOAS link for the rel *create*.  
{{< /hint >}}