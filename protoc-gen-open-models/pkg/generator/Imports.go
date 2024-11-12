package generator

import (
	"bytes"
	"text/template"
)

type ImportMap struct {
	// map[__modulepath__]map[__import__]bool
	Imports map[string]map[string]string
	//            ^           ^     ^
	//            |           |     |
	//        modulepath   class   fullqualified class
}

func (im ImportMap) AddImport(typescriptModulePath string, className string, fullQualifiedClassName string) {
	if im.Imports[typescriptModulePath] == nil {
		im.Imports[typescriptModulePath] = map[string]string{}
	}
	im.Imports[typescriptModulePath][className] = fullQualifiedClassName
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
{{ $first := true }}import { {{range $cls, $as := $clsMap}}{{if not $first}}, {{else}}{{$first = false}}{{end}}{{$cls}}{{if $as}} as {{$as}}{{end}}{{end}} } from "{{$import}}";
{{end}}
`
