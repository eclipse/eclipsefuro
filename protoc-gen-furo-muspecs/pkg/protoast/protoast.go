package protoast

import (
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"strings"
)

type ProtoAST struct {
	Request          *pluginpb.CodeGeneratorRequest
	Response         *pluginpb.CodeGeneratorResponse
	Parameters       map[string]string
	ProtoMap         map[string]*descriptorpb.FileDescriptorProto
	FileProtoMap     map[string]*descriptorpb.FileDescriptorProto
	ProtoMapLocation map[string]*descriptorpb.SourceCodeInfo_Location
}

// Get a prameter from the used command
func (a *ProtoAST) GetParameter(parameter string) (value string, found bool) {
	for k, v := range a.Parameters {
		if k == parameter {
			return v, true
		}
	}
	return "", false
}

// build a map of all protos
func (a *ProtoAST) BuildMap() {
	for _, f := range a.Request.ProtoFile {
		a.ProtoMap[*f.Name] = f
	}
}

func NewProtoAST(request *pluginpb.CodeGeneratorRequest) *ProtoAST {
	parameters := request.GetParameter()
	Params := map[string]string{}
	groupkv := strings.Split(parameters, ",")
	for _, element := range groupkv {
		kv := strings.Split(element, "=")
		if len(kv) > 1 {
			Params[kv[0]] = kv[1]
		}
	}
	ast := &ProtoAST{
		Request: request,
		Response: &pluginpb.CodeGeneratorResponse{
			Error: nil,
			File:  make([]*pluginpb.CodeGeneratorResponse_File, 0),
		},
		Parameters:   Params,
		ProtoMap:     map[string]*descriptorpb.FileDescriptorProto{},
		FileProtoMap: map[string]*descriptorpb.FileDescriptorProto{},
	}

	ast.BuildMap()
	for _, n := range ast.Request.FileToGenerate {
		ast.FileProtoMap[n] = ast.ProtoMap[n]
	}

	return ast
}
