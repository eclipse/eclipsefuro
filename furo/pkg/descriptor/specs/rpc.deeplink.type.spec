name: servicedeeplink
type: Servicedeeplink
description: URL information for the service
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
    description:
        type: string
        description: Describe the query params
        __proto:
            number: 1
        __ui:
            component: furo-data-textarea-input
            flags:
                - full
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: ""
            label: description
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    href:
        type: string
        description: The link pattern, like /api/xxx/{qp}/yyy
        __proto:
            number: 4
        __ui:
            component: ""
            flags:
                - double
                - condensed
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: The link pattern, like /api/xxx/{qp}/yyy
            label: href
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints:
            required:
                is: "true"
                message: ""
    method:
        type: string
        description: method of curl
        __proto:
            number: 3
        __ui:
            component: furo-data-collection-dropdown
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: Http vebs like PUT, PATCH,...
            label: Method
            options:
                flags: []
                list:
                    - '@type': descriptor.Optionitem
                      display_name: GET
                      id: GET
                      selected: true
                    - '@type': descriptor.Optionitem
                      display_name: PUT
                      id: PUT
                      selected: false
                    - '@type': descriptor.Optionitem
                      display_name: POST
                      id: POST
                      selected: false
                    - '@type': descriptor.Optionitem
                      display_name: PATCH
                      id: PATCH
                      selected: false
                    - '@type': descriptor.Optionitem
                      display_name: DELETE
                      id: DELETE
                      selected: false
                    - '@type': descriptor.Optionitem
                      display_name: OPTIONS
                      id: OPTIONS
                      selected: false
                    - '@type': descriptor.Optionitem
                      display_name: HEAD
                      id: HEAD
                      selected: false
                    - '@type': descriptor.Optionitem
                      display_name: TRACE
                      id: TRACE
                      selected: false
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    rel:
        type: string
        description: the relationship
        __proto:
            number: 2
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            placeholder: ""
            hint: like create, update, custommethod
            label: rel
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}


