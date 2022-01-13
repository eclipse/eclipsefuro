name: Request
type: Request
description: Request ...
lifecycle: null
__proto:
    package: helloworld
    targetfile: helloworld.proto
    imports: []
    options:
        go_package: '{{.RepositoryName}}/dist/pb/helloworld;helloworldpb'
        java_multiple_files: "true"
        java_outer_classname: HelloworldProto
        java_package: com.example.tutorial.helloworld
fields:
    name:
        type: string
        description: The request message containing the user's name.
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: helloworld.request.name.placeholder
            hint: ""
            label: helloworld.request.name.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
