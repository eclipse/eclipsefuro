package sourceinfo

import (
	"github.com/google/gnostic/openapiv3"
	options "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/types/descriptorpb"
	"path/filepath"
)

type SourceInfo struct {
	Messages            []MessageInfo
	InlineMessages      []MessageInfo
	Enums               []EnumInfo
	InlineEnums         []EnumInfo
	Services            []ServiceInfo
	FileDescriptorProto *descriptorpb.FileDescriptorProto
	Package             string
	Path                string //  path without the file
}

type ServiceInfo struct {
	Name    string
	Info    *descriptorpb.SourceCodeInfo_Location
	Service *descriptorpb.ServiceDescriptorProto
	Methods []MethodInfo
	Package string
}

type MethodInfo struct {
	Name     string
	Info     *descriptorpb.SourceCodeInfo_Location
	Method   *descriptorpb.MethodDescriptorProto
	Service  *descriptorpb.ServiceDescriptorProto // link the parent service
	HttpRule *ApiOptionInfo
}

type ApiOptionInfo struct {
	Info       *descriptorpb.SourceCodeInfo_Location
	ApiOptions *options.HttpRule
}

type MessageInfo struct {
	Name           string
	Info           *descriptorpb.SourceCodeInfo_Location
	FieldInfos     []FieldInfo
	Message        descriptorpb.DescriptorProto
	Package        string
	OpenApiSchema  *openapi_v3.Schema
	NestedMessages []MessageInfo
	ParentOfNested *MessageInfo
}

type EnumInfo struct {
	Name       string
	Info       *descriptorpb.SourceCodeInfo_Location
	ValuesInfo []ValueInfo
	AllowAlias bool
	Message    descriptorpb.DescriptorProto
	Package    string
}

type ValueInfo struct {
	Name  string
	Value int32
	Info  *descriptorpb.SourceCodeInfo_Location
}
type FieldInfo struct {
	Name              string
	Info              *descriptorpb.SourceCodeInfo_Location
	Field             *descriptorpb.FieldDescriptorProto
	Message           *descriptorpb.DescriptorProto
	Package           string
	OpenApiProperties *openapi_v3.Schema
}

