---
weight: 40
title: "Scalar Types"
---
# Scalar Value Types

Furo specs knows the same scalar types which are defined in [ProtocolBuffers](https://developers.google.com/protocol-buffers/docs/overview#scalar).

| type 	| note 	| golang  	| java  	| es6  	|
|---	|---	|---	|---	|---	|
| `string` 	| Please use UTF-8  	|  `*string` 	|  `String` 	|  `String` |
| `bytes` 	|   	|  `[]byte` 	|  `ByteString` 	|  `String` |
| `bool` 	|   	|  `*bool` 	|  `boolean` 	|  `Boolean` |
| `float` 	|   	|  `*float32` 	|  `float` 	|  `Number` |
| `double` 	|   	|  `*float64` 	|  `double` 	|  `Number` |
| `int32` 	|   	|  `*int32` 	|  `int` 	|  `Number` |
| `int64` 	|   	|  `*int64` 	|  `long` 	|  `Number` |
| `uint32` 	|   	|  `*uint32` 	|  `int` 	|  `Number` |
| `uint64` 	|   	|  `*uint64` 	|  `long` 	|  `Number` |
| `sint32` 	|   	|  `*int32` 	|  `int` 	|  `Number` |
| `sint64` 	|   	|  `*int64` 	|  `long` 	|  `Number` |
| `fixed32` 	|   	|  `*uint32` 	|  `int` 	|  `Number` |
| `fixed64` 	|   	|  `*uint64` 	|  `long` 	|  `Number` |
| `sfixed32` 	|   	|  `*int32` 	|  `int` 	|  `Number` |
| `sfixed64` 	|   	|  `*int64` 	|  `long` 	|  `Number` |