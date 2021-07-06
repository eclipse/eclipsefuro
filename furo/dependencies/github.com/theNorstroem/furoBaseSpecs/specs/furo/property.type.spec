name: property
type: Property
description: Type to define property values with type information
__proto:
    package: furo
    targetfile: property.proto
    imports:
        - google/protobuf/any.proto
        - furo/furo.proto
    options:
        cc_enable_arenas: "true"
        csharp_namespace: Furo.Property
        go_package: github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo/property;propertypb
        java_multiple_files: "true"
        java_outer_classname: PropertyProto
        java_package: pro.furo.property
        objc_class_prefix: FPB
fields:
    id:
        type: string
        description: Id of the property
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: Id
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: "true"
                message: is required
    display_name:
        type: string
        description: String representation of the property
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: Property
            options: null
            readonly: true
            repeated: false
            typespecific: null
        constraints: {}
    data:
        type: google.protobuf.Any
        description: data part of the property
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    meta:
        type: furo.Meta
        description: Meta for the response
        __proto:
            number: 4
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    code:
        type: string
        description: property code for additional settings
        __proto:
            number: 5
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    flags:
        type: string
        description: Optional attribute flags e.g. is-overwritable
        __proto:
            number: 6
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options: null
            readonly: true
            repeated: true
            typespecific: null
        constraints: {}
    is_overwritten:
        type: bool
        description: Optional flag indicating that the property differs from the original value
        __proto:
            number: 7
            oneof: ""
        __ui: null
        meta:
            default: ""
            hint: ""
            label: ""
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
