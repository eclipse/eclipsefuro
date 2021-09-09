package microtypes

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"regexp"
	"strings"
)

// unmarshal yaml/json to microtype list
func (l *MicroTypelist) Unmarshal(data []byte) {
	parseError := yaml.Unmarshal(data, &l.MicroTypes) //reads yaml and json because json is just a subset of yaml
	if parseError != nil {
		log.Fatal(parseError)
	}
	l.MicroTypesByName = map[string]*MicroType{}
	l.MicroTypesASTByName = map[string]*MicroTypeAst{}
	// build the map
	for _, t := range l.MicroTypes {
		regex := regexp.MustCompile(`(?s)^([^#(]*):?([^#]*)?(#(.*))?$`)
		matches := regex.FindStringSubmatch(t.Type)
		if len(matches) == 0 {
			fmt.Println("typeline not parseable", t.Type)
		}

		typeName := strings.TrimSpace(matches[1])
		l.MicroTypesByName[typeName] = t
		// convert the field values in list from yaml node to string
		t.Fields.Map(func(iKey interface{}, iValue interface{}) {
			var fieldstr string
			fieldYamlNode := iValue.(*yaml.Node)
			fieldYamlNode.Decode(&fieldstr)
			t.Fields.Set(iKey.(string), fieldstr)
		})
		l.MicroTypesASTByName[typeName] = t.ToMicroTypeAst()

	}
}
