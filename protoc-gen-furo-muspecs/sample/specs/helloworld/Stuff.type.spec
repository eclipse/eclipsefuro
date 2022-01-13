name: Stuff
type: Stuff
description: Stuuuuufff ...
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
    xxx:
        type: string
        description: The request message containing the user's name.
        __proto:
            number: 1
        __ui: null
        meta:
            default: ""
            placeholder: helloworld.stuff.xxx.placeholder
            hint: ""
            label: helloworld.stuff.xxx.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
    name:
        type: string
        description: no description
        __proto:
            number: 4
            oneof: test_oneof
        __ui: null
        meta:
            default: ""
            placeholder: helloworld.stuff.name.placeholder
            hint: ""
            label: helloworld.stuff.name.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    sub_message:
        type: helloworld.Request
        description: no description
        __proto:
            number: 9
            oneof: test_oneof
        __ui: null
        meta:
            default: ""
            placeholder: helloworld.stuff.submessage.placeholder
            hint: ""
            label: helloworld.stuff.submessage.label
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
            placeholder: helloworld.stuff.oname.placeholder
            hint: ""
            label: helloworld.stuff.oname.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    osub_message:
        type: helloworld.Reply
        description: no description
        __proto:
            number: 24
            oneof: ogher_oneof
        __ui: null
        meta:
            default: ""
            placeholder: helloworld.stuff.osubmessage.placeholder
            hint: ""
            label: helloworld.stuff.osubmessage.label
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
