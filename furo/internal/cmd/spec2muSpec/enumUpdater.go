package spec2muSpec

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/microenums"
	"github.com/eclipse/eclipsefuro/furo/pkg/orderedmap"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"strings"
)

func updateAndStoreMicroEnums(enumItems map[string]*UTshadowNode) {
	muEnumsPerFile := map[string][]*microenums.MicroEnum{}

	// transfer every enum assocList.EnumItemsByName
	for _, shadowNode := range enumItems {
		if shadowNode.edgeEnumNode != nil {

			if shadowNode.edgeMicroEnumNode == nil {
				// create muEnum because it does not exist
				shadowNode.edgeMicroEnumNode = &microenums.MicroEnumAst{
					SourceFile: viper.GetString("muSpec.dir") + "/" + shadowNode.edgeEnumNode.EnumSpec.XProto.Package + "/" + shadowNode.edgeEnumNode.EnumSpec.Type + ".enums.yaml",
				}
			}

			if muEnumsPerFile[shadowNode.edgeMicroEnumNode.SourceFile] == nil {
				muEnumsPerFile[shadowNode.edgeMicroEnumNode.SourceFile] = []*microenums.MicroEnum{}
			}

			enumLine := []string{}
			enumLine = append(enumLine, shadowNode.edgeEnumNode.EnumSpec.XProto.Package+"."+shadowNode.edgeEnumNode.EnumSpec.Type)

			enumLine = append(enumLine, "#"+shadowNode.edgeEnumNode.EnumSpec.Description)
			// field lines
			values := orderedmap.New()

			shadowNode.edgeEnumNode.EnumSpec.Values.Map(func(iKey interface{}, iValue interface{}) {

				values.Set(iKey, iValue)
			})

			muEnum := &microenums.MicroEnum{
				Enum:       strings.Join(enumLine, " "), //enum: "sample.Sample  #Sample"
				Values:     values,
				Target:     shadowNode.edgeMicroEnumNode.Target,
				AllowAlias: shadowNode.edgeEnumNode.EnumSpec.XProto.AllowAlias,
			}

			// add enum to "file"
			muEnumsPerFile[shadowNode.edgeMicroEnumNode.SourceFile] = append(muEnumsPerFile[shadowNode.edgeMicroEnumNode.SourceFile], muEnum)
		} else {
			// enum is not in spec
			// todo: implement delete
		}
	}

	// store every item in muEnumsPerFile (key is filename)
	for filename, muEnum := range muEnumsPerFile {
		// save the stuff
		file, _ := yaml.Marshal(muEnum)
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
