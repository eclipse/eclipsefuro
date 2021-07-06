name: Queryparam
type: Queryparam
description: Defines a queryparam field (for rpc type)
__proto:
    package: descriptor
    targetfile: descriptor.proto
    imports: []
    options:
        cc_enable_arenas: "true"
        go_package: github.com/eclipse/eclipsefuro/furo/pkg/descriptor;descriptorpb
        java_multiple_files: "true"
        java_outer_classname: FuroDescriptorProto
        java_package: pro.furo.descriptor
        objc_class_prefix: FPB
fields:
    constraints:
        type: map<string,descriptor.FieldConstraint>
        description: constraints for a field, like min{}, max{}, step{}. Not used at the moment
        __proto:
            number: 4
        __ui:
            component: ""
            flags:
                - full
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: constraints
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: { }
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
    meta:
        type: descriptor.FieldMeta
        description: meta information for the client, like label, default, repeated, options...
        __proto:
            number: 3
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: meta
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
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