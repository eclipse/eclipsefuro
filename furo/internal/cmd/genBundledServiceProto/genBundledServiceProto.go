package genBundledServiceProto

import (
	"bytes"
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/serviceAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/ast/typeAst"
	"github.com/eclipse/eclipsefuro/furo/pkg/protoTemplates"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"html/template"
	"io/ioutil"
	"log"
	"path"
	"sort"
	"strings"
)

type singleServiceTplData struct {
	Services             []specSpec.Service `json:"services,omitempty"`
	Imports              []string           `json:"imports,omitempty"`
	Package              string             `json:"package,omitempty"`
	preImport            map[string]bool
	Options              map[string]string `json:"options,omitempty"`
	GenAdditionalBinding bool
}

type allServicesTplData struct {
	AllServices map[string]*singleServiceTplData
	AllImports  []string
	Package     string
	Options     map[string]string
}

func Run(cmd *cobra.Command, args []string) {
	fmt.Println("genBundledServiceProto called")

	// Check for config
	c := viper.GetString("build.bundledservice.targetFile")
	if c == "" {
		log.Fatal("Config not found. Make sure you have a build.bundledservice section in your .furo config")
	}

	// todo implement flag
	withInstalled := true

	allServices := map[string]*specSpec.Service{}
	// types are needed for import checks
	Typelist := &typeAst.Typelist{}
	Typelist.LoadInstalledTypeSpecsFromDir(util.GetDependencyList()...)
	Typelist.LoadTypeSpecsFromDir(viper.GetString("specDir"))

	Servicelist := &serviceAst.Servicelist{}
	Servicelist.LoadInstalledServiceSpecsFromDir(util.GetDependencyList()...)
	Servicelist.LoadServiceSpecsFromDir(viper.GetString("specDir"))

	for k, t := range Servicelist.ServicesByName {
		allServices[k] = &t.ServiceSpec
	}
	if withInstalled {
		for k, t := range Servicelist.InstalledServicesByName {
			allServices[k] = &t.ServiceSpec
		}
	}

	// prepare templating
	fn := protoTemplates.GetSprigFuncs()
	tmpl, templateError := template.New("proto").Funcs(fn).Parse(protoTemplates.BundledServiceTemplate)
	if templateError != nil {
		log.Fatal(templateError)
	}

	protoTplDataServicelist := map[string]*singleServiceTplData{}

	// collect all types that belongs to one file
	for typeName, serviceData := range allServices {
		// typename is not the filename
		s := strings.Split(typeName, ".")
		filepath := strings.Join(s[:len(s)-1], "/") + "/" + serviceData.XProto.Targetfile

		if protoTplDataServicelist[filepath] == nil {
			protoTplDataServicelist[filepath] = &singleServiceTplData{
				preImport: map[string]bool{},
			}
		}

		protoTplDataServicelist[filepath].Services = append(protoTplDataServicelist[filepath].Services, *serviceData)
		protoTplDataServicelist[filepath].Package = serviceData.XProto.Package
		protoTplDataServicelist[filepath].Options = serviceData.XProto.Options

		// pre imports
		for _, imp := range serviceData.XProto.Imports {
			protoTplDataServicelist[filepath].preImport[imp] = true
		}
		// sort services by name
		sort.Slice(protoTplDataServicelist[filepath].Services, func(i, j int) bool {
			return protoTplDataServicelist[filepath].Services[i].Name < protoTplDataServicelist[filepath].Services[j].Name
		})

		// check if additional bindings should be created
		// if the request type has a field update_mask and method  PUT => set to true
		protoTplDataServicelist[filepath].GenAdditionalBinding = false

		for _, s := range protoTplDataServicelist[filepath].Services {
			r, found := s.Services.Get("Update")
			if found {
				rpc := r.(*specSpec.Rpc)
				reqType := protoTplDataServicelist[filepath].Package + "." + rpc.RpcName + "FuroGrpcRqst"
				// Services.sampleservice.UpdateSampleRequest
				_, ok := Typelist.AllAvailabeTypes[reqType].TypeSpec.Fields.Get("update_mask")
				if ok {
					protoTplDataServicelist[filepath].GenAdditionalBinding = true
				}

			}
		}

	}

	// process ipmports

	importmap := map[string]bool{} // is used to collect all import once

	for _, tplData := range protoTplDataServicelist {
		for imp, _ := range tplData.preImport {
			tplData.Imports = append(tplData.Imports, imp)
			importmap[imp] = true
		}
	}

	// collect ALL imports
	allData := allServicesTplData{
		AllServices: protoTplDataServicelist,
		AllImports:  []string{},
		Package:     viper.GetString("build.bundledservice.package"),
		Options:     viper.GetStringMapString("build.bundledservice.options"),
	}
	for i, _ := range importmap {
		allData.AllImports = append(allData.AllImports, i)
	}

	// sort the imports
	sort.Strings(allData.AllImports)

	// make the  proto
	filename := path.Join(viper.GetString("build.bundledservice.targetFile"))
	// create target dir => kann optimiert werden
	util.MkdirRelative(path.Dir(filename))

	var buff bytes.Buffer
	// build proto from tpl
	err := tmpl.Execute(&buff, &allData)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, buff.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
