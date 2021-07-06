name: FileOptions
type: FileOptions
description: descriptor does not have a description
__proto:
    package: google.protobuf
    targetfile: descriptor.proto
    imports: []
    options:
        cc_enable_arenas: "true"
        csharp_namespace: Google.Protobuf.Reflection
        go_package: github.com/golang/protobuf/protoc-gen-go/descriptor;descriptor
        java_outer_classname: DescriptorProtos
        java_package: com.google.protobuf
        objc_class_prefix: GPB
fields:
    java_package:
        type: string
        description: |-
            Sets the Java package where classes generated from this .proto will be
             placed.  By default, the proto package is used, but this is often
             inappropriate because proto packages do not normally start with backwards
             domain names.
        __proto:
            number: 1
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.java_package
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    java_outer_classname:
        type: string
        description: |-
            If set, all the classes from the .proto file are wrapped in a single
             outer class with the given name.  This applies to both Proto1
             (equivalent to the old "--one_java_file" option) and Proto2 (where
             a .proto always translates to a single class, but you may want to
             explicitly choose the class name).
        __proto:
            number: 8
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.java_outer_classname
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    java_multiple_files:
        type: bool
        description: |-
            If set true, then the Java code generator will generate a separate .java
             file for each top-level message, enum, and service defined in the .proto
             file.  Thus, these types will *not* be nested inside the outer class
             named by java_outer_classname.  However, the outer class will still be
             generated to contain the file's getDescriptor() method as well as any
             top-level extensions defined in the file.
        __proto:
            number: 10
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.java_multiple_files
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    java_generate_equals_and_hash:
        type: bool
        description: This option does nothing.
        __proto:
            number: 20
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.java_generate_equals_and_hash
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    java_string_check_utf8:
        type: bool
        description: |-
            If set true, then the Java2 code generator will generate code that
             throws an exception whenever an attempt is made to assign a non-UTF-8
             byte sequence to a string field.
             Message reflection will do the same.
             However, an extension field still accepts non-UTF-8 byte sequences.
             This option has no effect on when used with the lite runtime.
        __proto:
            number: 27
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.java_string_check_utf8
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    optimize_for:
        type: unknown
        description: ""
        __proto:
            number: 9
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.optimize_for
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    go_package:
        type: string
        description: |-
            Sets the Go package where structs generated from this .proto will be
             placed. If omitted, the Go package will be derived from the following:
               - The basename of the package import path, if provided.
               - Otherwise, the package statement in the .proto file, if present.
               - Otherwise, the basename of the .proto file, without extension.
        __proto:
            number: 11
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.go_package
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    cc_generic_services:
        type: bool
        description: |-
            Should generic services be generated in each language?  "Generic" services
             are not specific to any particular RPC system.  They are generated by the
             main code generators in each language (without additional plugins).
             Generic services were the only kind of service generation supported by
             early versions of google.protobuf.

             Generic services are now considered deprecated in favor of using plugins
             that generate code specific to your particular RPC system.  Therefore,
             these default to false.  Old code which depends on generic services should
             explicitly set them to true.
        __proto:
            number: 16
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.cc_generic_services
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    java_generic_services:
        type: bool
        description: ""
        __proto:
            number: 17
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.java_generic_services
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    py_generic_services:
        type: bool
        description: ""
        __proto:
            number: 18
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.py_generic_services
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    php_generic_services:
        type: bool
        description: ""
        __proto:
            number: 42
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.php_generic_services
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    deprecated:
        type: bool
        description: |-
            Is this file deprecated?
             Depending on the target platform, this can emit Deprecated annotations
             for everything in the file, or it will be completely ignored; in the very
             least, this is a formalization for deprecating files.
        __proto:
            number: 23
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.deprecated
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    cc_enable_arenas:
        type: bool
        description: |-
            Enables the use of arenas for the proto messages in this file. This applies
             only to generated classes for C++.
        __proto:
            number: 31
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.cc_enable_arenas
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    objc_class_prefix:
        type: string
        description: |-
            Sets the objective c class prefix which is prepended to all objective c
             generated classes from this .proto. There is no default.
        __proto:
            number: 36
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.objc_class_prefix
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    csharp_namespace:
        type: string
        description: Namespace for generated classes; defaults to the package.
        __proto:
            number: 37
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.csharp_namespace
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    swift_prefix:
        type: string
        description: |-
            By default Swift generators will take the proto package and CamelCase it
             replacing '.' with underscore and use that to prefix the types/symbols
             defined. When this options is provided, they will use this value instead
             to prefix the types/symbols defined.
        __proto:
            number: 39
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.swift_prefix
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    php_class_prefix:
        type: string
        description: |-
            Sets the php class prefix which is prepended to all php generated classes
             from this .proto. Default is empty.
        __proto:
            number: 40
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.php_class_prefix
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    php_namespace:
        type: string
        description: |-
            Use this option to change the namespace of php generated classes. Default
             is empty. When this option is empty, the package name will be used for
             determining the namespace.
        __proto:
            number: 41
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.php_namespace
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    php_metadata_namespace:
        type: string
        description: |-
            Use this option to change the namespace of php generated metadata classes.
             Default is empty. When this option is empty, the proto file name will be
             used for determining the namespace.
        __proto:
            number: 44
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.php_metadata_namespace
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    ruby_package:
        type: string
        description: |-
            Use this option to change the package of ruby generated classes. Default
             is empty. When this option is not set, the package name will be used for
             determining the ruby package.
        __proto:
            number: 45
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.ruby_package
            options:
                flags: []
                list: []
            readonly: false
            repeated: false
            typespecific: null
        constraints: {}
    uninterpreted_option:
        type: google.protobuf.UninterpretedOption
        description: |-
            The parser stores options it doesn't recognize here.
             See the documentation for the "Options" section above.
        __proto:
            number: 999
            oneof: ""
        __ui:
            component: ""
            flags: []
            noinit: false
            noskip: false
        meta:
            default: ""
            hint: ""
            label: label.FileOptions.uninterpreted_option
            options:
                flags: []
                list: []
            readonly: false
            repeated: true
            typespecific: null
        constraints: {}
