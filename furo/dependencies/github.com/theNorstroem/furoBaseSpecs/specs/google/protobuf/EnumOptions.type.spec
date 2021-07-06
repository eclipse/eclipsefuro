name: EnumOptions
type: EnumOptions
description: descriptor does not have a description
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
    allow_alias:
        type: bool
        description: |-
            Set this option to true to allow mapping different tag names to the same
             value.
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
            label: label.EnumOptions.allow_alias
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    deprecated:
        type: bool
        description: |-
            Is this enum deprecated?
             Depending on the target platform, this can emit Deprecated annotations
             for the enum, or it will be completely ignored; in the very least, this
             is a formalization for deprecating enums.
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
            label: label.EnumOptions.deprecated
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    uninterpreted_option:
        type: google.protobuf.UninterpretedOption
        description: The parser stores options it doesn't recognize here. See above.
        __proto:
            number: 999
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.EnumOptions.uninterpreted_option
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
