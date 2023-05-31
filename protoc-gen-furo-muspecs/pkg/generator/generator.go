package generator

import (
	"github.com/eclipse/eclipsefuro/furo/pkg/microenums"
	"github.com/eclipse/eclipsefuro/furo/pkg/microservices"
	"github.com/eclipse/eclipsefuro/furo/pkg/microtypes"
	"github.com/eclipse/eclipsefuro/furo/pkg/orderedmap"
	"github.com/eclipse/eclipsefuro/protoc-gen-furo-muspecs/pkg/protoast"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"gopkg.in/yaml.v3"
	"path"
	"strconv"
	"strings"
)

func Generate(protoAST *protoast.ProtoAST) error {

	// this is used to resolve the query params for the services
	typeMap := map[string]protoast.MessageInfo{}
	for _, descriptor := range protoAST.ProtoMap {
		si := protoast.GetSourceInfo(descriptor)
		for i, message := range descriptor.MessageType {
			typeMap[*descriptor.Package+"."+*message.Name] = si.Messages[i]
		}
	}

	for protofilename, descriptor := range protoAST.FileProtoMap {
		si := protoast.GetSourceInfo(descriptor)
		enumArr := []protoast.EnumInfo{}

		for i, _ := range descriptor.EnumType {
			enumArr = append(enumArr, si.Enums[i])
		}

		var fileName *string
		servicesInFile := []*microservices.MicroService{}
		typesInFile := []*microtypes.MicroType{}
		enumsInFile := []*microenums.MicroEnum{}
		// generate all the services

		for ServiceIndex, Service := range descriptor.Service {
			if shouldGenerateServiceSpec(protoAST, *Service.Name, descriptor, Service) {
				fn := strings.Replace(protofilename, ".proto", ".services.yaml", 1)
				fileName = &fn
				SourceInfo := protoast.GetSourceInfo(descriptor)
				//serviceSpecFileName := evaluateServiceSpecFileName(Service, descriptor)
				serviceSpecPackageName := evaluateServiceSpecPackageName(Service, descriptor)

				description := serviceSpecPackageName + " does not have a description"

				if SourceInfo.Services[ServiceIndex].Info.LeadingComments != nil {
					description = cleanDescription(*SourceInfo.Services[ServiceIndex].Info.LeadingComments)
				}

				serviceSpec := &microservices.MicroService{
					Name:        *Service.Name,
					Description: description,
					Package:     strings.Join(strings.Split(path.Dir(protofilename), "/"), "."),
					Target:      path.Base(protofilename),
					Methods:     getServices(SourceInfo.Services[ServiceIndex], SourceInfo, typeMap),
				}
				servicesInFile = append(servicesInFile, serviceSpec)

			}

		}
		if fileName != nil {
			// append the response
			var responseFile pluginpb.CodeGeneratorResponse_File
			responseFile.Name = fileName
			content, _ := yaml.Marshal(servicesInFile)
			s := string(content)
			responseFile.Content = &s

			protoAST.Response.File = append(protoAST.Response.File, &responseFile)
			fileName = nil
		}

		// generate all the messages
		for MessageIndex, Message := range descriptor.MessageType {
			if shouldGenerateTypeSpec(protoAST, *Message.Name, descriptor, Message) {

				_, packagename := FileAndPackageNameToGenerate(descriptor, Message)
				fn := strings.Replace(protofilename, ".proto", ".types.yaml", 1)
				fileName = &fn

				typesInFile = extractMessageType(descriptor, packagename, MessageIndex, protofilename, Message, typesInFile)

				// nested enums
				//for i, _ := range typeMap[*descriptor.Package+"."+*Message.Name].Message.EnumType {
				//	inline := protoast.GetSourceInfo(descriptor).InlineEnums
				//	v := inline[i]
				//	enumArr = append(enumArr, v)
				//}
			}

			typesInFile = buildNestedMessages(typesInFile, path.Base(protofilename), strings.Join(strings.Split(path.Dir(protofilename), "/"), ".")+"."+*Message.Name, Message.NestedType)

		}
		// nested enums
		inline := protoast.GetSourceInfo(descriptor).InlineEnums
		if len(inline) > 0 {
			enumArr = append(enumArr, inline...)
		}

		if fileName != nil {
			// append the response
			var responseFile pluginpb.CodeGeneratorResponse_File
			responseFile.Name = fileName
			content, _ := yaml.Marshal(typesInFile)
			s := string(content)
			responseFile.Content = &s

			protoAST.Response.File = append(protoAST.Response.File, &responseFile)
			fileName = nil
		}

		for _, Enum := range enumArr {

			fn := strings.Replace(protofilename, ".proto", ".enums.yaml", 1)
			fileName = &fn
			description := " does not have a description"
			if Enum.Info.LeadingComments != nil {
				description = cleanDescription(*Enum.Info.LeadingComments)
			}

			typeLine := []string{}
			typeLine = append(typeLine, strings.Join(strings.Split(path.Dir(protofilename), "/"), ".")+"."+Enum.Name)

			typeLine = append(typeLine, "#"+description)

			enumSpec := &microenums.MicroEnum{
				Enum:       strings.Join(typeLine, " "),
				Values:     getEnumValues(Enum),
				Target:     "ENUM_" + path.Base(protofilename),
				AllowAlias: Enum.AllowAlias,
			}

			enumsInFile = append(enumsInFile, enumSpec)
		}
		if fileName != nil {
			// append the response
			var responseFile pluginpb.CodeGeneratorResponse_File
			responseFile.Name = fileName
			content, _ := yaml.Marshal(enumsInFile)
			s := string(content)
			responseFile.Content = &s

			protoAST.Response.File = append(protoAST.Response.File, &responseFile)
			fileName = nil
		}
	}

	return nil
}

