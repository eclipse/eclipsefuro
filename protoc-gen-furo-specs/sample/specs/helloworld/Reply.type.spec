name: Reply
type: Reply
description: Reply ...
lifecycle: null
__proto:
    package: helloworld
    targetfile: helloworld.proto
    imports: []
    options:
        go_package: github.com/veith/puregrpc/dist/pb/helloworld;helloworldpb
        java_multiple_files: "true"
        java_outer_classname: HelloworldProto
        java_package: com.example.tutorial.helloworld
fields:
    message:
        type: string
        description: The response message containing the greetings.
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.Reply.message
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
