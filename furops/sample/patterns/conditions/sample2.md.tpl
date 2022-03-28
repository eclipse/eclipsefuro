## Variant 2


- UseService: {{.Var.UseService}}
- TypeName: {{.Var.TypeName}}
- ServiceName: {{.Var.ServiceName}}

## Range
{{range $key, $val := .Var.Service.services}}
**{{$key}}**: {{$val.description}}
{{end}}



## Payload

- .Conf.Data.pi: {{.Conf.Data.pi}}