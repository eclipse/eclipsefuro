name: MessageOptions
type: MessageOptions
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
    message_set_wire_format:
        type: bool
        description: |-
            Set true to use the old proto1 MessageSet wire format for extensions.
             This is provided for backwards-compatibility with the MessageSet wire
             format.  You should not use this for any other reason:  It's less
             efficient, has fewer features, and is more complicated.

             The message must be defined exactly as follows:
               message Foo {
                 option message_set_wire_format = true;
                 extensions 4 to max;
               }
             Note that the message cannot have any defined fields; MessageSets only
             have extensions.

             All extensions of your type must be singular messages; e.g. they cannot
             be int32s, enums, or repeated messages.

             Because this is an option, the above two restrictions are not enforced by
             the protocol compiler.
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
            label: label.MessageOptions.message_set_wire_format
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    no_standard_descriptor_accessor:
        type: bool
        description: |-
            Disables the generation of the standard "descriptor()" accessor, which can
             conflict with a field of the same name.  This is meant to make migration
             from proto1 easier; new code should avoid fields named "descriptor".
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
            label: label.MessageOptions.no_standard_descriptor_accessor
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
            Is this message deprecated?
             Depending on the target platform, this can emit Deprecated annotations
             for the message, or it will be completely ignored; in the very least,
             this is a formalization for deprecating messages.
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
            label: label.MessageOptions.deprecated
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    map_entry:
        type: bool
        description: |-
            Whether the message is an automatically generated map entry type for the
             maps field.

             For maps fields:
                 map<KeyType, ValueType> map_field = 1;
             The parsed descriptor looks like:
                 message MapFieldEntry {
                     option map_entry = true;
                     optional KeyType key = 1;
                     optional ValueType value = 2;
                 }
                 repeated MapFieldEntry map_field = 1;

             Implementations may choose not to generate the map_entry=true message, but
             use a native map in the target language to hold the keys and values.
             The reflection APIs in such implementations still need to work as
             if the field is a repeated message field.

             NOTE: Do not set the option in .proto files. Always use the maps syntax
             instead. The option should only be implicitly set by the proto compiler
             parser.
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
            label: label.MessageOptions.map_entry
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
            label: label.MessageOptions.uninterpreted_option
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
