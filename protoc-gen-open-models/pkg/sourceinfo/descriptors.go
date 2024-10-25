package sourceinfo

import (
	"fmt"
	openapi_v3 "github.com/google/gnostic-models/openapiv3"
	options "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Extract the google.api.http option
// more details here vendor/google.golang.org/genproto/googleapis/api/annotations/http.pb.go
func ExtractAPIOptions(meth *descriptorpb.MethodDescriptorProto) (*options.HttpRule, error) {
	if meth.Options == nil {
		return nil, nil
	}
	if !proto.HasExtension(meth.Options, options.E_Http) {
		return nil, nil
	}
	ext := proto.GetExtension(meth.Options, options.E_Http)
	opts, ok := ext.(*options.HttpRule)
	if !ok {
		return nil, fmt.Errorf("extension is %T; want an HttpRule", ext)
	}
	return opts, nil
}

func ExtractOpenApiFieldOptions(field *descriptorpb.FieldDescriptorProto) (*openapi_v3.Schema, error) {
	if field.Options == nil {
		return nil, nil
	}
	if !proto.HasExtension(field.Options, openapi_v3.E_Property) {
		return nil, nil
	}
	ext := proto.GetExtension(field.Options, openapi_v3.E_Property)
	opts, ok := ext.(*openapi_v3.Schema)
	if !ok {
		return nil, fmt.Errorf("extension is %T; want an OpenApi Property", ext)
	}
	return opts, nil
}

func ExtractOpenApiMessageOptions(message *descriptorpb.DescriptorProto) (*openapi_v3.Schema, error) {
	if message.Options == nil {
		return nil, nil
	}
	if !proto.HasExtension(message.Options, openapi_v3.E_Schema) {
		return nil, nil
	}
	ext := proto.GetExtension(message.Options, openapi_v3.E_Schema)
	opts, ok := ext.(*openapi_v3.Schema)
	if !ok {
		return nil, fmt.Errorf("extension is %T; want an OpenApi Property", ext)
	}
	return opts, nil
}