func buildNestedMessages(file []*microtypes.MicroType, target string, prefix string, nestedType []*descriptorpb.DescriptorProto) []*microtypes.MicroType {
	for _, message := range nestedType {
		typeLine := []string{}
		typeLine = append(typeLine, prefix+"."+*message.Name)
		typeLine = append(typeLine, "# ")

		typeSpec := &microtypes.MicroType{
			Type:   strings.Join(typeLine, " "),
			Fields: getNestedFields(message.Field),
			Target: target,
		}

		file = append(file, typeSpec)
		file = buildNestedMessages(file, target, prefix+"."+*message.Name, message.NestedType)
	}

	return file
}

func getNestedFields(fields []*descriptorpb.FieldDescriptorProto) *orderedmap.OrderedMap {
	omap := orderedmap.New()
	for _, f := range fields {
		fieldline := []string{}

		// set repeated, must be false on maps!
		// repeated is in f.Field.Label
		isRepeated := false
		if *f.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
			isRepeated = true
		}

		if isRepeated {
			fieldline = append(fieldline, "[]")
		}
		fieldinfo := &protoast.FieldInfo{Field: f}
		fieldline = append(fieldline, extractTypeNameFromField(fieldinfo)+":"+strconv.Itoa(int(*f.Number)))

		fieldline = append(fieldline, "#")

		field := strings.Join(fieldline, " ")
		omap.Set(*f.Name, field)
	}
	return omap
}

func extractMessageType(descriptor *descriptorpb.FileDescriptorProto, packagename string, MessageIndex int, protofilename string, Message *descriptorpb.DescriptorProto, typesInFile []*microtypes.MicroType) []*microtypes.MicroType {
	SourceInfo := protoast.GetSourceInfo(descriptor)
	description := packagename + " does not have a description"

	if SourceInfo.Messages[MessageIndex].Info.LeadingComments != nil {
		description = cleanDescription(*SourceInfo.Messages[MessageIndex].Info.LeadingComments)
	}
	typeLine := []string{}
	typeLine = append(typeLine, *descriptor.Package+"."+*Message.Name)

	typeLine = append(typeLine, "#"+description)

	typeSpec := &microtypes.MicroType{
		Type:   strings.Join(typeLine, " "),
		Fields: getFields(SourceInfo.Messages[MessageIndex]),
		Target: path.Base(protofilename),
	}

	typesInFile = append(typesInFile, typeSpec)

	return typesInFile
}

func getEnumValues(info protoast.EnumInfo) *orderedmap.OrderedMap {
	om := orderedmap.New()
	for _, e := range info.ValuesInfo {
		om.Set(e.Name, e.Value)
	}
	return om
}

func cleanDescription(s string) string {
	res := s[1 : len(s)-1]
	strings.Replace(s, "\n", "\\n", -1)
	return res
}

