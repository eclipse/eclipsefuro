package clientspec

import (
	"github.com/eclipse/eclipsefuro/furo/pkg/orderedmap"
	"github.com/eclipse/eclipsefuro/furo/pkg/specSpec"
)

func CreateClientTypeFromAstType(ast *specSpec.Type) (t *Type) {

	t = &Type{
		Name:   ast.Name,
		Type:   ast.Type,
		Fields: orderedmap.New(),
	}

	ast.Fields.Map(func(iKey interface{}, iValue interface{}) {
		astField := iValue.(*specSpec.Field)
		field := &Field{
			Type:        astField.Type,
			Meta:        astField.Meta,
			Constraints: astField.Constraints,
			XProto:      astField.XProto,
		}
		t.Fields.Set(iKey, field)

	})

	return t
}

func CreateClientTypeFromEnum(ast *specSpec.Enum) (t *Type) {

	t = &Type{
		Type:   ast.Type,
		Values: orderedmap.New(),
	}

	ast.Values.Map(func(iKey interface{}, iValue interface{}) {
		t.Values.Set(iKey, iValue)
	})

	return t
}

func CreateServiceFromAstService(ast *specSpec.Service, fullname string) (t *Service) {

	t = &Service{
		Name:     ast.Name,
		Services: orderedmap.New(),
	}

	ast.Services.Map(func(iKey interface{}, iValue interface{}) {
		astField := iValue.(*specSpec.Rpc)
		s := &CompressedService{
			Data: &specSpec.Servicereqres{
				Request:   resolveFullQualifiedTypename(astField.Data.Request, fullname),
				Response:  resolveFullQualifiedTypename(astField.Data.Response, fullname),
				Bodyfield: astField.Data.Bodyfield,
			},
			Deeplink: &Deeplink{
				Href:   astField.Deeplink.Href,
				Method: astField.Deeplink.Method,
				Rel:    astField.Deeplink.Rel,
			},
			Query: astField.Query, // todo check if this is still needed
		}
		t.Services.Set(iKey, s)
	})

	return t
}

type Service struct {
	// Name of the type
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" yaml:"name"`

	// services => which is a CompressedService
	Services *orderedmap.OrderedMap `json:"services" yaml:"services" `
}

type CompressedService struct {
	Data     *specSpec.Servicereqres `json:"data,omitempty" yaml:"data,omitempty"`
	Deeplink *Deeplink               `json:"deeplink,omitempty" yaml:"deeplink,omitempty"`
	Query    *orderedmap.OrderedMap  `json:"query,omitempty" yaml:"query,omitempty"`
}
type Deeplink struct {
	// The link pattern, like /api/xxx/{qp}/yyy
	Href string `json:"href,omitempty" yaml:"href,omitempty"`
	// method of curl
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
	// the relationship
	Rel string `json:"rel,omitempty" yaml:"rel,omitempty"`
}

// Defines a type in the furo env spec
type Type struct {
	// Name of the type
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	// the type
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" yaml:"type"`
	// fields of a type
	Fields *orderedmap.OrderedMap `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// values for enum types
	Values *orderedmap.OrderedMap `protobuf:"bytes,5,rep,name=fields,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
type Field struct {
	// the field type, https://developers.google.com/protocol-buffers/docs/proto3#scalar
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" yaml:"type"`
	// meta information for the client, like label, default, repeated, options...
	Meta *specSpec.FieldMeta `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta" yaml:"meta"`
	// constraints for a field, like min{}, max{}, step{}
	Constraints map[string]*specSpec.FieldConstraint `protobuf:"bytes,4,rep,name=constraints,proto3" json:"constraints" json:"yaml" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// information for the proto generator, like number, type
	XProto *specSpec.Fieldproto `json:"__proto" yaml:"__proto"`
}
