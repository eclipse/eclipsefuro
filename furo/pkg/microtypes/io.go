package microtypes

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/orderedmap"
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
		makeEntity := false
		makeCollection := false
		if matches[2] != "" {
			if strings.Contains(matches[2], "c") {
				makeCollection = true
			}
			if strings.Contains(matches[2], "e") {
				makeEntity = true
			}
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

		if makeEntity {

			fields := orderedmap.New()
			fields.Set("data", typeName+":1 #the data contains a "+typeName)
			fields.Set("links", "[]furo.Link:2 #the Hateoas links")
			fields.Set("meta", "furo.Meta:3 #Meta for the response")
			entity := &MicroType{
				Type:   typeName + "Entity #Entitycontainer which holds a " + typeName,
				Fields: fields,
				Target: "",
			}
			l.MicroTypes = append(l.MicroTypes, entity)
		}
		if makeCollection {
			fields := orderedmap.New()
			fields.Set("entities", "[]"+typeName+"Entity:1 #the data contains a "+typeName)
			fields.Set("links", "[]furo.Link:2 #the Hateoas links")
			fields.Set("meta", "furo.Meta:3 #Meta for the response")
			collection := &MicroType{
				Type:   typeName + "Collection #Collectioncontainer which holds a " + typeName,
				Fields: fields,
				Target: "",
			}
			l.MicroTypes = append(l.MicroTypes, collection)
		}

	}
}
