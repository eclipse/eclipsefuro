name: fieldmeta
type: FieldMeta
description: Metas for a field
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
    default:
        type: string
        description: The default value as JSON string, this must match with the data structure of your type
        __proto:
            number: 3
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: Default value
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    placeholder:
        type: string
        description: Informative text which can be displayed on the ui
        __proto:
            number: 8
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: Usualy used as placeholder text on empty elements
            label: Hint
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    hint:
        type: string
        description: A hint
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: Also used for input-fields
            label: Hint
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    label:
        type: string
        description: Informative text which can be displayed on the ui
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: Usualy used to label the input field
            label: Label
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    options:
        type: descriptor.Fieldoption
        description: Fieldoptions
        __proto:
            number: 6
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: options
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    readonly:
        type: bool
        description: Marks the field as readonly. Used for validations.
        __proto:
            number: 4
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: readonly
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    repeated:
        type: bool
        description: repeated
        __proto:
            number: 5
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: repeated
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    typespecific:
        type: google.protobuf.Any
        description: Put in type specific metas for your fields here
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
            label: typespecific meta
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
