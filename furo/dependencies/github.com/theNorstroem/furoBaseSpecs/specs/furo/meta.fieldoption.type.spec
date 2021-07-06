name: fieldoption
type: Fieldoption
description: Metas for a field
__proto:
    package: furo
    targetfile: furo.proto
    imports:
        - google/protobuf/any.proto
    options:
        cc_enable_arenas: "true"
        csharp_namespace: Furo
        go_package: github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo;furopb
        java_multiple_files: "true"
        java_outer_classname: FuroProto
        java_package: pro.furo
        objc_class_prefix: FPB
fields:
    list:
        type: google.protobuf.Any
        description: a list with options, use furo.optionitem or your own
        __proto:
            number: 1
            oneof: ""
        __ui:
            component: ""
            flags:
                - full
                - condensed
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    flags:
        type: string
        description: "Add flags for your field. This can be something like \"searchable\". \n//The flags can be used by generators, ui components,...\n"
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
            hint: optional flags
            label: flags
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
