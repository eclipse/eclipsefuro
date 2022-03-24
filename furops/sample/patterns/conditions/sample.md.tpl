## Variables


- UseService: {{.Var.UseService}}
- TypeName: {{.Var.TypeName}}
- ServiceName: {{.Var.ServiceName}}

{{if .Var.UseService}}
USE SERVICE was True

## Range
{{range $key, $val := .Var.Service.services}}
**{{$key}}**: {{$val.description}}
{{end}}

{{end}}


## Payload

- .Conf.Data.pi: {{.Conf.Data.pi}}