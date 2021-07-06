package clientspec

import (
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"strings"
)

type ClientTypeList map[string]*Type
type ClientServiceList map[string]*Service

var availableTypes = ClientTypeList{} // holds all av. types
var availableServices = ClientServiceList{}

func AddTypesToResolver(tlist map[string]*typeAst.TypeAst) {
	for k, t := range tlist {
		availableTypes[k] = CreateClientTypeFromAstType(&t.TypeSpec)
	}
}

func AddServicesToResolver(slist map[string]*serviceAst.ServiceAst) {
	for fullname, t := range slist {
		availableServices[t.ServiceSpec.Name] = CreateServiceFromAstService(&t.ServiceSpec, fullname)
	}
}

func GetAllTypes() ClientTypeList {
	return availableTypes
}

func GetAllServices() ClientServiceList {
	return availableServices
}

func TransformCPlusStyleToAbsolutTypes() {

	// https://developers.google.com/protocol-buffers/docs/proto3
	// Packages and Name Resolution
	// Type name resolution in the protocol buffer language works like C++: first the innermost scope is searched, then the next-innermost, and so on, with each package considered to be "inner" to its parent package. A leading '.' (for example, .foo.bar.Baz) means to start from the outermost scope instead.

	// The protocol buffer compiler resolves all type names by parsing the imported .proto files. The code generator for each language knows how to refer to each type in that language, even if it has different scoping rules.

	// resolves the c++ notation to client notation which is always from root

	// check all fields in all types
	for fullQualifiedName, typeordermap := range availableTypes { //fullQualifiedName contains the package..

		typeordermap.Fields.Map(func(iFname interface{}, iField interface{}) {
			fieldname := iFname.(string)
			iEnvField, _ := availableTypes[fullQualifiedName].Fields.Get(fieldname)
			es6Field := iEnvField.(*Field)

			if strings.HasPrefix(es6Field.Type, ".") {
				// type starts from root, just remove the .
				es6Field.Type = es6Field.Type[1:len(es6Field.Type)]
			} else {
				// find first occurrence which can match the field

				es6Field.Type = resolveFullQualifiedTypename(es6Field.Type, fullQualifiedName)

			}

			availableTypes[fullQualifiedName].Fields.Set(fieldname, es6Field)

		})

	}

}

func resolveFullQualifiedTypename(typename string, pkg string) string {
	// absolut type given, nothing special to do
	if strings.HasPrefix(typename, ".") {
		// type starts from root, just remove the .
		return typename[1:len(typename)]
	}

	pathArr := strings.Split(pkg, ".")
	// if we are in type a.b.c.d and want type x.y we look for
	// a.b.c.d.x.y
	// a.b.c.x.y
	// a.b.x.y
	// a.x.y
	// x.y
	for i := len(pathArr) - 1; i >= 0; i-- {
		sub := strings.Join(pathArr[0:i], ".")
		ftype := sub + "." + typename

		if availableTypes[ftype] != nil {
			// match
			return ftype
			i = 0
		}
		// we are at root
		if i == 0 && strings.HasPrefix(ftype, ".") {
			// remove .
			return ftype[1:len(ftype)]
		}
	}

	return typename
}
