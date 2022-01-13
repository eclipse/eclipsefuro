name: SayHiRESTGreeterFuroRequest
type: SayHiRESTGreeterFuroRequest
description: request message for SayHiRESTGreeter
lifecycle: null
__proto:
    package: helloworld
    targetfile: reqmsgs.proto
    imports:
        - google/protobuf/empty.proto
    options:
        go_package: '{{.RepositoryName}}/dist/pb/helloworld;helloworldpb'
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: com.example.tutorial.helloworld
fields:
    body:
        type: .google.protobuf.Empty
        description: Body with google.protobuf.Empty
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: helloworld.sayhirestgreeterfurorequest.body.placeholder
            hint: ""
            label: helloworld.sayhirestgreeterfurorequest.body.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    name:
        type: string
        description: The request message containing the user's name.
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: helloworld.sayhirestgreeterfurorequest.name.placeholder
            hint: ""
            label: helloworld.sayhirestgreeterfurorequest.name.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
