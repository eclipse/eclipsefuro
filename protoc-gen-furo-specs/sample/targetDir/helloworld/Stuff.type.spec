name: Stuff
type: Stuff
description: Stuuuuufff ...
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
    sub_message:
        type: helloworld.Request
        description: ""
        __proto:
            number: 9
            oneof: test_oneof
        __ui: null
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: label.Stuff.sub_message
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
