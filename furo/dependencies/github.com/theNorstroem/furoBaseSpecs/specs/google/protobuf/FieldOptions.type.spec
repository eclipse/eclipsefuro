name: FieldOptions
type: FieldOptions
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
    ctype:
        type: unknown
        description: |-
            The ctype option instructs the C++ code generator to use a different
             representation of the field than it normally would.  See the specific
             options below.  This option is not yet implemented in the open source
             release -- sorry, we'll try to include it in a future version!
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
            label: label.FieldOptions.ctype
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    packed:
        type: bool
        description: |-
            The packed option can be enabled for repeated primitive fields to enable
             a more efficient representation on the wire. Rather than repeatedly
             writing the tag and type for each element, the entire array is encoded as
             a single length-delimited blob. In proto3, only explicit setting it to
             false will avoid using packed encoding.
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
            label: label.FieldOptions.packed
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    jstype:
        type: unknown
        description: |-
            The jstype option determines the JavaScript type used for values of the
             field.  The option is permitted only for 64 bit integral and fixed types
             (int64, uint64, sint64, fixed64, sfixed64).  A field with jstype JS_STRING
             is represented as JavaScript string, which avoids loss of precision that
             can happen when a large value is converted to a floating point JavaScript.
             Specifying JS_NUMBER for the jstype causes the generated JavaScript code to
             use the JavaScript "number" type.  The behavior of the default option
             JS_NORMAL is implementation dependent.

             This option is an enum to permit additional types to be added, e.g.
             goog.math.Integer.
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
            label: label.FieldOptions.jstype
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    lazy:
        type: bool
        description: |-
            Should this field be parsed lazily?  Lazy applies only to message-type
             fields.  It means that when the outer message is initially parsed, the
             inner message's contents will not be parsed but instead stored in encoded
             form.  The inner message will actually be parsed when it is first accessed.

             This is only a hint.  Implementations are free to choose whether to use
             eager or lazy parsing regardless of the value of this option.  However,
             setting this option true suggests that the protocol author believes that
             using lazy parsing on this field is worth the additional bookkeeping
             overhead typically needed to implement it.

             This option does not affect the public interface of any generated code;
             all method signatures remain the same.  Furthermore, thread-safety of the
             interface is not affected by this option; const methods remain safe to
             call from multiple threads concurrently, while non-const methods continue
             to require exclusive access.


             Note that implementations may choose not to check required fields within
             a lazy sub-message.  That is, calling IsInitialized() on the outer message
             may return true even if the inner message has missing required fields.
             This is necessary because otherwise the inner message would have to be
             parsed in order to perform the check, defeating the purpose of lazy
             parsing.  An implementation which chooses not to check required fields
             must be consistent about it.  That is, for any particular sub-message, the
             implementation must either *always* check its required fields, or *never*
             check its required fields, regardless of whether or not the message has
             been parsed.
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
            label: label.FieldOptions.lazy
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
            Is this field deprecated?
             Depending on the target platform, this can emit Deprecated annotations
             for accessors, or it will be completely ignored; in the very least, this
             is a formalization for deprecating fields.
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
            label: label.FieldOptions.deprecated
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    weak:
        type: bool
        description: For Google-internal migration only. Do not use.
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
            label: label.FieldOptions.weak
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
            label: label.FieldOptions.uninterpreted_option
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
