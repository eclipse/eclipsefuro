package genEsModule

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/enumAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/clientspec"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

func Run(cmd *cobra.Command, args []string) {

	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	// ENUMS
	Enumlist := &enumAst.Enumlist{}
	Enumlist.LoadEnumSpecsFromDir(viper.GetString("specDir"))
	Enumlist.LoadInstalledEnumSpecsFromDir(util.GetDependencyList()...)

	clientspec.AddEnumsToResolver(Enumlist.EnumsByName)
	clientspec.AddTypesToResolver(Typelist.TypesByName)
	clientspec.AddTypesToResolver(Typelist.InstalledTypesByName)
	// after adding all types we can build up the type resolutions
	// in the es module we use absolute type names only
	clientspec.TransformCPlusStyleToAbsolutTypes()
	allTypes := clientspec.GetAllTypes()

	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))
	clientspec.AddServicesToResolver(Servicelist.ServicesByName)
	clientspec.AddServicesToResolver(Servicelist.InstalledServicesByName)
	allServices := clientspec.GetAllServices()

	// check for enum fields in types and add the options
	for _, t := range allTypes {
		if t.Fields != nil {
			// convert fields from yaml.Node to Field type
			for pair := t.Fields.Oldest(); pair != nil; pair = pair.Next() {
				field := pair.Value.(*clientspec.Field)

				if allTypes[field.Type] != nil && allTypes[field.Type].Values != nil {

					field.Meta.Options.Flags = append(field.Meta.Options.Flags, "enum")
					// loop all enum values to build the options
					isenum := false
					for enum := allTypes[field.Type].Values.Oldest(); enum != nil; enum = enum.Next() {
						option := map[string]interface{}{}
						option["display_name"] = strings.ToLower("enum." + field.Type + "." + enum.Key.(string) + ".label")
						option["id"] = enum.Value.(*int32)
						option["@type"] = "type.googleapis.com/furo.Optionitem"
						option["selected"] = false
						field.Meta.Options.List = append(field.Meta.Options.List, &option)
						isenum = true
					}

					// set the type on the field to enum
					if isenum {
						field.Meta.Typespecific = &map[string]interface{}{"enum": field.Type}
						field.Type = "enum"
					}
				}

			}
		}

	}

	td, _ := json.Marshal(allTypes)

	// only escape " \n and \t and not just \ because we have strings like \u003c which is a <
	escapedType := strings.ReplaceAll(string(td), "\\\"", "\\\\\"")
	escapedType = strings.ReplaceAll(escapedType, "\\n", "\\\\n")
	escapedType = strings.ReplaceAll(escapedType, "\\t", "\\\\t")

	typeLine := "export const Types = JSON.parse(\n`" + escapedType + "`,\n);\n"
	sd, _ := json.Marshal(allServices)

	escapedServices := strings.ReplaceAll(string(sd), "\\\"", "\\\\\"")
	escapedServices = strings.ReplaceAll(escapedServices, "\\n", "\\\\n")
	escapedServices = strings.ReplaceAll(escapedServices, "\\t", "\\\\t")

	serviceLine := "export const Services = JSON.parse(\n`" + escapedServices + "`,\n);\n"

	err := ioutil.WriteFile(viper.GetString("build.esModule.targetFile"), []byte("/* eslint-disable */\n"+typeLine+"\n"+serviceLine), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Environment.js written to ", viper.GetString("build.esModule.targetFile"))

}
