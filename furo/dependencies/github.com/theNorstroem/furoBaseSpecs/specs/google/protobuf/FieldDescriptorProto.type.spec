name: FieldDescriptorProto
type: FieldDescriptorProto
description: Describes a field within a message.
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
    name:
        type: string
        description: ""
        __proto:
            number: 1
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FieldDescriptorProto.name
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    number:
        type: int32
        description: ""
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
            label: label.FieldDescriptorProto.number
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    label:
        type: unknown
        description: ""
        __proto:
            number: 4
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FieldDescriptorProto.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    type:
        type: unknown
        description: |-
            If type_name is set, this need not be set.  If both this and type_name
             are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP.
        __proto:
            number: 5
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FieldDescriptorProto.type
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    type_name:
        type: string
        description: |-
            For message and enum types, this is the name of the type.  If the name
             starts with a '.', it is fully-qualified.  Otherwise, C++-like scoping
             rules are used to find the type (i.e. first the nested types within this
             message are searched, then within the parent, on up to the root
             namespace).
        __proto:
            number: 6
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FieldDescriptorProto.type_name
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    extendee:
        type: string
        description: |-
            For extensions, this is the name of the type being extended.  It is
             resolved in the same manner as type_name.
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
            label: label.FieldDescriptorProto.extendee
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    default_value:
        type: string
        description: |-
            For numeric types, contains the original text representation of the value.
             For booleans, "true" or "false".
             For strings, contains the default text contents (not escaped in any way).
             For bytes, contains the C escaped value.  All bytes >= 128 are escaped.
             TODO(kenton):  Base-64 encode?
        __proto:
            number: 7
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FieldDescriptorProto.default_value
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    oneof_index:
        type: int32
        description: |-
            If set, gives the index of a oneof in the containing type's oneof_decl
             list.  This field is a member of that oneof.
        __proto:
            number: 9
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FieldDescriptorProto.oneof_index
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    json_name:
        type: string
        description: |-
            JSON name of this field. The value is set by protocol compiler. If the
             user has set a "json_name" option on this field, that option's value
             will be used. Otherwise, it's deduced from the field's name by converting
             it to camelCase.
        __proto:
            number: 10
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FieldDescriptorProto.json_name
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    options:
        type: google.protobuf.FieldOptions
        description: ""
        __proto:
            number: 8
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FieldDescriptorProto.options
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    proto3_optional:
        type: bool
        description: |-
            If true, this is a proto3 "optional". When a proto3 field is optional, it
             tracks presence regardless of field type.

             When proto3_optional is true, this field must be belong to a oneof to
             signal to old proto3 clients that presence is tracked for this field. This
             oneof is known as a "synthetic" oneof, and this field must be its sole
             member (each proto3 optional field gets its own synthetic oneof). Synthetic
             oneofs exist in the descriptor only, and do not generate any API. Synthetic
             oneofs must be ordered after all "real" oneofs.

             For message fields, proto3_optional doesn't create any semantic change,
             since non-repeated message fields always track presence. However it still
             indicates the semantic detail of whether the user wrote "optional" or not.
             This can be useful for round-tripping the .proto file. For consistency we
             give message fields a synthetic oneof also, even though it is not required
             to track presence. This is especially important because the parser can't
             tell if a field is a message or an enum, so it must always create a
             synthetic oneof.

             Proto2 optional fields do not set this flag, because they already indicate
             optional with `LABEL_OPTIONAL`.
        __proto:
            number: 17
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FieldDescriptorProto.proto3_optional
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
