package muSrvSanitize

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/microservices"
	"github.com/eclipse/eclipsefuro/furo/pkg/orderedmap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"regexp"
)

type ServicesPerFile struct {
	MicroServices []*microservices.MicroService
	hasChanged    bool
}

func checkQueryParams(microServicesByName map[string]*microservices.MicroService) {
	fileChanges := map[string]*ServicesPerFile{}

	for _, ms := range microServicesByName {
		if fileChanges[ms.SourceFile] == nil {
			fileChanges[ms.SourceFile] = &ServicesPerFile{
				MicroServices: nil,
				hasChanged:    false,
			}
		}
		fileChanges[ms.SourceFile].MicroServices = append(fileChanges[ms.SourceFile].MicroServices, ms)
		// sourcefile no longer needed, and we do not want it in yaml
		thisSrc := ms.SourceFile
		ms.SourceFile = ""

		for sIndex, rpc := range ms.Methods {
			regex := regexp.MustCompile(`^([^:]+):\s?([A-Z]*)\s?([^\s]*)\s?([^#]*)\s?,\s?([^#]*)\s?#?(?s:(.*))$`)
			matches := regex.FindStringSubmatch(rpc.Md)
			if len(matches) == 0 {
				fmt.Println("typeline not parseable", rpc.Md)
			}
			// match 3 contains the url
			url := matches[3]

			rgxqp := regexp.MustCompile(`({([^{]*)})`)
			qpmatches := rgxqp.FindAllStringSubmatch(url, -1)

			if rpc.Qp == nil {
				rpc.Qp = orderedmap.New()
			}
			for _, qpshort := range qpmatches {
				_, has := rpc.Qp.Get(qpshort[2])
				if !has {
					rpc.Qp.Set(qpshort[2], "string #The query param "+qpshort[2]+" stands for XXX id.")
					fileChanges[thisSrc].hasChanged = true
				}
			}

			if rpc.Qp.Len() == 0 {
				rpc.Qp = nil
			}

			ms.Methods[sIndex] = rpc

		}

	}

	// save if contentChanged
	for filename, content := range fileChanges {
		if content.hasChanged {
			file, _ := yaml.Marshal(content.MicroServices)
			err := ioutil.WriteFile("./"+filename, file, 0644)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
