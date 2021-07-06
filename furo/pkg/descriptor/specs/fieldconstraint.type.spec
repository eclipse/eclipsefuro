name: fieldconstraint
type: FieldConstraint
description: a single fieldconstraint
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
    is:
        type: string
        description: the constraint value as string, even it is a number
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: the constraint value as string, even it is a number
            label: is
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    message:
        type: string
        description: The message to display on constraint violation
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: message
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
