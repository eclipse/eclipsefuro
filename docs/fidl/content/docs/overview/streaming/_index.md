---
title: "Streaming"
weight: 100
# bookFlatSection: false
# bookToc: true
# bookHidden: false
# bookComments: true
bookCollapseSection: true
---

# Streaming

You can do streaming by adding the keyword **stream** to your response type. 

{{< hint info >}}
Streaming is only supported as response type (Server streaming RPC or in simle words: from server to the client) at the moment.
{{< /hint >}}

{{< hint warning >}}
You need at least [**furo 1.24.1**](https://furo.pro/) to use this feature 
{{< /hint >}}

Read more about [grpc streaming here](https://grpc.io/docs/what-is-grpc/core-concepts/#server-streaming-rpc)

#### Streaming big files to the client
```yaml
- md: 'Get: GET /files/{fid} google.protobuf.Empty , stream google.api.HttpBody #Returns the raw file'
    qp:
      fid: 'string #The query param fid stands for the id of a file.'
```
With google.api.HttpBody you can send any content to the client. Read more about [google.api.HttpBody here](https://github.com/googleapis/googleapis/blob/master/google/api/httpbody.proto). 

#### Stream a message type
```yaml
- md: 'Get: GET /messagestream google.protobuf.Empty, stream message.Message #Returns a stream of messages'
```
You also can stream a message type to the client.

{{< hint info >}}
The response will be packed in a attribute "result",  
{{< /hint >}}



Please read the [documentation on the grpc-gateway page](https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_your_gateway/) for configuring your gateway, or have a look on the [example subsite](/docs/overview/streaming/grpc_gateway/.)
