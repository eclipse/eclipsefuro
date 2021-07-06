name: GeneratedCodeInfo
type: GeneratedCodeInfo
description: |-
    Describes the relationship between generated code and its original source
     file. A GeneratedCodeInfo message is associated with only one generated
     source file, but may contain references to different source .proto files.
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
    annotation:
        type: google.protobuf.GeneratedCodeInfo.Annotation
        description: |-
            An Annotation connects some span of text in generated code to an element
             of its generating .proto file.
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
            label: label.GeneratedCodeInfo.annotation
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
