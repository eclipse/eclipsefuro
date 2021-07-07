---
title: "Style Guide"
weight: 300
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookCollapseSection: false
# bookComments: true
---
# Style Guide
The style guide is nearly similar to the [protocol buffers style guide](https://developers.google.com/protocol-buffers/docs/style) 
and follows also the [google api design guidelines](https://cloud.google.com/apis/design/standard_fields).

## File structure
Files should be named lower_snake_case.types.yaml. 

You can put the services and types to the same folder (thematic group). You can move the files around without any effect.
The resulting protos does not depend on the FIDL file names or structure, they will follow the given package definitions.

```bash
muspecs
├── auth
│   ├── auth.services.yaml
│   └── auth.types.yaml
└── fruits
    ├── fruit.services.yaml
    └── fruits.types.yaml


dist/protos
├── Services
│   └── auth
│       ├── auth.proto
│       └── reqmsgs.proto
├── auth
│   └── auth.proto
└── fruit
    ├── fruit.proto
    ├── fruitservice.proto
    └── reqmsgs.proto

```

## Type and field names
### package names
Use camelCase (without an initial capital) for package names.

{{< hint info >}} 
**sample**.SampleRequest
{{< /hint >}}

### type names
Use CamelCase (with an initial capital) for type names.


{{< hint info >}}
sample.**SampleRequest**
{{< /hint >}}

### field names
Use underscore_separated_names for field names for example, birth_date.

{{< hint info >}}
birth_date
{{< /hint >}}


```yaml
- type: 'sample.Sample #A sample type'
  fields:
    password: '* string:1 #The password.'
    birth_date: '* google.type.Date:2 #The username or email, or something to identify.'
    details: 'sample.Details:3 #Details.'  
```


If your field name contains a number, the number should appear after the letter instead of after the underscore. e.g., use song_name1 instead of song_name_1

#### Repeated fields
Use pluralized names for repeated fields.


{{< hint info >}}
**tags**: '[] string:1 #Some tags.'
{{< /hint >}}


## Services

You should use CamelCase (with an initial capital) for both the service name and any method names:


{{< hint info >}}
Look at **- name: FruitService** , **- md: 'ListFruits: GET /...**. You can find a good guide and explanation on the [Standard Methods](https://cloud.google.com/apis/design/standard_methods) page of the cloud API design guide page. 
{{< /hint >}} 

```yaml
- name: FruitService
  description: |
    Fruits are healthy, so having a service which can list some fruits would be nice.
    We do not cover all fruits, but some. The list will grow with time, hopefully.
  package: fruit
  target: fruit_service.proto
  methods:
    - md: 'ListFruits: GET /fruits google.protobuf.Empty , fruit.FruitCollection #Filterable and searchable list of fruits with pagination.'
      qp:
        q: 'string #Use this to search for a fruit.'
        filter: 'string #Use this field to filter the fruits, this is not searching.'
        order_by: 'string #Use this field to specify the ordering.'
        page: 'string #Use this field to specify page to display.'
    - md: 'Get: GET /fruits/{frt} google.protobuf.Empty , fruit.FruitEntity #Returns a single fruit.'
      qp:
        frt: 'string #The query param frt stands for the FRuiT id.'
```

### URL Path
Usualy the path part is a noun in plural form. Use singular nouns only on singleton ressources.
Do **not** append a prefix like */api* to your paths. Use **/fruits**.

{{< hint warning >}}
TIPP: Assume that your API is a host by its own. So you will address it with **api.xy.com/fruits**. 
Having *api.xy.com/api/fruits* will look strange in that moment.

Adding prefixes can be done by infrastructure. The furo client libs also have the posibillity to prefix your specs according
to the situation.
{{< /hint >}}

{{< hint danger >}}
You will loose portability capabilities when you prefix your paths. 
{{< /hint >}}

### URL Placeholer / Query Params
Use short names for the placeholders because they will apear on every request you make and also in the HATEOAS of every 
response. 

Use some consonants only and use them that you can recognize the word behind, when possible. 
The first letter of a word is allowed to be a vovel.

Main objects of your domain should use only two or three letters. A

{{< hint danger >}}
**q** is "reserved" for search query.
{{< /hint >}}

**some examples**
- fruit => **frt**
- example_data => **exd**
- support_cases **spc**
- support_analysis **spa**