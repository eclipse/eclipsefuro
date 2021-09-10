package spec2muSpec

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/microtypes"
	"github.com/eclipse/eclipsefuro/furo/pkg/orderedmap"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

func updateAndStoreMicroTypes(typeItems map[string]*UTshadowNode) {
	muTypesPerFile := map[string][]*microtypes.MicroType{}

	// transfer every type assocList.TypeItemsByName
	for _, shadowNode := range typeItems {
		if shadowNode.edgeTypeNode != nil {

			if shadowNode.edgeMicroTypeNode == nil {
				// create muType because it does not exist
				shadowNode.edgeMicroTypeNode = &microtypes.MicroTypeAst{
					SourceFile: viper.GetString("muSpec.dir") + "/" + shadowNode.edgeTypeNode.TypeSpec.XProto.Package + "/" + shadowNode.edgeTypeNode.TypeSpec.Type + ".types.yaml",
				}
			}

			if muTypesPerFile[shadowNode.edgeMicroTypeNode.SourceFile] == nil {
				muTypesPerFile[shadowNode.edgeMicroTypeNode.SourceFile] = []*microtypes.MicroType{}
			}

			typeLine := []string{}
			typeLine = append(typeLine, shadowNode.edgeTypeNode.TypeSpec.XProto.Package+"."+shadowNode.edgeTypeNode.TypeSpec.Type)

			typeLine = append(typeLine, "#"+shadowNode.edgeTypeNode.TypeSpec.Description)
			// field lines
			fields := orderedmap.New()

			shadowNode.edgeTypeNode.TypeSpec.Fields.Map(func(iKey interface{}, iValue interface{}) {
				f := iValue.(*specSpec.Field) //*string:1 # A * before the type means required
				fieldline := []string{}

				if f.Meta != nil && f.Meta.Readonly {
					fieldline = append(fieldline, "-")
				}

				if f.Constraints["required"] != nil {
					fieldline = append(fieldline, "*")
				}

				if f.Meta != nil && f.Meta.Repeated {
					fieldline = append(fieldline, "[]")
				}
				fieldline = append(fieldline, f.Type+":"+strconv.Itoa(int(f.XProto.Number)))
				fieldline = append(fieldline, "#"+f.Description)
				fields.Set(iKey, strings.Join(fieldline, " "))
			})

			muType := &microtypes.MicroType{
				Type:   strings.Join(typeLine, " "), //type: "sample.Sample  #Sample"
				Fields: fields,
				Target: shadowNode.edgeMicroTypeNode.Target,
			}

			// add type to "file"
			muTypesPerFile[shadowNode.edgeMicroTypeNode.SourceFile] = append(muTypesPerFile[shadowNode.edgeMicroTypeNode.SourceFile], muType)
		} else {
			// type is not in spec
			// todo: implement delete
		}
	}

	// store every item in muTypesPerFile (key is filename)
	for filename, muType := range muTypesPerFile {
		// save the stuff
		file, _ := yaml.Marshal(muType)
		if !util.DirExists(path.Dir(filename)) {
			util.MkdirRelative(path.Dir(filename))
		}
		err := ioutil.WriteFile("./"+filename, file, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
	// transfer every service assocList.ServiceItemsByName

}
