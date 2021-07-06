name: ServiceDescriptorProto
type: ServiceDescriptorProto
description: Describes a service.
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
    name:
        type: string
        description: ""
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
            label: label.ServiceDescriptorProto.name
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    method:
        type: google.protobuf.MethodDescriptorProto
        description: ""
        __proto:
            number: 2
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.ServiceDescriptorProto.method
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    options:
        type: google.protobuf.ServiceOptions
        description: ""
        __proto:
            number: 3
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.ServiceDescriptorProto.options
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
