package generator

import (
	"bytes"
	"protoc-gen-open-models/pkg/sourceinfo"
	"text/template"
)

type TransportType struct {
	Name   string
	Fields []TransportFields
}

type TransportFields struct {
	LeadingComments []string
	TrailingComment string
	FieldName       string // name of the field in interface types and models
	FieldProtoName  string // name of the field in transport types
	Type            string
}

func (r *TransportType) Render() string {

	t, err := template.New("TransportType").Parse(TransportTypeTemplate)
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

var TransportTypeTemplate = `export interface T{{.Name}} {
{{- range .Fields}}{{if .LeadingComments}}{{range $i, $commentLine := .LeadingComments}}
  // {{$commentLine}}{{end}}{{end}}
  {{.FieldProtoName}}?: {{.Type}};{{if .TrailingComment}} // {{.TrailingComment}}{{end}}{{end}}
}
`

func prepareTransportType(message *sourceinfo.MessageInfo, imports ImportMap) TransportType {
	transportType := TransportType{
		Name:   fullQualifiedName(message.Package, message.Name),
		Fields: nil,
	}
	for _, field := range message.FieldInfos {
		transportType.Fields = append(transportType.Fields, TransportFields{
			LeadingComments: multilineComment(field.Info.GetLeadingComments()),
			TrailingComment: field.Info.GetTrailingComments(),
			FieldName:       field.Field.GetJsonName(), // todo: check preserve proto names
			FieldProtoName:  field.Field.GetName(),     // todo: check  preserve proto names
			Type:            resolveInterfaceType(imports, field, "T"),
		})
	}
	return transportType
}
