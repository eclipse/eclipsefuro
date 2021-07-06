name: servicereqres
type: Servicereqres
description: Repuest and response types for services, used in service.type
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
    request:
        type: string
        description: Define the request type, leave this field empty if not needed
        __proto:
            number: 1
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: leave empty if not needed
            label: request type
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    response:
        type: string
        description: Define the response type, leave this field empty if not needed
        __proto:
            number: 2
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: leave empty if not needed
            label: response type
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    bodyfield:
        type: string
        description: |
            Define the body field in request the type
            The name of the request field whose value is mapped to the HTTP request
            body, or `*` for mapping all request fields not captured by the path
            pattern to the HTTP body, or omitted for not having any HTTP request body.
            NOTE: the referred field must be present at the top-level of the request
            message type.
        __proto:
            number: 3
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta: null
        constraints: {}
