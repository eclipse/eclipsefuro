name: fieldoption
type: Fieldoption
description: Options like flags and list
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
    flags:
        type: string
        description: "Add flags for your field. This can be something like \"searchable\". \n//The flags can be used by generators, ui components,...\n"
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
            hint: optional flags
            label: flags
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    list:
        type: google.protobuf.Any
        description: a list with options, use descriptor.Optionitem or your own
        __proto:
            number: 1
        __ui:
            component: ""
            flags:
                - full
                - condensed
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
            repeated: true
            typespecific: null
        constraints: {}
