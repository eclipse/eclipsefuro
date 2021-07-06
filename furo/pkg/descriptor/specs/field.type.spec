name: field
type: Field
description: Defines a field in the furo spec
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
        description: the field type, https://developers.google.com/protocol-buffers/docs/proto3#scalar
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: Use a scalar type or a already defined one
            label: type
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    description:
        type: string
        description: the field description
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
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    __proto:
        type: descriptor.Fieldproto
        description: information for the proto generator, like number, type
        __proto:
            number: 6
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: proto
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    __ui:
        type: descriptor.Uiextension
        description: ""
        __proto:
            number: 7
        __ui: null
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
    meta:
        type: descriptor.FieldMeta
        description: meta information for the client, like label, default, repeated, options...
        __proto:
            number: 3
        __ui: null
        meta: null
        constraints: {}
    constraints:
        type: map<string,descriptor.FieldConstraint>
        description: constraints for a field, like min{}, max{}, step{}
        __proto:
            number: 4
        __ui: null
        meta: null
        constraints: {}
    extensions:
        type: map<string,google.protobuf.Any>
        description: Custom extension
        __proto:
            number: 5
        __ui: null
        meta: null
        constraints: {}
