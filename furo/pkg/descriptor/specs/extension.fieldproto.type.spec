name: fieldproto
type: Fieldproto
description: Proto options for a field
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
    number:
        type: int32
        description: The field numbers are used to identify your fields in the message binary format, and should not be changed once your message type is in use.
        __proto:
            number: 2
        __ui:
            component: furo-data-number-input
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: must be unique in fields
            label: Proto field number
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: "true"
                message: a unique number is needed
    oneof:
        type: string
        description: Assign field to a protobuf oneof group.
        __proto:
            number: 3
        __ui:
            component: furo-data-text-input
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: Proto oneof group
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
