package generator

import (
	"bytes"
	"google.golang.org/genproto/googleapis/api/annotations"
	"protoc-gen-open-models/pkg/sourceinfo"
	"strings"
	"text/template"
)

type ServiceType struct {
	Name            string
	Methods         []ServiceMethods
	LeadingComments []string
	Package         string
}

type ServiceMethods struct {
	Name                string // GetList
	RequestTypeLiteral  string //LFuroCubeCubeServiceGetRequest
	ResponseTypeLiteral string // LFuroCubeCubeServiceGetResponse
	Verb                string // GET
	Path                string // /v1/cubes/{cube_id}
	Body                string // "" | "*" | "data"
	LeadingComments     []string
	TrailingComment     string
}

var ServiceTemplate = `{{if .LeadingComments}}{{range $i, $commentLine := .LeadingComments}}
// {{$commentLine}}{{end}}{{end}}

export class {{.Name}} {

{{range $i, $method := .Methods}}{{if .LeadingComments}}{{range $i, $commentLine := .LeadingComments}}
  // {{$commentLine}}{{end}}{{end}}
  public {{.Name}}: Fetcher<{{.RequestTypeLiteral}},{{.ResponseTypeLiteral}}> = new Fetcher<{{.RequestTypeLiteral}},{{.ResponseTypeLiteral}}>(
    API_OPTIONS,
    '{{.Verb}}',
    '{{.Path}}',{{if .Body}}
    '{{.Body}}'{{end}}
  ); {{if .TrailingComment}}// {{.TrailingComment}}{{end}}
{{end}}

}
`

func (r *ServiceType) Render() string {

	t, err := template.New("ServiceType").Parse(ServiceTemplate)
	if err != nil {
		panic(err)
	}

	var res bytes.Buffer
	err = t.Execute(&res, r)
	if err != nil {
		panic(err)
	}
	return res.String()
}

func prepareServiceType(service sourceinfo.ServiceInfo, imports ImportMap) ServiceType {

	imports.AddImport("@furo/open-models/dist/Fetcher", "Fetcher")
	imports.AddImport("../API_OPTIONS", "API_OPTIONS")

	serviceType := ServiceType{
		Name:            fullQualifiedName(service.Package, service.Name),
		Methods:         make([]ServiceMethods, 0, len(service.Methods)),
		LeadingComments: multilineComment(service.Info.GetLeadingComments()),
		Package:         service.Package,
	}

	for _, method := range service.Methods {

		requestType := baseTypeName(method.Method.GetInputType())
		requestTypeFQ := fullQualifiedTypeName(method.Method.GetInputType())

		responseType := baseTypeName(method.Method.GetOutputType())
		responseTypeFQ := fullQualifiedTypeName(method.Method.GetOutputType())

		imports.AddImport("./"+requestType, "I"+requestTypeFQ)
		imports.AddImport("./"+responseType, "I"+responseTypeFQ)

		verb, path := extractPathAndPattern(method.HttpRule.ApiOptions)
		serviceMethods := ServiceMethods{
			Name:                method.Name,
			RequestTypeLiteral:  "I" + requestTypeFQ,
			ResponseTypeLiteral: "I" + responseTypeFQ,
			Verb:                verb,
			Path:                path,
			Body:                method.HttpRule.ApiOptions.GetBody(),
			LeadingComments:     multilineComment(method.Info.GetLeadingComments()),
			TrailingComment:     method.Info.GetTrailingComments(),
		}

		serviceType.Methods = append(serviceType.Methods, serviceMethods)
	}

	return serviceType
}

func baseTypeName(typeName string) string {
	ts := strings.Split(typeName, ".")
	return ts[len(ts)-1]
}

func fullQualifiedTypeName(typeName string) string {
	p := []string{}
	for _, s := range strings.Split(typeName[1:], ".") {
		p = append(p, strings.ToUpper(s[:1])+s[1:])
	}
	return strings.Join(p, "")
}

func extractPathAndPattern(rule *annotations.HttpRule) (path string, pattern string) {
	switch r := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return "GET", r.Get

	case *annotations.HttpRule_Put:
		return "PUT", r.Put

	case *annotations.HttpRule_Post:
		return "POST", r.Post

	case *annotations.HttpRule_Patch:
		return "PATCH", r.Patch

	case *annotations.HttpRule_Delete:
		return "DELETE", r.Delete

	case *annotations.HttpRule_Custom:
		return r.Custom.Kind, r.Custom.Path

	}
	// should not happen
	return "", ""
}
