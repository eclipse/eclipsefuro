name: enum
type: Enum
description: Defines a type in the furo spec
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
    type:
        type: string
        description: 'This is the Typename'
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
            hint: The typename, without package. i.e. Person
            label: type
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    description:
        type: string
        description: the type description
        __proto:
            number: 3
        __ui:
            component: furo-data-textarea-input
            flags:
                - full
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: Describe what this type is for
            label: description
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    lifecycle:
        type: descriptor.Lifecycle
        description: Type lifecycle information, setting deprecated to true will log a warning when running furo
        __proto:
            number: 7
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    __proto:
        type: descriptor.Enumproto
        description: information for the proto generator, should be removed for the client spec
        __proto:
            number: 4
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
    values:
        type: map<string,int32>
        description: Value of a enum
        __proto:
            number: 5
        __ui: null
        meta: null
        constraints: {}
    extensions:
        type: map<string,google.protobuf.Any>
        description: Custom extension
        __proto:
            number: 6
        __ui: null
        meta: null
        constraints: {}
