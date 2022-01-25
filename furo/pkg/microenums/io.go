package microenums

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"regexp"
	"strings"
)

// unmarshal yaml/json to microEnum list
func (l *MicroEnumlist) Unmarshal(data []byte) {
	parseError := yaml.Unmarshal(data, &l.MicroEnums) //reads yaml and json because json is just a subset of yaml
	if parseError != nil {
		log.Fatal(parseError)
	}
	l.MicroEnumsByName = map[string]*MicroEnum{}
	l.MicroEnumsASTByName = map[string]*MicroEnumAst{}
	// build the map
	for _, t := range l.MicroEnums {
		regex := regexp.MustCompile(`(?s)^([^#(]*):?([^#]*)?(#(.*))?$`)
		matches := regex.FindStringSubmatch(t.Enum)
		if len(matches) == 0 {
			fmt.Println("typeline not parseable", t.Enum)
		}

		typeName := strings.TrimSpace(matches[1])
		l.MicroEnumsByName[typeName] = t
		// convert the field values in list from yaml node to string
		t.Values.Map(func(iKey interface{}, iValue interface{}) {
			var valuestr string
			fieldYamlNode := iValue.(*yaml.Node)
			fieldYamlNode.Decode(&valuestr)
			t.Values.Set(iKey.(string), valuestr)
		})
		l.MicroEnumsASTByName[typeName] = t.ToMicroEnumAst()

	}
}
