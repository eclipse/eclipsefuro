package microenums

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/enumAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/orderedmap"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/viper"
	"path"
	"regexp"
	"strconv"
	"strings"
)

type MicroEnumlist struct {
	MicroEnumsByName    map[string]*MicroEnum
	MicroEnumsASTByName map[string]*MicroEnumAst
	MicroEnums          []*MicroEnum `yaml:"enums"`
}

func (l *MicroEnumlist) UpateTypelist(typelist *enumAst.Enumlist, deleteSpecs bool, overwriteSpecOptions bool) {
	// build list to delete specs which are not enums.yaml
	deleteList := map[string]bool{}
	for typeName, _ := range typelist.EnumsByName {
		// mark every item as deletable
		deleteList[typeName] = true
	}

	for typename, mType := range l.MicroEnumsASTByName {
		deleteList[typename] = false
		// create type on Typelist if it does not exist
		if typelist.EnumsByName == nil {
			typelist.EnumsByName = map[string]*enumAst.EnumAst{}
		}

		AstType, ok := typelist.EnumsByName[typename]
		if !ok {
			typelist.EnumsByName[typename] = &enumAst.EnumAst{
				Path:     mType.TargetPath,
				FileName: mType.Type + ".enum.spec",
				EnumSpec: specSpec.Enum{},
			}
			AstType = typelist.EnumsByName[typename]
		}

		AstType.EnumSpec.Type = mType.Type
		AstType.EnumSpec.Description = mType.Description
		if AstType.EnumSpec.XProto == nil {
			AstType.EnumSpec.XProto = &specSpec.Enumproto{
				Imports:    []string{},
				Options:    map[string]string{},
				Package:    "",
				Targetfile: "",
			}
		}

		AstType.EnumSpec.XProto.Package = mType.Package
		AstType.EnumSpec.XProto.Targetfile = mType.Target
		AstType.EnumSpec.XProto.AllowAlias = mType.AllowAlias
		// check for empty options
		if AstType.EnumSpec.XProto.Options == nil || overwriteSpecOptions {
			AstType.EnumSpec.XProto.Options = map[string]string{}
		}
		// set option only if it does not exist
		_, ok = AstType.EnumSpec.XProto.Options["go_package"]
		if !ok {
			AstType.EnumSpec.XProto.Options["go_package"] = util.GetGoPackageName(mType.TargetPath)
		}
		_, ok = AstType.EnumSpec.XProto.Options["java_package"]
		if !ok {
			AstType.EnumSpec.XProto.Options["java_package"] = viper.GetString("muSpec.javaPackagePrefix") + mType.Package
		}
		_, ok = AstType.EnumSpec.XProto.Options["java_outer_classname"]
		if !ok {
			AstType.EnumSpec.XProto.Options["java_outer_classname"] = strings.Title(strings.Replace(path.Base(mType.Target), ".proto", "Proto", 1))
		}
		_, ok = AstType.EnumSpec.XProto.Options["java_multiple_files"]
		if !ok {
			AstType.EnumSpec.XProto.Options["java_multiple_files"] = "true"
		}

		valueDeleteList := map[string]bool{}
		if AstType.EnumSpec.Values != nil {
			AstType.EnumSpec.Values.Map(func(iKey interface{}, iValue interface{}) {
				valueDeleteList[iKey.(string)] = true
			})
		}
		for pair := mType.Values.Oldest(); pair != nil; pair = pair.Next() {
			mFieldname := pair.Key.(string)

			// check for values create if they do not exist and update later
			if AstType.EnumSpec.Values == nil {
				AstType.EnumSpec.Values = orderedmap.New()
			}

			// remove field from deletelist
			valueDeleteList[mFieldname] = false

			if !ok {

			}

			// Assign to Node again
			AstType.EnumSpec.Values.Set(mFieldname, pair.Value)
		}

		for fieldname, del := range valueDeleteList {
			if del {
				AstType.EnumSpec.Values.Delete(fieldname)
			}
		}
	}
	// delete the item
	for typename, del := range deleteList {
		if del {
			if deleteSpecs {
				typelist.DeleteType(typename)
				fmt.Println(typename, "deleted")
			}

		}
	}
}

// holds a single type from microspec
type MicroEnum struct {
	Enum       string                 `yaml:"enum"`
	Values     *orderedmap.OrderedMap `yaml:"values,omitempty"`
	Target     string                 `yaml:"target,omitempty"`
	SourceFile string                 `yaml:"_,omitempty"`
	AllowAlias bool                   `yaml:"alias,omitempty"`
}

func (mt MicroEnum) ToMicroEnumAst() *MicroEnumAst {

	// make the fieldmaps
	values := orderedmap.New() //was map[string]FieldMap{}
	for pair := mt.Values.Oldest(); pair != nil; pair = pair.Next() {
		valuestring := pair.Value.(string)
		fieldName := pair.Key.(string)
		intval, _ := strconv.Atoi(valuestring)
		values.Set(fieldName, intval)
	}

	// parse title and description
	regex := regexp.MustCompile(`(?s)^([^#(]*):?([^#]*)?(#(.*))?$`)
	matches := regex.FindStringSubmatch(mt.Enum)
	if len(matches) == 0 {
		fmt.Println("typeline not parseable", mt.Enum)
	}

	typedescription := ""
	if matches[4] != "" {
		typedescription = strings.TrimSpace(matches[4])
	}
	typeName := strings.TrimSpace(matches[1])
	typeArr := strings.Split(typeName, ".")
	targetpath := strings.ToLower(typeArr[0])
	packagename := strings.ToLower(typeArr[0])
	targetname := strings.ToLower(typeArr[0]) + ".proto"

	if len(typeArr) > 1 {
		// last segement is typename
		typeName = strings.TrimSpace(typeArr[len(typeArr)-1])
		// the other segments are the package
		packagename = strings.Join(typeArr[:len(typeArr)-1], ".")
		targetpath = strings.Join(typeArr[:len(typeArr)-1], "/")
		// target file name
		targetname = typeArr[len(typeArr)-2] + ".proto"
		if mt.Target != "" {
			// if optional target was given
			targetname = mt.Target
		}
	}

	mAst := MicroEnumAst{
		Type:        typeName,
		Package:     packagename,
		TargetPath:  targetpath,
		Description: typedescription,
		Values:      values,
		Target:      targetname,
		SourceFile:  mt.SourceFile,
		AllowAlias:  mt.AllowAlias,
	}

	return &mAst
}

type MicroEnumAst struct {
	Type         string                 `yaml:"type"`
	ProtoImports []string               `yaml:"imports"`
	Package      string                 `yaml:"package,omitempty"`
	TargetPath   string                 // to find out the file to write
	Description  string                 `yaml:"description"`
	Values       *orderedmap.OrderedMap `yaml:"values,omitempty"`
	Target       string                 `yaml:"target,omitempty"`
	AllowAlias   bool                   `yaml:"alias,omitempty"`
	SourceFile   string
}

// updates a type ast
func (mAst MicroEnumAst) UpdateEnumAst(ast enumAst.EnumAst) {

}

func (mAst MicroEnumAst) Save() {

}
