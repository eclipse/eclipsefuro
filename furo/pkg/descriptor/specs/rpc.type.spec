name: rpc
type: Rpc
description: Defines a rpc for a service
__proto:
    package: descriptor
    targetfile: descriptor.proto
    imports:
        - google/protobuf/any.proto
    options:
        cc_enable_arenas: "true"
        go_package: github.com/eclipse/eclipsefuro/furo/pkg/descriptor;descriptorpb
        java_multiple_files: "true"
        java_outer_classname: FuroDescriptorProto
        java_package: pro.furo.descriptor
        objc_class_prefix: FPB
fields:
    description:
        type: string
        description: the service description
        __proto:
            number: 1
        __ui:
            component: furo-data-textarea-input
            flags:
                - full
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: description
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    data:
        type: descriptor.Servicereqres
        description: Request and response types for the service
        __proto:
            number: 3
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: proto
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    deeplink:
        type: descriptor.Servicedeeplink
        description: This data is needed for...
        __proto:
            number: 5
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: deeplink
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    query:
        type: map<string,descriptor.Queryparam>
        description: Query params, it is recomended to use string types
        __proto:
            number: 4
        __ui: null
        meta: null
        constraints: {}
    rpc_name:
        type: string
        description: RPC name https://developers.google.com/protocol-buffers/docs/proto3#services
        __proto:
            number: 2
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: The rpc name
            label: rpc name
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    extensions:
        type: map<string,google.protobuf.Any>
        description: Custom extension
        __proto:
            number: 6
        __ui: null
        meta: null
        constraints: {}
