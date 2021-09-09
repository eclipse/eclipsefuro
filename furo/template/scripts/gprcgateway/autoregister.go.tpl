package gateway
{{$config := .config}}{{$importlist := dict }}
import (
    "context"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	{{- range $servicename, $service:= .services}}
	{{- $pkgopts := split ";" $service.__proto.options.go_package }}
	//{{$servicename}}
	{{- if hasKey $importlist $pkgopts._1 | not}}
	{{$pkgopts._1}} "{{$pkgopts._0}}"
    {{end}}
    {{- $_ := set $importlist $pkgopts._1 "imported"}}{{- end}}
{{- range $servicename, $service:= .installedServices}}
//Installed service {{$servicename}}
	{{- $pkgopts := split ";" $service.__proto.options.go_package }}
	{{- if hasKey $importlist $pkgopts._1 | not}}
	{{$pkgopts._1}} "{{$pkgopts._0}}"
    {{end}}
    {{- $_ := set $importlist $pkgopts._1 "imported"}}
{{- end}}
	"google.golang.org/grpc"
	"net/http"
    {{- range $i , $import:= .config.extensions.gen_transcoder.additional_imports}}
    _ "{{$import}}"
    {{end}}
)


// newGateway returns a new gateway server which translates HTTP into gRPC.
func NewGateway(ctx context.Context, conn *grpc.ClientConn, opts []gwruntime.ServeMuxOption) (http.Handler, error) {

	mux := gwruntime.NewServeMux(opts...)

	for _, f := range []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{

{{- range $servicename, $service:= .services}}
{{- $pkgopts := split ";" $service.__proto.options.go_package }}
	// {{$servicename}}
	{{$pkgopts._1}}.Register{{$service.name}}Handler,
{{end}}

//installed services
{{- range $servicename, $service:= .installedServices}}
{{- $pkgopts := split ";" $service.__proto.options.go_package }}
	// {{$servicename}}
	{{$pkgopts._1}}.Register{{$service.name}}Handler,
{{end}}
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}