func getFields(messageInfo protoast.MessageInfo) *orderedmap.OrderedMap {
	omap := orderedmap.New()
	for _, f := range messageInfo.FieldInfos {

		fielddescription := ""
		if f.Info.LeadingComments != nil {
			fielddescription = cleanDescription(*f.Info.LeadingComments)
		}

		if f.Info.TrailingComments != nil {
			fielddescription = fielddescription + "\n" + cleanDescription(*f.Info.TrailingComments)
		}
		fieldline := []string{}

		// set repeated, must be false on maps!
		// repeated is in f.Field.Label
		isRepeated := false
		if *f.Field.Label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
			isRepeated = !strings.HasPrefix(extractTypeNameFromField(&f), "map<")
		}

		if isRepeated {
			fieldline = append(fieldline, "[]")
		}

		fieldline = append(fieldline, extractTypeNameFromField(&f)+":"+strconv.Itoa(int(*f.Field.Number)))

		if f.Field.OneofIndex != nil {
			fieldline = append(fieldline, "["+*f.Message.OneofDecl[*f.Field.OneofIndex].Name+"]")
		}

		fieldline = append(fieldline, "#"+fielddescription)

		field := strings.Join(fieldline, " ")

		omap.Set(f.Name, field)
	}

	return omap
}

func extractTypeNameFromField(fieldinfo *protoast.FieldInfo) string {
	// If type_name is set, this need not be set.  If both this and type_name
	// are set, this must be one of TYPE_ENUM, TYPE_MESSAGE or TYPE_GROUP.
	// --> Type *FieldDescriptorProto_Type `protobuf:"varint,5,opt,name=type,enum=google.protobuf.FieldDescriptorProto_Type" json:"type,omitempty"`
	// For message and enum types, this is the name of the type.  If the name
	// starts with a '.', it is fully-qualified.  Otherwise, C++-like scoping
	// rules are used to find the type (i.e. first the nested types within this
	// message are searched, then within the parent, on up to the root
	// namespace).
	// --> TypeName *string `protobuf:"bytes,6,opt,name=type_name,json=typeName" json:"type_name,omitempty"`

	// get primitive types first
	// vendor/google.golang.org/protobuf/types/specSpec/descriptor.pb.go Line 54
	field := fieldinfo.Field

	if field.Type != nil {
		t := field.Type.String()
		if !(*field.Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE ||
			*field.Type == descriptorpb.FieldDescriptorProto_TYPE_ENUM ||
			*field.Type == descriptorpb.FieldDescriptorProto_TYPE_GROUP) {
			return strings.ToLower(t[5:len(t)])
		}
		// if we have message, we look in Typename
		if *field.Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
			// check for nested type map<string,xxx>
			if fieldinfo.Message.NestedType == nil {
				// must be type
				f := *field.TypeName
				return f[1:len(f)]
			}
			for _, nested := range fieldinfo.Message.NestedType {
				if nested.Options != nil {
					if *nested.Options.MapEntry {
						if strings.Title(fieldinfo.Name)+"Entry" == *nested.Name {
							// this is a map
							maptype := "not_evaluated"
							if !(*nested.Field[1].Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE ||
								*nested.Field[1].Type == descriptorpb.FieldDescriptorProto_TYPE_ENUM ||
								*nested.Field[1].Type == descriptorpb.FieldDescriptorProto_TYPE_GROUP) {
								t := nested.Field[1].Type.String()
								maptype = strings.ToLower(t[5:len(t)])
							} else {
								// can be a message or a primitive
								if *nested.Field[1].Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
									// message
									m := *nested.Field[1].TypeName
									maptype = m[1:len(m)]
								}
							}
							return "map<string," + maptype + ">"
						}
					}
				}
			}

			// check for nested types and rewrite them to package.typename_nested_name

			f := *field.TypeName
			return f[1:]

		}

		// inline enum
		if *field.Type == descriptorpb.FieldDescriptorProto_TYPE_ENUM {
			// name is messagename_enumnname

			parts := strings.Split(*field.TypeName, ".")
			mi := len(parts) - 2
			fi := len(parts) - 1
			parts[mi] = parts[mi] + "." + parts[fi]
			f := strings.Join(parts[:len(parts)-1], ".")
			return f[1:]
		}
	}

	return "unknown"
}

