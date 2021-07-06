package protoTemplates

var TypeTemplate = `// Code generated by furo-proto-gen. DO NOT EDIT.
syntax = "proto3";
package {{.Package}};
{{- range $key, $option := .Options}}
option {{$key}} = {{if  ($option | eq "true")}}{{$option}};{{else}}"{{$option}}";{{end}}{{end}}
{{range $import := .Imports}}
import "{{$import}}";{{end}}
{{range $message := .Types}}

{{if $message.Description}}// {{$message.Description | replace "\n" "\n// "}}{{end}}{{$fp := $message.Fields | fieldpairs}}
message {{$message.Type}} {  {{range  $fmap := $fp -}}{{ $fieldname := $fmap.Fieldname }}{{$field := $fmap.Field -}}{{if not $field.XProto.Oneof}}

    {{if $field.Description}}// {{$field.Description | replace "\n" "\n// " | noescape}}{{end}}
    {{if $field.Meta}}{{if $field.Meta.Repeated}}repeated {{end}}{{end}}{{$field.Type | noescape }} {{$fieldname}} = {{$field.XProto.Number}};{{end}}{{end}}{{$oneofFields := $fp | collectoneof}}{{range $oneof_name, $fields := $oneofFields}}
    oneof {{$oneof_name}} {
    {{- range $fieldname, $field := $fields}}

        {{if $field.Description}}// {{$field.Description | replace "\n" "\n// "  | noescape}}{{end}}
        {{if $field.Meta}}{{if $field.Meta.Repeated}}repeated {{end}}{{end}}{{$field.Type | noescape }} {{$fieldname}} = {{$field.XProto.Number}};
    {{- end}}
    }{{end}}
}{{end -}}
`