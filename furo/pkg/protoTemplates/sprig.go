package protoTemplates

import (
	"github.com/Masterminds/sprig"
	"github.com/eclipse/eclipsefuro/furo/pkg/orderedmap"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
	"github.com/spf13/viper"
	"html/template"
	"regexp"
	"strings"
)

type FieldMap struct {
	Fieldname string
	Field     *specSpec.Field
}

func fieldpairs(orderedMap orderedmap.OrderedMap) (f []FieldMap) {
	f = []FieldMap{}

	for p := orderedMap.Oldest(); p != nil; p = p.Next() {
		m := FieldMap{
			Fieldname: p.Key.(string),
			Field:     p.Value.(*specSpec.Field),
		}
		f = append(f, m)
	}

	return f
}

func rpcmap(orderedMap *orderedmap.OrderedMap) (f map[string]*specSpec.Rpc) {
	f = map[string]*specSpec.Rpc{}

	for p := orderedMap.Oldest(); p != nil; p = p.Next() {
		f[p.Key.(string)] = p.Value.(*specSpec.Rpc)
	}

	return f
}

func noescape(str string) template.HTML {
	return template.HTML(str)
}

func collectoneof(fields []FieldMap) (oneoflist map[string]map[string]*specSpec.Field) {
	oneoflist = map[string]map[string]*specSpec.Field{}
	for _, f := range fields {
		field := f.Field
		if field.XProto.Oneof != "" {
			if oneoflist[field.XProto.Oneof] == nil {
				oneoflist[field.XProto.Oneof] = map[string]*specSpec.Field{}
			}
			oneoflist[field.XProto.Oneof][f.Fieldname] = field
		}
	}
	return oneoflist
}

type Qparams struct {
	Description string
	Type        string
	Name        string
	FieldId     int
}

// extract request params from service
func requestParamsFromRpc(rpc *specSpec.Rpc) []Qparams {
	// path params regexFindAll  "{[a-zA-Z]*}" $method.deeplink.href 10
	regex := regexp.MustCompile(`{([^{]+)}`)
	matches := regex.FindAllStringSubmatch(rpc.Deeplink.Href, 50)
	params := []Qparams{}
	i := 1
	for _, match := range matches {
		params = append(params, Qparams{
			Description: "deeplink params for path segment {" + match[1] + "} of " + rpc.Deeplink.Href,
			Type:        "string",
			Name:        match[1],
			FieldId:     i,
		})
		i++
	}

	// request type
	if rpc.Data.Request != "" {
		params = append(params, Qparams{
			Description: "request type " + rpc.Data.Request,
			Type:        rpc.Data.Request,
			Name:        "data",
			FieldId:     i,
		})
		i++
	}

	// query
	if rpc.Query != nil {
		rpc.Query.Map(func(iKey interface{}, iValue interface{}) {
			p := iValue.(*specSpec.Queryparam)
			params = append(params, Qparams{
				Description: p.Description,
				Type:        p.Type,
				Name:        iKey.(string),
				FieldId:     i,
			})
			i++
		})
	}

	return params
}

// builds the request for the rpc
// rpc {{$method.RpcName}} ({{$method | rpcRequest}}) returns ({{$method.Data.Response}}){
func rpcRequest(method *specSpec.Rpc) string {
	if strings.HasPrefix(method.Data.Request, "stream ") {
		return method.Data.Request
	}
	return method.RpcName + viper.GetString("muSpec.requestTypeSuffix")
}

func isNotStream(method *specSpec.Rpc) bool {
	if strings.HasPrefix(method.Data.Request, "stream ") {
		return false
	}
	return true
}

func GetSprigFuncs() template.FuncMap {
	fn := sprig.FuncMap()
	fn["noescape"] = noescape
	fn["collectoneof"] = collectoneof
	fn["fieldpairs"] = fieldpairs
	fn["rpcmap"] = rpcmap
	fn["requestParamsFromRpc"] = requestParamsFromRpc
	fn["rpcRequest"] = rpcRequest
	fn["isNotStream"] = isNotStream
	return fn
}
