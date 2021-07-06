package protoTemplates

var BundledServiceTemplate = `// Code generated by furo spectools. DO NOT EDIT.

syntax = "proto3";
package {{.Package}};
{{- range $key, $option := .Options}}
option {{$key}} = {{if  ($option | eq "true")}}{{$option}};{{else}}"{{$option}}";{{end}}{{end}}
{{range $import := .AllImports}}
import "{{$import}}";{{end}}

{{range $k, $Services := .AllServices}}
{{$GenAdditionalBinding := $Services.GenAdditionalBinding}}
{{range $Service := $Services.Services}}
{{$serviceName := camelcase (snakecase $Service.Name)}}
{{if $Service.Description}}// {{$Service.Description | replace "\n" "\n// "}}{{end}}
service {{$serviceName}} { 
{{- $rpcmap := $Service.Services | rpcmap}}
{{range $rpckey, $method := $rpcmap}}
  {{if $method.Description}}// {{$method.Description | replace "\n" "\n// " | noescape}}{{end}}
  rpc {{$method.RpcName}} ({{if $method | isNotStream}}{{$Services.Package}}.{{end}}{{$method | rpcRequest}}) returns ({{$method.Data.Response}}){
	{{if $method.Deeplink.Description}}//{{$method.Deeplink.Description | replace "\n" "\n// "}}{{else}}// developer was to lazy to describe the rpc, sorry{{end}}
	option (google.api.http) = {
		{{ lower $method.Deeplink.Method}}: "{{$method.Deeplink.Href}}"{{ if $method.Data.Request}}
		{{ if or (eq $method.Deeplink.Method "POST") (eq $method.Deeplink.Method "PATCH") (eq $method.Deeplink.Method "PUT")}}body: "{{$method.Data.Bodyfield}}"{{end}}{{end}}
		{{- if and (eq $method.Deeplink.Method "PUT") (eq $GenAdditionalBinding true)  }}
		additional_bindings {
            patch: "{{$method.Deeplink.Href}}"
            body: "{{$method.Data.Bodyfield}}"
        }{{end}}
	};
  }
{{end}}
}
{{end}}
{{end}}

`