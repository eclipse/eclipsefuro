name: link
type: Link
description: link
__proto:
    package: furo
    targetfile: furo.proto
    imports: []
    options:
        cc_enable_arenas: "true"
        csharp_namespace: Furo
        go_package: github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo;furopb
        java_multiple_files: "true"
        java_outer_classname: FuroProto
        java_package: pro.furo
        objc_class_prefix: FPB
fields:
    rel:
        type: string
        description: the relationship like self...
        __proto:
            number: 1
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    method:
        type: string
        description: method of curl v1.0.0
        __proto:
            number: 2
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    href:
        type: string
        description: link
        __proto:
            number: 3
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    type:
        type: string
        description: mime type
        __proto:
            number: 4
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
    service:
        type: string
        description: name of the service which can handle this link
        __proto:
            number: 5
            oneof: ""
        __ui: null
        meta: null
        constraints: {}
