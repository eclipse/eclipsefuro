package genEsModule

import (
	"encoding/json"
	"fmt"
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