func GetSourceInfo(descr *descriptorpb.FileDescriptorProto) SourceInfo {
	sourceInfo := SourceInfo{FileDescriptorProto: descr, Package: descr.GetPackage(), Path: filepath.Dir(descr.GetName())}

	for _, location := range descr.GetSourceCodeInfo().GetLocation() {

		// 6 111 => 6 ServiceIndex
		// location info for descriptor.ServiceType[111]Field[222]
		if len(location.GetPath()) == 2 && location.Path[0] == 6 {
			msgIndex := location.Path[1]
			sourceInfo.Services = append(sourceInfo.Services, ServiceInfo{
				Name:    *descr.Service[msgIndex].Name,
				Service: descr.Service[msgIndex],
				Info:    location,
				Methods: []MethodInfo{},
				Package: descr.GetPackage(),
			})
		}

		// Methods
		// 6 111 2 222 4 9999 =>	 4 ServiceIndex 2 Method 4 Option
		// for field with index 222 in Service with index 111
		// location info for descriptor.ServiceType[111]Field[222]
		if len(location.GetPath()) == 4 && location.Path[0] == 6 && location.Path[2] == 2 {
			sIndex := location.Path[1]
			methodIndex := location.Path[3]
			apiOptions, _ := ExtractAPIOptions(descr.Service[sIndex].Method[methodIndex])
			fi := MethodInfo{
				Name:    *descr.Service[sIndex].Method[methodIndex].Name,
				Info:    location,
				Method:  descr.Service[sIndex].Method[methodIndex],
				Service: descr.Service[sIndex],
				HttpRule: &ApiOptionInfo{
					Info:       nil,
					ApiOptions: apiOptions,
				},
			}
			sourceInfo.Services[sIndex].Methods = append(sourceInfo.Services[sIndex].Methods, fi)
		}

		// Method options
		if len(location.GetPath()) == 6 && location.Path[0] == 6 && location.Path[2] == 2 && location.Path[4] == 4 {
			sIndex := location.Path[1]
			methodIndex := location.Path[3]
			sourceInfo.Services[sIndex].Methods[methodIndex].HttpRule.Info = location
		}

		// 5 111 => 5 EnumIndex
		// location info for descriptor.ServiceType[111]Field[222]
		if len(location.GetPath()) == 2 && location.Path[0] == 5 {
			msgIndex := location.Path[1]
			alias := false
			if descr.EnumType[msgIndex].Options != nil && descr.EnumType[msgIndex].Options.AllowAlias != nil {
				alias = *descr.EnumType[msgIndex].Options.AllowAlias
			}
			// values
			vals := []ValueInfo{}
			for _, v := range descr.EnumType[msgIndex].Value {
				fi := ValueInfo{
					Name:  *v.Name,
					Value: *v.Number,
				}
				vals = append(vals, fi)
			}

			sourceInfo.Enums = append(sourceInfo.Enums, EnumInfo{
				Name:       *descr.EnumType[msgIndex].Name,
				ValuesInfo: vals,
				AllowAlias: alias,
				Info:       location,
				Package:    descr.GetPackage(),
			})
		}

		// enum values
		if len(location.GetPath()) == 4 && location.Path[0] == 5 {

			msgIndex := location.Path[1]
			enumIndex := location.Path[3]
			sourceInfo.Enums[msgIndex].ValuesInfo[enumIndex].Info = location

		}

		// 4 111 2 222 => 4 MessageIndex 2 FieldIndex
		// for field with index 222 in message with index 111
		// location info for descriptor.MessageType[111]Field[222]
		if len(location.GetPath()) == 2 && location.Path[0] == 4 {
			msgIndex := location.Path[1]
			OpenApiProps, _ := ExtractOpenApiMessageOptions(descr.MessageType[msgIndex])
			sourceInfo.Messages = append(sourceInfo.Messages, MessageInfo{
				Name:          *descr.MessageType[msgIndex].Name,
				Message:       *descr.MessageType[msgIndex],
				Info:          location,
				FieldInfos:    []FieldInfo{},
				Package:       descr.GetPackage(),
				OpenApiSchema: OpenApiProps,
			})

			// check for inline enums
			for _, nestedEnum := range descr.MessageType[msgIndex].EnumType {

				alias := false
				if nestedEnum.Options != nil && nestedEnum.Options.AllowAlias != nil {
					alias = *nestedEnum.Options.AllowAlias
				}
				// values
				vals := []ValueInfo{}
				for _, v := range nestedEnum.Value {
					fi := ValueInfo{
						Name:  *v.Name,
						Value: *v.Number,
					}
					vals = append(vals, fi)
				}

				sourceInfo.InlineEnums = append(sourceInfo.InlineEnums, EnumInfo{
					Name:       *descr.MessageType[msgIndex].Name + "." + *nestedEnum.Name,
					ValuesInfo: vals,
					AllowAlias: alias,
					Info:       location,
				})

			}

			// check for nested messages
			for _, nestedMessage := range descr.MessageType[msgIndex].NestedType {

				mi := MessageInfo{
					Name:           *descr.MessageType[msgIndex].Name + "." + *nestedMessage.Name,
					Info:           nil,
					FieldInfos:     []FieldInfo{},
					Message:        descriptorpb.DescriptorProto{},
					Package:        descr.GetPackage(),
					OpenApiSchema:  nil,
					ParentOfNested: &sourceInfo.Messages[msgIndex],
				}

				for _, field := range nestedMessage.Field {
					fi := FieldInfo{
						Name:    field.GetName(),
						Info:    location,
						Field:   field,
						Message: nestedMessage,
						Package: descr.GetPackage(),
					}
					mi.FieldInfos = append(mi.FieldInfos, fi)
				}

				sourceInfo.Messages[msgIndex].NestedMessages = append(sourceInfo.Messages[msgIndex].NestedMessages, mi)
				sourceInfo.InlineMessages = append(sourceInfo.InlineMessages, mi)

			}
		}

		// 4 111 3 0 2 123 =>	 4 MessageIndex 3 NestedMessageIndex 2 NestedFieldIndex
		// for field with index 222 in message with index 111
		// location info for descriptor.MessageType[111]Field[222]
		if len(location.GetPath()) == 6 && location.Path[0] == 4 && location.Path[2] == 3 {
			msgIndex := location.Path[1]
			NestedMessageIndex := location.Path[3]
			NestedFieldIndex := location.Path[5]

			sourceInfo.Messages[msgIndex].NestedMessages[NestedMessageIndex].FieldInfos[NestedFieldIndex].Info = location

		}

		// 4 111 2 222 =>	 4 MessageIndex 2 FieldIndex
		// for field with index 222 in message with index 111
		// location info for descriptor.MessageType[111]Field[222]
		if len(location.GetPath()) == 4 && location.Path[0] == 4 && location.Path[2] == 2 {
			msgIndex := location.Path[1]
			fieldIndex := location.Path[3]
			OpenApiProps, _ := ExtractOpenApiFieldOptions(descr.MessageType[msgIndex].Field[fieldIndex])

			fi := FieldInfo{
				Name:              *descr.MessageType[msgIndex].Field[fieldIndex].Name,
				Info:              location,
				Field:             descr.MessageType[msgIndex].Field[fieldIndex],
				Message:           descr.MessageType[msgIndex],
				Package:           descr.GetPackage(),
				OpenApiProperties: OpenApiProps,
			}

			sourceInfo.Messages[msgIndex].FieldInfos = append(sourceInfo.Messages[msgIndex].FieldInfos, fi)
		}

	}
	return sourceInfo
}
