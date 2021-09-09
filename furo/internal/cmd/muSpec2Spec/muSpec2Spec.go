package muSpec2Spec

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/microservices"
	"github.com/eclipse/eclipsefuro/furo/pkg/microtypes"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func Run(cmd *cobra.Command, args []string) {
	deleteSpecs := false
	f := cmd.Flag("delete")
	if f != nil {
		deleteSpecs = f.Value.String() == "true"
	}

	if viper.GetBool("muSpec.forceSync") {
		deleteSpecs = true
	}

	fmt.Println("running muSpec2Spec")

	microList := &microtypes.MicroTypelist{
		MicroTypesByName:    map[string]*microtypes.MicroType{},
		MicroTypesASTByName: map[string]*microtypes.MicroTypeAst{},
		MicroTypes:          []*microtypes.MicroType{},
	} // holds all muspecs

	microServicesList := &microservices.MicroServiceList{
		MicroServicesByName:    map[string]*microservices.MicroService{},
		MicroServicesASTByName: map[string]*microservices.MicroServiceAst{},
		MicroServices:          []*microservices.MicroService{},
	} // holds all muspecs

	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))
	Servicelist.LoadInstalledServiceSpecsFromDir(viper.GetStringSlice("importedServiceSpecs")...)

	Typelist := &typeAst.Typelist{}
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)

	serviceglobs := viper.GetStringSlice("muSpec.services")
	typeglobs := viper.GetStringSlice("muSpec.types")
	for _, glob := range typeglobs {
		list, err := filepath.Glob(glob)
		if err != nil {
			log.Fatal(err)
		}
		LoadTypes(list, microList)
	}
	for _, glob := range serviceglobs {
		list, err := filepath.Glob(glob)
		if err != nil {
			log.Fatal(err)
		}
		LoadServices(list, microServicesList)
	}

	// build the service name and ast map
	for _, t := range microServicesList.MicroServices {

		serviceName := strings.TrimSpace(t.Package) + "." + strings.TrimSpace(t.Name)
		microServicesList.MicroServicesByName[serviceName] = t
		microServicesList.MicroServicesASTByName[serviceName] = t.ToMicroServiceAst()
	}

	overwriteSpecOptions := false
	if cmd.Flag("overwrite-spec-options") != nil &&
		cmd.Flag("overwrite-spec-options").Value.String() == "true" {
		overwriteSpecOptions = true
	}

	// update the services ast from microspecs
	microServicesList.UpateServicelist(Servicelist, deleteSpecs, microList, overwriteSpecOptions) //microList is to create request types...

	// build the new name and ast map
	for _, t := range microList.MicroTypes {
		regex := regexp.MustCompile(`(?s)^([^#(]*):?([^#]*)?(#(.*))?$`)
		matches := regex.FindStringSubmatch(t.Type)
		if len(matches) == 0 {
			fmt.Println("typeline not parseable", t.Type)
		}

		typeName := strings.TrimSpace(matches[1])
		microList.MicroTypesByName[typeName] = t
		microList.MicroTypesASTByName[typeName] = t.ToMicroTypeAst()
	}

	Servicelist.UpdateAllImports(Typelist)

	// update the typelist from microspecs
	microList.UpateTypelist(Typelist, deleteSpecs, overwriteSpecOptions)
	Typelist.UpdateImports()

	// save types and services
	typeAst.Format = viper.GetString("specFormat")
	serviceAst.Format = viper.GetString("specFormat")

	Typelist.SaveAllTypeSpecsToDir(viper.GetString("specDir"))
	Servicelist.SaveAllServiceSpecsToDir(viper.GetString("specDir"))
}
