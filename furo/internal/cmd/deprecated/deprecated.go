package deprecated

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path"
)

func Run(cmd *cobra.Command, args []string) {
	installedTypes := map[string]interface{}{}
	allTypes := map[string]specSpec.Type{}
	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	for k, t := range Typelist.TypesByName {

		allTypes[k] = t.TypeSpec

	}

	for k, t := range Typelist.InstalledTypesByName {
		installedTypes[k] = t.TypeSpec
		allTypes[k] = t.TypeSpec
	}

	allServices := map[string]specSpec.Service{}
	installedServices := map[string]interface{}{}
	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))
	for k, s := range Servicelist.ServicesByName {

		allServices[k] = s.ServiceSpec

	}

	for k, s := range Servicelist.InstalledServicesByName {
		installedServices[k] = s.ServiceSpec
		allServices[k] = s.ServiceSpec

	}

	// check for deprecated types in the project types
	for tname, ast := range Typelist.TypesByName {
		ast.TypeSpec.Fields.Map(func(iKey interface{}, iValue interface{}) {
			f := iValue.(*specSpec.Field) //*string:1 # A * before the type means required
			if allTypes[f.Type].Lifecycle != nil && allTypes[f.Type].Lifecycle.Deprecated {
				fmt.Println(util.ScanForStringPosition(f.Type, path.Join(viper.GetString("specDir"), ast.Path, ast.FileName))+":WARNING: field", iKey.(string), "in type", tname, "uses deprecated type", f.Type)
				fmt.Println(allTypes[f.Type].Lifecycle.Info)
			}
		})
	}

	// check for deprecated types in the project services
	for sname, ast := range Servicelist.ServicesByName {
		ast.ServiceSpec.Services.Map(func(iKey interface{}, iValue interface{}) {
			method := iValue.(*specSpec.Rpc)
			if allTypes[method.Data.Request].Lifecycle != nil && allTypes[method.Data.Request].Lifecycle.Deprecated {

				fmt.Println(util.ScanForStringPosition(method.Data.Request, path.Join(viper.GetString("specDir"), ast.Path, ast.FileName))+":WARNING: request type on method", iKey.(string), "in service", sname, "uses deprecated type", method.Data.Request)
				fmt.Println(allTypes[method.Data.Request].Lifecycle.Info)

			}
			if allTypes[method.Data.Response].Lifecycle != nil && allTypes[method.Data.Response].Lifecycle.Deprecated {
				fmt.Println(util.ScanForStringPosition(method.Data.Response, path.Join(viper.GetString("specDir"), ast.Path, ast.FileName))+":WARNING: response type on method", iKey.(string), "in service", sname, "uses deprecated type", method.Data.Response)
				fmt.Println(allTypes[method.Data.Response].Lifecycle.Info)
			}

		})
	}

	// check installed services
	for sname, ast := range Servicelist.InstalledServicesByName {
		if ast.ServiceSpec.Lifecycle != nil && ast.ServiceSpec.Lifecycle.Deprecated {
			fmt.Println("WARNING: Deprecated service installed (" + sname + "). If you don't use it, you can ignore this message.")
			fmt.Println(ast.ServiceSpec.Lifecycle.Info)
		}
	}

}
