## Variables

- Name: {{.Var.Name}}
- TypeName: {{.Var.TypeName}}
- ServiceName: {{.Var.ServiceName}}

## Deep access

- .Var.Type.__proto.package: {{.Var.Type.__proto.package}}
- .Var.Service.__proto.package:

## Range
{{range $key, $val := .Var.Service.services}}
**{{$key}}**: {{$val.description}}
{{end}}

## Payload

- .Conf.Data.pi: {{.Conf.Data.pi}}