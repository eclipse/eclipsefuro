name: Meeter
version: ""
description: The greeting service definition.
lifecycle: null
__proto:
    package: helloworld
    targetfile: helloworldservice.proto
    imports:
        - google/api/annotations.proto
        - helloworld/helloworld.proto
    options:
        go_package: github.com/veith/puregrpc/dist/pb/helloworld;helloworldpb
        java_multiple_files: "true"
        java_outer_classname: HelloworldserviceProto
        java_package: com.example.tutorial.helloworld
services:
    SayHello:
        description: Sends a greeting.
        data:
            request: helloworld.Request
            response: helloworld.Reply
            bodyfield: ""
        deeplink:
            description: ""
            href: ""
            method: ""
            rel: ""
        query: null
        rpc_name: SayHello
    SayHelloREST:
        description: Sends a greeting.
        data:
            request: helloworld.Request
            response: helloworld.Reply
            bodyfield: sub_message
        deeplink:
            description: |
                relname: This is a comment and should appear on the SayHelloREST.deeplink.description field
            href: /api/hi
            method: POST
            rel: relname
        query:
            xxx:
                type: string
                description: The request message containing the user's name.
                __proto:
                    number: 1
                __ui: null
                meta:
                    default: ""
                    placeholder: ""
                    hint: ""
                    label: label.Stuff.xxx
                    options:
                        flags: []
                        list: []
                    readonly: false
                    repeated: true
                    typespecific: null
                constraints: {}
            name:
                type: string
                description: ""
                __proto:
                    number: 4
                    oneof: test_oneof
                __ui: null
                meta:
                    default: ""
                    placeholder: ""
                    hint: ""
                    label: label.Stuff.name
                    options:
                        flags: []
                        list: []
                    readonly: false
                    repeated: false
                    typespecific: null
                constraints: {}
            oname:
                type: string
                description: |-
                    davor
                    Da geht noch was
                __proto:
                    number: 2
                    oneof: ogher_oneof
                __ui: null
                meta:
                    default: ""
                    placeholder: ""
                    hint: ""
                    label: label.Stuff.oname
                    options:
                        flags: []
                        list: []
                    readonly: false
                    repeated: false
                    typespecific: null
                constraints: {}
            osub_message:
                type: helloworld.Reply
                description: ""
                __proto:
                    number: 24
                    oneof: ogher_oneof
                __ui: null
                meta:
                    default: ""
                    placeholder: ""
                    hint: ""
                    label: label.Stuff.osub_message
                    options:
                        flags: []
                        list: []
                    readonly: false
                    repeated: false
                    typespecific: null
                constraints: {}
        rpc_name: SayHelloREST
    SayHiREST:
        description: Sends a greeting.
        data:
            request: ""
            response: helloworld.Reply
            bodyfield: ""
        deeplink:
            description: |
                This is a comment and should appear on the SayHelloREST.deeplink.description field
                The relname was not set
            href: /api/hi/{me}:custom
            method: GET
            rel: self
        query:
            name:
                type: string
                description: The request message containing the user's name.
                __proto:
                    number: 1
                __ui: null
                meta:
                    default: ""
                    placeholder: ""
                    hint: ""
                    label: label.Request.name
                    options:
                        flags: []
                        list: []
                    readonly: false
                    repeated: false
                    typespecific: null
                constraints: {}
        rpc_name: SayHiREST
