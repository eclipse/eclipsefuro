package generator

import (
	"google.golang.org/protobuf/types/descriptorpb"
	"path/filepath"
	"protoc-gen-open-models/pkg/sourceinfo"
	"strings"
)

var WellKnownTypesMap = map[string]string{
	"StringValue": "string",
	"BytesValue":  "string",
	"BoolValue":   "boolean",
	"Int32Value":  "number",
	"Int64Value":  "number",
	"FloatValue":  "number",
	"DoubleValue": "number",
	"UInt32Value": "number",
	"UInt64Value": "number",
}

var ModelTypesMap = map[string]string{
	"TYPE_STRING":   "STRING",
	"TYPE_BYTES":    "BYTES",
	"TYPE_BOOL":     "BOOLEAN",
	"TYPE_INT32":    "INT32",
	"TYPE_INT64":    "INT64",
	"TYPE_DOUBLE":   "DOUBLE",
	"TYPE_FLOAT":    "FLOAT",
	"TYPE_UINT32":   "UINT32",
	"TYPE_UINT64":   "UINT64",
	"TYPE_FIXED32":  "FIXED32",
	"TYPE_FIXED64":  "FIXED64",
	"TYPE_SFIXED32": "SFIXED32",
	"TYPE_SFIXED64": "SFIXED64",
	"TYPE_SINT32":   "SINT32",
	"TYPE_SINT64":   "SINT64",
}

