---
title: "GraphQL"
weight: 16
---


# How to expose a GraphQL API from furo-specs
Furo does not produce GraphQL directly.

Furo produces `proto` files. From the `proto` you have to generate a [**grpc-gateway**](https://grpc-ecosystem.github.io/grpc-gateway/) 
and the [**OAS**](https://www.openapis.org/) specs. 

The final step is to use IBMs [openapi-to-graphql](https://github.com/IBM/openapi-to-graphql) or 
if you just want an instant server to test something the [openapi-to-graphql-cli](https://github.com/IBM/openapi-to-graphql/tree/master/packages/openapi-to-graphql-cli).

<img src="/grpcio-ar21.svg" style="width: 120px;">
<span style="font-size: 80px; color:#999999">➔</span>
<img src="/openapis-icon.svg" style="width: 60px">
<span style="font-size: 80px; color:#999999">➔</span>
<img src="/graphQL.svg" style="width: 60px">


## Example Request
*collection*
<img src="/collection.png" >

*entity*
<img src="/entity.png" >
