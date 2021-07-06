name: lifecycle
type: Lifecycle
description: Lifecycle information for a type or a service
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
    deprecated:
        type: bool
        description: Is this version deprecated
        __proto:
            number: 1
        __ui: null
        meta: null
        constraints: {}
    info:
        type: string
        description: Inform about the replacement here, if you have one
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: Inform about the replacement here, if you have one
            label: Deprecation info
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
