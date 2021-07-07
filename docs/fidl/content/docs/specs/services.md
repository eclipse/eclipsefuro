---
title: "Services"
---

# Anatomy of a service spec
A service spec consists of 2 main sections (services and __proto) and some properties on the root node.


## name
The type of the service without the package name. 
Follow the [style guide](/docs/overview/style_guide/#services) and write it in CamelCase with an initial capital.


## version
Put any version information of the service. This field does not affect the URL of the services/methods.

This property has just informative character at the moment.

## description
Describe the intention of your service in some sentences. You can ommit this field. 
Furo will add a default description in the generates “developer was to lazy to give a description”.

## lifecycle
Lifecycle information for the service. If you set to true, put in a info with an alternative solution and maybe a deadline.

```yaml
lifecycle:
    deprecated: false
    info: This version is still valid
```

## __proto
The proto section defines some properties to generate the proto files. 

```yaml
__proto:
  package: fruit
  targetfile: fruitservice.proto
  imports:
    - google/api/annotations.proto
    - fruit/reqmsgs.proto
    - google/protobuf/empty.proto
    - fruit/fruit.proto
    - google/protobuf/field_mask.proto
  options:
    go_package: github.com/veith/doit-specs/dist/pb/fruit;fruitpb
    java_multiple_files: "true"
    java_outer_classname: FruitserviceProto
    java_package: com.furo.basefruit
```
### Field `package` *string*
This is the desired proto package name, his field is **not optional**


### Field `imports` *[]string*
The imports are checked by Furo (missing imports will be added).
Additional imports will not be removed (maybe you have to do a import for side effects).

This can be done with the command `furo checkImports` and is done in some other commands too.
Imports that can not be found would be reported. So you can check for typos like `fruit.FruitCollections` which should be a `fruit.FruitCollection`.

```bash

specs/Fruitservice.service.spec :Import fruit.FruitCollections not found in Service FruitService on param ListFruitService

```
### field options *map<string,string>*
When needed, you can add options for your protos. This can be something like the following:

```yaml
  options:
    go_package: github.com/veith/doit-specs/dist/pb/fruit;fruitpb
    java_multiple_files: "true"
    java_outer_classname: FruitserviceProto
    java_package: com.furo.basefruit
```

## services *map<string, service>*
The most important part of a service spec are the services, what a surprise. 
The service section contains a map with services. The key is the "name" of the service and usualy something 
like *List*, *Get*, *Delete*, *Create*, *Update*, *CustomName*. This names are not the rpc_name.

A service itself has the properties  `description`, `data`, `deeplink`, `query` and `rpc_name`.

**a single service in detail**
```yaml
  List:
    description: List fruits with pagination.
    data:
      request: google.protobuf.Empty
      response: fruit.FruitCollection
      bodyfield: body
    deeplink:
      description: 'List: GET /fruits google.protobuf.Empty , fruit.FruitCollection #List fruits with pagination.'
      href: /fruits
      method: GET
      rel: list
    query:
      q:
        description: Use this to search for a fruit.
        type: string
      filter:
        description: Use this field to filter the fruits, this is not searching.
        type: string
      order_by:
        description: Use this field to specify the ordering.
        type: string
      page:
        description: Use this field to specify page to display.
        type: string
    rpc_name: ListFruits
```

### Field `description` *string*
It is a good practice to give a good description of the type. This description will go to the generated protos and other generates.

### Field `data` *Servicereqres*
Data contains the repuest and response types for the service.

```yaml
    data:
      request: google.protobuf.Empty
      response: fruit.FruitCollection
      bodyfield: body
```

#### `request` *string*
Define the request type. Use google.protobuf.Empty if not needed.

#### `response` *string*
Define the response type. Use google.protobuf.Empty if not needed.

#### `bodyfield` *string* 
This defines the body field in request the type
The name of the request field whose value is mapped to the HTTP request
body, or `*` for mapping all request fields not captured by the path
pattern to the HTTP body, or omitted for not having any HTTP request body.

NOTE: the referred field must be present at the top-level of the request message type.

### Field `deeplink` *Servicedeeplink*
General URL Paht information for the service.

**deeplink for GetFruit**
```yaml
    deeplink:
      description: 'Get: GET /fruits/{fr} google.protobuf.Empty , fruit.FruitEntity #Returns a single fruit.'
      href: /fruits/{fr}
      method: GET
      rel: self
```
#### `description` *string*
Writing something like 'Get: GET /fruits/{fr} google.protobuf.Empty , fruit.FruitEntity #Returns a single fruit.' is a good idea.
Give additional information for the dev. if needed.

#### `href` *string*
The URL pattern with placeholders, like `/fruits/{fr}` the service is listening to . 

The placeholders are defined in the request types.

#### `method` *string*
The request method/verb the service is listening to.

This should be one of the following verbs: 
- GET
- PUT 
- PATCH
- POST
- DELETE


{{< hint info >}}
**Side note:**
If you set the verb to PUT and add a field *update_mask* of *type google.protobuf.FieldMask* to the request type `update_mask: 'google.protobuf.FieldMask #Needed to patch a record'` , an additional binding for the PATCH verb will be created in the proto.
Because it is assumable that your service has patch and put capabilities.
{{< /hint >}}

#### `rel` *string*
Give an according relation type for your link. For GET on entities this is usualy a **self** and on collections a **list**.
For DELETE a rel **delete** is set. And on custom methods it is mostly the name of the custom method. 

Use lowercase for the rel.

### Field `query` *Queryparam*
The query params for this service. This fields are used by the client lib to proove the capabilities of the service.
Furo will update this list for you, when you come from µSpec. In near future this will be removed, because the information is already
available in the request type and must not be written twice.

A query param consist of a descritpion for the documentation and a type.

```yaml
    query:
      q:
        description: Use this to search for a fruit.
        type: string
      filter:
        description: Use this field to filter the fruits, this is not searching.
        type: string
```

{{< hint warning >}}
The types and their values **must** be url safe. They will appear in the query string of the request. 
{{< /hint >}}

### Field `rpc_name` *string*
The rpc name which should appear in the proto. In the example below you can see that the rpc_name was set to *CreateFruit*

```proto
service FruitService {

  // Use this to create new fruits.
  rpc CreateFruit (CreateFruitRequest) returns (google.protobuf.Empty){
	//Create: POST /fruits fruit.Fruit , google.protobuf.Empty #Use this to create new fruits.
	option (google.api.http) = {
		post: "/fruits"
		body: "body"
	};
  }
}

```

## Example of a "Complete" Service

```yaml

name: FruitService
version: ""
description: |
  Fruits are healthy, so having a service which can list some fruits would be nice.
  We do not cover all fruits, but some. The list will grow with time, hopefully.
lifecycle: null
__proto:
  package: fruit
  targetfile: fruitservice.proto
  imports:
    - google/api/annotations.proto
    - fruit/reqmsgs.proto
    - google/protobuf/empty.proto
    - fruit/fruit.proto
    - google/protobuf/field_mask.proto
  options:
    go_package: github.com/veith/doit-specs/dist/pb/fruit;fruitpb
    java_multiple_files: "true"
    java_outer_classname: FruitserviceProto
    java_package: com.furo.basefruit
services:
  List:
    description: List fruits with pagination.
    data:
      request: google.protobuf.Empty
      response: fruit.FruitCollection
      bodyfield: body
    deeplink:
      description: 'List: GET /fruits google.protobuf.Empty , fruit.FruitCollection #List fruits with pagination.'
      href: /fruits
      method: GET
      rel: list
    query:
      q:
        description: Use this to search for a fruit.
        type: string
      filter:
        description: Use this field to filter the fruits, this is not searching.
        type: string
      order_by:
        description: Use this field to specify the ordering.
        type: string
      page:
        description: Use this field to specify page to display.
        type: string
    rpc_name: ListFruits
  Get:
    description: Returns a single fruit.
    data:
      request: google.protobuf.Empty
      response: fruit.FruitEntity
      bodyfield: body
    deeplink:
      description: 'Get: GET /fruits/{fr} google.protobuf.Empty , fruit.FruitEntity #Returns a single fruit.'
      href: /fruits/{fr}
      method: GET
      rel: self
    query:
      fr:
        description: The query param fr stands for FR id.
        type: string
    rpc_name: GetFruit
  Create:
    description: Use this to create new fruits.
    data:
      request: fruit.Fruit
      response: google.protobuf.Empty
      bodyfield: body
    deeplink:
      description: 'Create: POST /fruits fruit.Fruit , google.protobuf.Empty #Use this to create new fruits.'
      href: /fruits
      method: POST
      rel: create
    query: { }
    rpc_name: CreateFruit
  Update:
    description: Use this to update existing fruits.
    data:
      request: fruit.Fruit
      response: fruit.FruitEntity
      bodyfield: body
    deeplink:
      description: 'Update: PUT /fruits/{fr} fruit.Fruit , fruit.FruitEntity #Use this to update existing fruits.'
      href: /fruits/{fr}
      method: PUT
      rel: update
    query:
      fr:
        description: fr string.
        type: string
      update_mask:
        description: Needed to patch a record
        type: google.protobuf.FieldMask
    rpc_name: UpdateFruit
  Delete:
    description: Use this to delete existing fruits.
    data:
      request: google.protobuf.Empty
      response: google.protobuf.Empty
      bodyfield: body
    deeplink:
      description: 'Delete: DELETE /fruits/{fr} google.protobuf.Empty , google.protobuf.Empty #Use this to delete existing fruits.'
      href: /fruits/{fr}
      method: DELETE
      rel: delete
    query:
      fr:
        description: fr string.
        type: string
    rpc_name: DeleteFruit
  DeleteAll:
    description: Use this to delete ALL fruits.
    data:
      request: google.protobuf.Empty
      response: google.protobuf.Empty
      bodyfield: body
    deeplink:
      description: 'DeleteAll: DELETE /fruits google.protobuf.Empty , google.protobuf.Empty #Use this to delete ALL fruits.'
      href: /fruits
      method: DELETE
      rel: deleteall
    query: { }
    rpc_name: DeleteAllFruits
  Ferment:
    description: Fermented fruits tastes very good in liquid form.
    data:
      request: google.protobuf.Empty
      response: google.protobuf.Empty
      bodyfield: body
    deeplink:
      description: 'Ferment: POST /fruits/{fr}:ferment google.protobuf.Empty , google.protobuf.Empty #Custom methods are always POST.'
      href: /fruits/{fr}:ferment
      method: POST
      rel: ferment
    query:
      fr:
        description: fr is the placeholder for the fruit id.
        type: string
    rpc_name: FermentFruit


```