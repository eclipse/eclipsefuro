package checkImports

import (
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Run(cmd *cobra.Command, args []string) {
	Typelist := &typeAst.Typelist{}
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))

	Typelist.UpdateImports()

	typeAst.Format = viper.GetString("specFormat")
	Typelist.SaveAllTypeSpecsToDir(viper.GetString("specDir"))

	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))

	Servicelist.UpdateAllImports(Typelist)

	serviceAst.Format = viper.GetString("specFormat")

	Servicelist.SaveAllServiceSpecsToDir(viper.GetString("specDir"))
}
