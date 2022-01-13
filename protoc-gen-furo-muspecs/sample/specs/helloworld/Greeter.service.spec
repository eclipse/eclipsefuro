name: Greeter
version: ""
description: The greeting service definition.
lifecycle: null
__proto:
    package: helloworld
    targetfile: helloworldservice.proto
    imports:
        - google/api/annotations.proto
        - helloworld/reqmsgs.proto
        - google/protobuf/empty.proto
        - helloworld/helloworld.proto
    options:
        go_package: '{{.RepositoryName}}/dist/pb/helloworld;helloworldpb'
        java_multiple_files: "true"
        java_outer_classname: HelloworldserviceProto
        java_package: com.example.tutorial.helloworld
services:
    AAAASayHello:
        description: Sends a pure grpc greeting.
        data:
            request: helloworld.Request
            response: helloworld.Reply
            bodyfield: body
        deeplink:
            description: 'AAAASayHello:   helloworld.Request , helloworld.Reply #Sends a pure grpc greeting.'
            href: ""
            method: ""
            rel: aaaasayhello
        query: {}
        rpc_name: AAAASayHello
    SayHelloREST:
        description: Sends a greeting.
        data:
            request: '*'
            response: helloworld.Reply
            bodyfield: '*'
        deeplink:
            description: 'SayHelloREST: POST /api/hi * , helloworld.Reply #Sends a greeting.'
            href: /api/hi
            method: POST
            rel: sayhellorest
        query:
            name:
                constraints: {}
                description: The request message containing the user's name.
                meta: null
                type: string
        rpc_name: SayHelloRESTGreeter
    SayHiREST:
        description: Sends a greeting.
        data:
            request: google.protobuf.Empty
            response: helloworld.Reply
            bodyfield: body
        deeplink:
            description: 'SayHiREST: GET /api/hi google.protobuf.Empty , helloworld.Reply #Sends a greeting.'
            href: /api/hi
            method: GET
            rel: sayhirest
        query:
            name:
                constraints: {}
                description: The request message containing the user's name.
                meta: null
                type: string
        rpc_name: SayHiRESTGreeter
