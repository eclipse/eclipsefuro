name: typeproto
type: Typeproto
description: Main proto for a type
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
    package:
        type: string
        description: the package this type belogs to
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: package
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    targetfile:
        type: string
        description: the target proto file for this type
        __proto:
            number: 3
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: targetfile
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    imports:
        type: string
        description: needed imports like [ "spec/descriptor.proto", "google/protobuf/empty.proto" ]
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: imports
            options: null
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    options:
        type: map<string,string>
        description: 'Proto options Todo: find a solution for boolean options'
        __proto:
            number: 4
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: Proto options
            options: null
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
