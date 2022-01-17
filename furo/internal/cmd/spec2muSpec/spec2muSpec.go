package spec2muSpec

// Generate µ-specs from Specs
import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/internal/cmd/muSpec2Spec"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/enumAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/microenums"
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
	deleteMuSpecs := false
	f := cmd.Flag("delete")
	if f != nil {
		deleteMuSpecs = f.Value.String() == "true"
	}

	if viper.GetBool("muSpec.forceSync") {
		deleteMuSpecs = true
	}

	fmt.Println("running spec2muSpec")
	assocList := NewUTShadowList()

	microTypesList := &microtypes.MicroTypelist{
		MicroTypesByName:    map[string]*microtypes.MicroType{},
		MicroTypesASTByName: map[string]*microtypes.MicroTypeAst{},
		MicroTypes:          []*microtypes.MicroType{},
	} // holds all muspecs

	microEnumsList := &microenums.MicroEnumlist{
		MicroEnumsByName:    map[string]*microenums.MicroEnum{},
		MicroEnumsASTByName: map[string]*microenums.MicroEnumAst{},
		MicroEnums:          []*microenums.MicroEnum{},
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

	Enumlist := &enumAst.Enumlist{}
	Enumlist.LoadEnumSpecsFromDir(viper.GetString("specDir"))
	Enumlist.LoadInstalledEnumSpecsFromDir(util.GetDependencyList()...)

	// fill assoc
	for _, ast := range Servicelist.ServicesByName {
		assocList.AddServiceNode(ast)
	}

	// fill assoc
	for typename, ast := range Typelist.TypesByName {
		assocList.AddTypeNode(typename, ast)
	}
	// fill assoc
	for enumname, ast := range Enumlist.EnumsByName {
		assocList.AddEnumNode(enumname, ast)
	}

	serviceglobs := viper.GetStringSlice("muSpec.services")
	typeglobs := viper.GetStringSlice("muSpec.types")
	enumglobs := viper.GetStringSlice("muSpec.enums")
	for _, glob := range typeglobs {
		list, err := filepath.Glob(glob)
		if err != nil {
			log.Fatal(err)
		}
		muSpec2Spec.LoadTypes(list, microTypesList)
	}
	for _, glob := range enumglobs {
		list, err := filepath.Glob(glob)
		if err != nil {
			log.Fatal(err)
		}
		muSpec2Spec.LoadEnums(list, microEnumsList)
	}
	for _, glob := range serviceglobs {
		list, err := filepath.Glob(glob)
		if err != nil {
			log.Fatal(err)
		}
		muSpec2Spec.LoadServices(list, microServicesList)
	}

	// build the service name and ast map
	for _, t := range microServicesList.MicroServices {

		serviceName := strings.TrimSpace(t.Package) + "." + strings.TrimSpace(t.Name)
		microServicesList.MicroServicesByName[serviceName] = t
		microServicesList.MicroServicesASTByName[serviceName] = t.ToMicroServiceAst()

		assocList.AddMicroServiceNode(microServicesList.MicroServicesASTByName[serviceName])
	}

	// build the new name and ast map
	for _, t := range microTypesList.MicroTypes {
		regex := regexp.MustCompile(`(?s)^([^#(]*):?([^#]*)?(#(.*))?$`)
		matches := regex.FindStringSubmatch(t.Type)
		if len(matches) == 0 {
			fmt.Println("typeline not parseable", t.Type)
		}

		typeName := strings.TrimSpace(matches[1])
		microTypesList.MicroTypesByName[typeName] = t
		microTypesList.MicroTypesASTByName[typeName] = t.ToMicroTypeAst()
		assocList.AddMicroTypeNode(microTypesList.MicroTypesASTByName[typeName])
	}

	// build the new name and ast map
	for _, t := range microEnumsList.MicroEnums {
		regex := regexp.MustCompile(`(?s)^([^#(]*):?([^#]*)?(#(.*))?$`)
		matches := regex.FindStringSubmatch(t.Enum)
		if len(matches) == 0 {
			fmt.Println("enumline not parseable", t.Enum)
		}

		enumName := strings.TrimSpace(matches[1])
		microEnumsList.MicroEnumsByName[enumName] = t
		microEnumsList.MicroEnumsASTByName[enumName] = t.ToMicroEnumAst()
		assocList.AddMicroEnumNode(microEnumsList.MicroEnumsASTByName[enumName])
	}

	// this is the main part :-)
	updateAndStoreMicroTypes(assocList.TypeItemsByName)
	updateAndStoreMicroEnums(assocList.EnumItemsByName)
	updateAndStoreMicroServices(assocList.ServiceItemsByName)

	e := assocList.GetUnconnectedMicroTypes()
	if deleteMuSpecs && len(e) > 0 {
		fmt.Println(len(e), "muSpecs deleted")
	}

}
