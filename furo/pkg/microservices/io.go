package microservices

import (
	"gopkg.in/yaml.v3"
	"log"
)

// unmarshal yaml/json to microtype list
func (l *MicroServiceList) Unmarshal(data []byte) {
	parseError := yaml.Unmarshal(data, &l.MicroServices) //reads yaml and json because json is just a subset of yaml
	if parseError != nil {
		log.Fatal(parseError)
	}
	l.MicroServicesByName = map[string]*MicroService{}
	l.MicroServicesASTByName = map[string]*MicroServiceAst{}

}
