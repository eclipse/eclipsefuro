package exportAsYaml

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func Run(cmd *cobra.Command, args []string) {
	fullExport := false
	// todo implement flag
	exportInstalled := true
	f := cmd.Flag("full")
	if f != nil {
		fullExport = f.Value.String() == "true"
	}

	allTypes := map[string]interface{}{}
	installedTypes := map[string]interface{}{}
	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	for k, t := range Typelist.TypesByName {
		if fullExport {
			allTypes[k] = t
		} else {
			allTypes[k] = t.TypeSpec
		}

	}

	if exportInstalled {
		for k, t := range Typelist.InstalledTypesByName {
			if fullExport {
				installedTypes[k] = t
			} else {
				installedTypes[k] = t.TypeSpec
			}
		}
	}

	allServices := map[string]interface{}{}
	installedServices := map[string]interface{}{}
	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))
	for k, s := range Servicelist.ServicesByName {
		if fullExport {
			allServices[k] = s
		} else {
			allServices[k] = s.ServiceSpec
		}

	}
	if exportInstalled {
		for k, s := range Servicelist.InstalledServicesByName {
			if fullExport {
				installedServices[k] = s
			} else {
				installedServices[k] = s.ServiceSpec
			}
		}
	}

	output := map[string]interface{}{}

	output["config"] = viper.AllSettings()
	output["installedServices"] = installedServices
	output["installedTypes"] = installedTypes
	output["services"] = allServices
	output["types"] = allTypes

	outputstr, _ := yaml.Marshal(output)

	fmt.Print(string(outputstr))
}
