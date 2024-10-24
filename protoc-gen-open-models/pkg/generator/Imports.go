package generator

import (
	"bytes"
	"text/template"
)

type ImportMap struct {
	// map[__modulepath__]map[__import__]bool
	Imports map[string]map[string]bool
}

func (im ImportMap) AddImport(typescriptModulePath string, cls string) {
	if im.Imports[typescriptModulePath] == nil {
		im.Imports[typescriptModulePath] = map[string]bool{}
	}
	im.Imports[typescriptModulePath][cls] = true
}

func (im ImportMap) Render() string {
	t, err := template.New("Imports").Parse(ImportTemplate)
	if err != nil {
		panic(err)
	}

	var res bytes.Buffer
	err = t.Execute(&res, im)
	if err != nil {
		panic(err)
	}
	return res.String()
}

var ImportTemplate = `
{{- range $import, $clsMap := .Imports}}
{{ $first := true }}import { {{range $cls, $_ := $clsMap}}{{if not $first}}, {{else}}{{$first = false}}{{end}}{{$cls}}{{end}} } from "{{$import}}";
{{end}}
`
