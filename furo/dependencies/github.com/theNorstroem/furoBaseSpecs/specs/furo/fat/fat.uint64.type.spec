name: uint64
type: Uint64
description: Furo annotated type wrapper message for `uint64`.  The range constraints are set to Number.MAX_SAFE_INTEGER because of browser limitations
__proto:
    package: furo.fat
    targetfile: fat.proto
    imports: []
    options:
        cc_enable_arenas: "true"
        csharp_namespace: Furo.Fat # von package
        go_package: github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo/fat;fatpb # aus .furo.module + pb + package + targetfile ohne .proto + pb
        java_multiple_files: "true"
        java_outer_classname: FatProto # von targetfile
        java_package: pro.furo.fat # tld + package (pro + furo.fat)
        objc_class_prefix: FPB # (F)uro(P)roto(B)uff
fields:
    value:
        type: uint64
        description: The JSON representation for `Uint64Value` is JSON number, range is set to 0 - Number.MAX_SAFE_INTEGER
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
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    labels:
        type: map<string,bool>
        description: Labels / flags for the value, something like unspecified, empty, confidential, absent,... Can be used for AI, UI, Business Logic,...
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
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    attributes:
        type: map<string,string>
        description: 'Attributes for a value, something like confidential-msg: you are not allowed to see this value '
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
            label: ""
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
