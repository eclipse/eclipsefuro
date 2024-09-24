package generator

import (
	"bytes"
	"protoc-gen-open-models/pkg/sourceinfo"
	"strings"
	"text/template"
)

type LiteralType struct {
	Name   string
	Fields []LiteralFields
}

type LiteralFields struct {
	LeadingComments    []string
	TrailingComment    string
	FieldName          string // name of the field in interface types and models
	FieldTransportName string // name of the field in transport types
	Type               string
}

var LiteralTypeTemplate = `export interface L{{.Name}} {
{{- range .Fields}}{{if .LeadingComments}}{{range $i, $commentLine := .LeadingComments}}
  // {{$commentLine}}{{end}}{{end}}
  {{.FieldName}}?: {{.Type}};{{if .TrailingComment}} // {{.TrailingComment}}{{end}}{{end}}
}
`

func (r *LiteralType) Render() string {

	t, err := template.New("LiteralType").Parse(LiteralTypeTemplate)
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

func prepareLiteralType(message *sourceinfo.MessageInfo, imports ImportMap) LiteralType {
	literalType := LiteralType{
		Name:   fullQualifiedName(message.Package, message.Name),
		Fields: nil,
	}
	for _, field := range message.FieldInfos {
		literalType.Fields = append(literalType.Fields, LiteralFields{
			LeadingComments:    multilineComment(field.Info.GetLeadingComments()),
			TrailingComment:    field.Info.GetTrailingComments(),
			FieldName:          field.Field.GetJsonName(), // todo: check preserve proto names
			FieldTransportName: field.Field.GetName(),     // todo: check  preserve proto names
			Type:               resolveInterfaceType(imports, field, "L"),
		})
	}
	return literalType
}

func fullQualifiedName(pkg string, name string) string {
	p := []string{}
	for _, s := range strings.Split(pkg, ".") {
		p = append(p, strings.ToUpper(s[:1])+s[1:])
	}
	return strings.Join(p, "") + name
}
