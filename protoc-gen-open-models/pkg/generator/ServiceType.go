package generator

import (
	"bytes"
	"errors"
	"github.com/eclipse/eclipsefuro/protoc-gen-open-models/pkg/sourceinfo"
	"google.golang.org/genproto/googleapis/api/annotations"
	"path/filepath"
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

	pathSegments := strings.Split(service.Package, ".")
	for i, _ := range pathSegments {
		pathSegments[i] = ".."
	}

	imports.AddImport(strings.Join(pathSegments, "/")+"/API_OPTIONS", "API_OPTIONS")

	serviceType := ServiceType{
		Name:            fullQualifiedName(service.Package, service.Name),
		Methods:         make([]ServiceMethods, 0, len(service.Methods)),
		LeadingComments: multilineComment(service.Info.GetLeadingComments()),
		Package:         service.Package,
	}

	for _, method := range service.Methods {

		requestTypeFQ := fullQualifiedTypeName(method.Method.GetInputType())
		responseTypeFQ := fullQualifiedTypeName(method.Method.GetOutputType())

		fieldPackage := strings.Split("."+service.Package, ".")

		relIn, _ := filepath.Rel(strings.Join(fieldPackage, "/"), "/"+typenameToPath(method.Method.GetInputType()))
		relOut, _ := filepath.Rel(strings.Join(fieldPackage, "/"), "/"+typenameToPath(method.Method.GetOutputType()))
		if !strings.HasPrefix(relIn, "..") {
			relIn = "./" + relIn
		}
		if !strings.HasPrefix(relOut, "..") {
			relOut = "./" + relOut
		}
		imports.AddImport(relIn, "I"+requestTypeFQ)
		imports.AddImport(relOut, "I"+responseTypeFQ)

		verb, path, err := extractPathAndPattern(method.HttpRule.ApiOptions)
		// on err, we have no REST endpoints
		if err == nil {
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

func extractPathAndPattern(rule *annotations.HttpRule) (path string, pattern string, err error) {
	if rule == nil {
		return "", "", errors.New("No REST endpoint available")
	}
	switch r := rule.Pattern.(type) {
	case *annotations.HttpRule_Get:
		return "GET", r.Get, nil

	case *annotations.HttpRule_Put:
		return "PUT", r.Put, nil

	case *annotations.HttpRule_Post:
		return "POST", r.Post, nil

	case *annotations.HttpRule_Patch:
		return "PATCH", r.Patch, nil

	case *annotations.HttpRule_Delete:
		return "DELETE", r.Delete, nil

	case *annotations.HttpRule_Custom:
		return r.Custom.Kind, r.Custom.Path, nil

	}
	// should not happen
	return "", "", errors.New("No match")
}
