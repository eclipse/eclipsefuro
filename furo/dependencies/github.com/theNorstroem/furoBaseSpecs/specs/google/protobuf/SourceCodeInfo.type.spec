name: SourceCodeInfo
type: SourceCodeInfo
description: |-
    Encapsulates information about the original source file from which a
     FileDescriptorProto was generated.
__proto:
    package: google.protobuf
    targetfile: descriptor.proto
    imports: []
    options:
        cc_enable_arenas: "true"
        csharp_namespace: Google.Protobuf.Reflection
        go_package: github.com/golang/protobuf/protoc-gen-go/descriptor;descriptor
        java_outer_classname: DescriptorProtos
        java_package: com.google.protobuf
        objc_class_prefix: GPB
fields:
    location:
        type: google.protobuf.SourceCodeInfo.Location
        description: |-
            A Location identifies a piece of source code in a .proto file which
             corresponds to a particular definition.  This information is intended
             to be useful to IDEs, code indexers, documentation generators, and similar
             tools.

             For example, say we have a file like:
               message Foo {
                 optional string foo = 1;
               }
             Let's look at just the field definition:
               optional string foo = 1;
               ^       ^^     ^^  ^  ^^^
               a       bc     de  f  ghi
             We have the following locations:
               span   path               represents
               [a,i)  [ 4, 0, 2, 0 ]     The whole field definition.
               [a,b)  [ 4, 0, 2, 0, 4 ]  The label (optional).
               [c,d)  [ 4, 0, 2, 0, 5 ]  The type (string).
               [e,f)  [ 4, 0, 2, 0, 1 ]  The name (foo).
               [g,h)  [ 4, 0, 2, 0, 3 ]  The number (1).

             Notes:
             - A location may refer to a repeated field itself (i.e. not to any
               particular index within it).  This is used whenever a set of elements are
               logically enclosed in a single code segment.  For example, an entire
               extend block (possibly containing multiple extension definitions) will
               have an outer location whose path refers to the "extensions" repeated
               field without an index.
             - Multiple locations may have the same path.  This happens when a single
               logical declaration is spread out across multiple places.  The most
               obvious example is the "extend" block again -- there may be multiple
               extend blocks in the same scope, each of which will have the same path.
             - A location's span is not always a subset of its parent's span.  For
               example, the "extendee" of an extension declaration appears at the
               beginning of the "extend" block and is shared by all extensions within
               the block.
             - Just because a location's span is a subset of some other location's span
               does not mean that it is a descendant.  For example, a "group" defines
               both a type and a field in a single declaration.  Thus, the locations
               corresponding to the type and field and their components will overlap.
             - Code which tries to interpret locations should probably be designed to
               ignore those that it doesn't understand, as more types of locations could
               be recorded in the future.
        __proto:
            number: 1
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.SourceCodeInfo.location
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