func resolveModelType(imports ImportMap, field sourceinfo.FieldInfo) (
	ModelType string, SetterCommand string, SetterType string, GetterType string, MapValueConstructor string, FieldConstructor string) {
	tn := field.Field.GetTypeName()

	fieldType := field.Field.Type.String()

	if t, ok := ModelTypesMap[fieldType]; ok {
		primitiveType := PrimitivesMap[fieldType]
		if field.Field.Label.String() == "LABEL_REPEATED" {
			imports.AddImport("@furo/open-models/dist/index", "ARRAY")
			return "ARRAY<" + t + ", " + primitiveType + ">",
				"__TypeSetter",
				primitiveType + "[]",
				"ARRAY<" + t + ", " + primitiveType + ">",
				"", // ARRAY is uses a typesetter
				t
		}
		imports.AddImport("@furo/open-models/dist/index", ModelTypesMap[fieldType])
		return t, "__PrimitivesSetter", primitiveType, t, "", t
	}

	// Maps
	for _, nested := range field.Message.NestedType {
		if nested.Options != nil {
			if *nested.Options.MapEntry {
				if strings.Title(field.Name)+"Entry" == *nested.Name {
					// this is a map
					maptype := "not_evaluated"
					if !(*nested.Field[1].Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE ||
						*nested.Field[1].Type == descriptorpb.FieldDescriptorProto_TYPE_ENUM ||
						*nested.Field[1].Type == descriptorpb.FieldDescriptorProto_TYPE_GROUP) {
						t := nested.Field[1].Type.String()
						maptype = t
					} else {
						// can be a message or a primitive
						if *nested.Field[1].Type == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
							// message
							m := *nested.Field[1].TypeName
							maptype = m[1:len(m)]
							// todo:implement map<string,MESSAGETYPE>
							panic("implement map<string,MESSAGETYPE>")
						}
					}
					// for model types return "MAP<string, STRING, string>;"
					imports.AddImport("@furo/open-models/dist/index", "MAP")
					imports.AddImport("@furo/open-models/dist/index", ModelTypesMap[maptype])
					return "MAP<string," + ModelTypesMap[maptype] + "," + PrimitivesMap[maptype] + ">",
						"__TypeSetter",
						"{ [key: string]: " + PrimitivesMap[maptype] + " }",
						"MAP<string," + ModelTypesMap[maptype] + "," + PrimitivesMap[maptype] + ">",
						ModelTypesMap[maptype],
						"MAP<string," + ModelTypesMap[maptype] + "," + PrimitivesMap[maptype] + ">"

				}
			}
		}
	}

	if fieldType == "TYPE_MESSAGE" {
		// WELL KNOWN
		if strings.HasPrefix(tn, ".google.protobuf.") {
			ts := strings.Split(tn, ".")
			typeName := ts[len(ts)-1]

			// ANY
			if typeName == "Any" {
				imports.AddImport("@furo/open-models/dist/index", "type IAny")
				imports.AddImport("@furo/open-models/dist/index", "ANY")
				return "ANY", "__TypeSetter", "IAny", "ANY", "", "ANY"
			}
			primitiveType := WellKnownTypesMap[typeName]
			imports.AddImport("@furo/open-models/dist/index", typeName)
			return typeName, "__TypeSetter", primitiveType + "| null", typeName, "", typeName
		}

		// MESSAGE
		t := field.Field.GetTypeName()
		if strings.HasPrefix(t, ".") {
			t = t[1:]
		}
		if strings.HasPrefix(t, field.Package) {
			// we are in the same package
			// import is just ./[TypeName]
			ss := strings.Split(field.Field.GetTypeName(), ".")
			importFile := ss[len(ss)-1]
			// do not add import for the same file (direct recursion types)
			t = fullQualifiedName(t, "")
			if field.Message.GetName() != importFile {
				imports.AddImport("./"+importFile, t)
			}

			if field.Field.Label.String() == "LABEL_REPEATED" {
				imports.AddImport("@furo/open-models/dist/index", "ARRAY")
				return "ARRAY<" + t + ", I" + t + ">",
					"__TypeSetter",
					"I" + t + "[]",
					"ARRAY<" + t + ", I" + t + ">",
					"",
					t
			}
			// if a field type equals the package name + message type we have a direct recusrion
			if field.Field.GetTypeName() == "."+field.Package+"."+field.Message.GetName() {
				imports.AddImport("@furo/open-models/dist/index", "RECURSION")
				return "RECURSION<" + t + ", I" + t + ">",
					"__TypeSetter",
					"I" + t,
					"RECURSION<" + t + ", I" + t + ">",
					"",
					t
			}
			// todo: check for deep recursion
			if deepRecursionCheck(field.Field.GetTypeName()) {
				imports.AddImport("@furo/open-models/dist/index", "RECURSION")
				return "RECURSION<" + t + ", I" + t + ">",
					"__TypeSetter",
					"I" + t,
					"RECURSION<" + t + ", I" + t + ">",
					"",
					t
			}

			return t, "__TypeSetter", "I" + t, t, "", t
		}

		// find relative path to import target,

		if _, ok := projectFiles[typenameToPath(tn)]; ok {
			// definition is in project root
			ss := strings.Split(field.Field.GetTypeName(), ".")
			importFile := ss[len(ss)-1]
			fieldPackage := strings.Split("."+field.Package, ".")
			importPackage := ss[0 : len(ss)-1]

			rel, _ := filepath.Rel(strings.Join(fieldPackage, "/"), strings.Join(importPackage, "/"))

			// do not add import for the same file (direct recursion types)
			t = fullQualifiedName(t, "")
			if field.Message.GetName() != importFile {
				imports.AddImport(rel+"/"+importFile, t)
			}
			return t, "__TypeSetter", "I" + t, t, "", t
		}

		return field.Field.GetTypeName(), "__TypeSetter", "todo:resolve dependency", "???", "", field.Field.GetTypeName()
	}
	if fieldType == "TYPE_ENUM" {
		t := field.Field.GetTypeName()
		if strings.HasPrefix(t, ".") {
			t = t[1:]
		}
		if strings.HasPrefix(t, field.Package) {
			// we are in the same package
			// import is just ./[TypeName]
			ss := strings.Split(field.Field.GetTypeName(), ".")
			importFile := ss[len(ss)-1]
			t = fullQualifiedName(t, "")

			imports.AddImport("@furo/open-models/dist/index", "ENUM")
			imports.AddImport("./"+importFile, t)
			return "ENUM<" + t + ">", "__TypeSetter", t, "ENUM<" + t + ">", "", "ENUM<" + t + ">"
		}
		return "ENUM:UNRECOGNIZED", "__TypeSetter", "???", "???", "", "ENUM<" + t + ">"
	}

	return "UNRECOGNIZED", "UNRECOGNIZED", "UNRECOGNIZED", "UNRECOGNIZED", "", "UNRECOGNIZED"
}

func deepRecursionCheck(typename string) bool {
	return deepRecursionCheckRecursion(typename, typename)
}
func deepRecursionCheckRecursion(startAt string, lookFor string) bool {

	for _, info := range allTypes[startAt].FieldInfos {
		if info.Field.GetTypeName() == lookFor {
			return true
		}
		if info.Field.Type.String() == "TYPE_MESSAGE" {
			return deepRecursionCheckRecursion(info.Field.GetTypeName(), lookFor)
		}

	}
	return false
}

func typenameToPath(tn string) string {
	return strings.Replace(tn[1:], ".", "/", -1)
}