// get all known options
func getProtoOptions(options *descriptorpb.FileOptions) map[string]string {
	opts := map[string]string{}

	if options.JavaPackage != nil {
		opts[strcase.ToSnake("JavaPackage ")] = *options.JavaPackage //= {*string | 0xc000011c80} "pro.furo.bigdecimal"
	}

	if options.JavaOuterClassname != nil {
		opts[strcase.ToSnake("JavaOuterClassname")] = *options.JavaOuterClassname // = {*string | 0xc000011c90} "BigdecimalProto"
	}

	if options.JavaMultipleFiles != nil {
		if *options.JavaMultipleFiles {
			opts[strcase.ToSnake("JavaMultipleFiles")] = "true" // = {*bool | 0xc00001cc5f} true
		} else {
			opts[strcase.ToSnake("JavaMultipleFiles")] = "false" // = {*bool | 0xc00001cc5f} true
		}
	}

	if options.JavaStringCheckUtf8 != nil {
		if *options.JavaStringCheckUtf8 {
			opts[strcase.ToSnake("JavaStringCheckUtf8")] = "true" // = {*bool} nil
		} else {
			opts[strcase.ToSnake("JavaStringCheckUtf8")] = "false" // = {*bool | 0xc00001cc5f} true
		}
	}
	if options.GoPackage != nil {
		opts[strcase.ToSnake("GoPackage")] = *options.GoPackage // = {*string | 0xc000011ca0} "github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo/bigdecimal;bigdecimalpb"
	}
	if options.CcGenericServices != nil {
		if *options.CcGenericServices {
			opts[strcase.ToSnake("CcGenericServices")] = "true" // = {*bool} nil
		} else {
			opts[strcase.ToSnake("CcGenericServices")] = "false" // = {*bool | 0xc00001cc5f} true
		}
	}
	if options.JavaGenericServices != nil {
		if *options.JavaGenericServices {
			opts[strcase.ToSnake("JavaGenericServices")] = "true" // = {*bool} nil
		} else {
			opts[strcase.ToSnake("JavaGenericServices")] = "false" // = {*bool | 0xc00001cc5f} true
		}
	}
	if options.PyGenericServices != nil {
		if *options.PyGenericServices {
			opts[strcase.ToSnake("PyGenericServices")] = "true" // = {*bool} nil
		} else {
			opts[strcase.ToSnake("PyGenericServices")] = "false" // = {*bool | 0xc00001cc5f} true
		}
	}
	if options.PhpGenericServices != nil {
		if *options.PhpGenericServices {
			opts[strcase.ToSnake("PhpGenericServices")] = "true" // = {*bool} nil
		} else {
			opts[strcase.ToSnake("PhpGenericServices")] = "false" // = {*bool | 0xc00001cc5f} true
		}
	}
	if options.Deprecated != nil {
		if *options.Deprecated {
			opts[strcase.ToSnake("Deprecated")] = "true" // = {*bool} nil
		} else {
			opts[strcase.ToSnake("Deprecated")] = "false" // = {*bool | 0xc00001cc5f} true
		}
	}
	if options.CcEnableArenas != nil {
		if *options.CcEnableArenas {
			opts[strcase.ToSnake("CcEnableArenas")] = "true" // = {*bool | 0xc00001cc60} true
		} else {
			opts[strcase.ToSnake("CcEnableArenas")] = "false" // = {*bool | 0xc00001cc5f} true
		}
	}

	if options.ObjcClassPrefix != nil {
		opts[strcase.ToSnake("ObjcClassPrefix")] = *options.ObjcClassPrefix // = {*string | 0xc000011cb0} "FPB"
	}
	if options.CsharpNamespace != nil {
		opts[strcase.ToSnake("CsharpNamespace")] = *options.CsharpNamespace // = {*string | 0xc000011cc0} "Furo.Bigdecimal"
	}
	if options.SwiftPrefix != nil {
		opts[strcase.ToSnake("SwiftPrefix")] = *options.SwiftPrefix // = {*string} nil
	}
	if options.PhpClassPrefix != nil {
		opts[strcase.ToSnake("PhpClassPrefix")] = *options.PhpClassPrefix // = {*string} nil
	}
	if options.PhpNamespace != nil {
		opts[strcase.ToSnake("PhpNamespace")] = *options.PhpNamespace // = {*string} nil
	}
	if options.PhpMetadataNamespace != nil {
		opts[strcase.ToSnake("PhpMetadataNamespace")] = *options.PhpMetadataNamespace // = {*string} nil
	}
	if options.RubyPackage != nil {
		opts[strcase.ToSnake("RubyPackage")] = *options.RubyPackage // = {*string} nil
	}
	return opts
}
