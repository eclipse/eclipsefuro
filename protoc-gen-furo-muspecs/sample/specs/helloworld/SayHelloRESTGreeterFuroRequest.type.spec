name: SayHelloRESTGreeterFuroRequest
type: SayHelloRESTGreeterFuroRequest
description: request message for SayHelloRESTGreeter
lifecycle: null
__proto:
    package: helloworld
    targetfile: reqmsgs.proto
    imports: []
    options:
        go_package: '{{.RepositoryName}}/dist/pb/helloworld;helloworldpb'
        java_multiple_files: "true"
        java_outer_classname: ReqmsgsProto
        java_package: com.example.tutorial.helloworld
fields:
    name:
        type: string
        description: The request message containing the user's name.
        __proto:
            number: 2
        __ui: null
        meta:
            default: ""
            placeholder: helloworld.sayhellorestgreeterfurorequest.name.placeholder
            hint: ""
            label: helloworld.sayhellorestgreeterfurorequest.name.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
