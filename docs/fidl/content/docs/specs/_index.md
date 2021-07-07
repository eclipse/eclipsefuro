---
weight: 20
title: "Specs"
bookCollapseSection: true
---

# Specs

The specs or "standard specs" are the extended notation format of furo FIDL. 
If you are just interested in protobuf as output, consider to work with µSpecs.

The specs have extensionpoints on many places. If you need to add your extension,
you have to add them at the extension points now. In older versions of the specs this was
experimental, now they are an integral part of the specs. 

The extensionpoint are ,in simple words, just a `map<string,google.protobuf.Any>`.
Furoc gives you nice helpers to get your extensions in your generators. If you work with js based generators, the only thing that changes is the *name*.

## [Types](/docs/specs/types/)
A type spec in FIDL is compareable to a message in proto and have a lot of similarities. 

When you look at a field definition i.e. the only *new* field to the specs was `oneof`. 
The `number` field was added, because just using the `index+1` was not enough.   

There is a __proto option for the type itself, which defines the
- package
- targetfile
- imports
- options


## [Services](/docs/specs/services/)
A service spec in FIDL is compareable to a proto service definition with `option (google.api.http)` on every `rpc`.
There is only a __proto extension for the service itself, which defines the 
- package
- targetfile
- imports
- options
