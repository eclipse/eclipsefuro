name: uiextension
type: Uiextension
description: ui hints for a field
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
    component:
        type: string
        description: component hint for ui-builder
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ui-builder will use this element
            label: Component hint
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: ""
                message: ""
    flags:
        type: string
        description: UI element flags like full, double, hidden,...
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: UI element flags like full, double, hidden,...
            label: Element flags
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    no_init:
        type: bool
        description: Skip adding this field on ui init
        __proto:
            number: 3
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: Skip initialization of this field in ui-builder
            label: Skip init
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    no_skip:
        type: bool
        description: do not skip this field, even it is in the default skip list
        __proto:
            number: 4
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: never skip initialization of this field in ui-builder
            label: do not skip
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
