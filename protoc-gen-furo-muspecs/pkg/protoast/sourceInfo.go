package protoast

import (
	options "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/types/descriptorpb"
)

type SourceInfo struct {
	Messages []MessageInfo
	Enums    []EnumInfo
	Services []ServiceInfo
}

type ServiceInfo struct {
	Name    string
	Info    *descriptorpb.SourceCodeInfo_Location
	Service *descriptorpb.ServiceDescriptorProto
	Methods []MethodInfo
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
	Name       string
	Info       *descriptorpb.SourceCodeInfo_Location
	FieldInfos []FieldInfo
	Message    descriptorpb.DescriptorProto
}

type EnumInfo struct {
	Name       string
	Info       *descriptorpb.SourceCodeInfo_Location
	ValuesInfo []ValueInfo
	Message    descriptorpb.DescriptorProto
}

type ValueInfo struct {
	Name  string
	Value int32
}
type FieldInfo struct {
	Name    string
	Info    *descriptorpb.SourceCodeInfo_Location
	Field   *descriptorpb.FieldDescriptorProto
	Message descriptorpb.DescriptorProto
}

func GetSourceInfo(descr *descriptorpb.FileDescriptorProto) SourceInfo {
	SourceInfo := SourceInfo{}

	for _, location := range descr.GetSourceCodeInfo().GetLocation() {

		// 6 111 => 6 ServiceIndex
		// location info for descriptor.ServiceType[111]Field[222]
		if len(location.GetPath()) == 2 && location.Path[0] == 6 {
			msgIndex := location.Path[1]
			SourceInfo.Services = append(SourceInfo.Services, ServiceInfo{
				Name:    *descr.Service[msgIndex].Name,
				Service: descr.Service[msgIndex],
				Info:    location,
				Methods: []MethodInfo{},
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
			SourceInfo.Services[sIndex].Methods = append(SourceInfo.Services[sIndex].Methods, fi)
		}

		// Method options
		if len(location.GetPath()) == 6 && location.Path[0] == 6 && location.Path[2] == 2 && location.Path[4] == 4 {
			sIndex := location.Path[1]
			methodIndex := location.Path[3]
			SourceInfo.Services[sIndex].Methods[methodIndex].HttpRule.Info = location
		}

		// 5 111 => 5 EnumIndex
		// location info for descriptor.ServiceType[111]Field[222]
		if len(location.GetPath()) == 2 && location.Path[0] == 5 {
			msgIndex := location.Path[1]
			SourceInfo.Enums = append(SourceInfo.Enums, EnumInfo{
				Name:       *descr.EnumType[msgIndex].Name,
				ValuesInfo: []ValueInfo{},
				Info:       location,
			})
		}

		if len(location.GetPath()) == 4 && location.Path[0] == 5 {
			msgIndex := location.Path[1]
			fieldIndex := location.Path[3]
			fi := ValueInfo{
				Name:  *descr.EnumType[msgIndex].Value[fieldIndex].Name,
				Value: *descr.EnumType[msgIndex].Value[fieldIndex].Number,
			}
			SourceInfo.Enums[msgIndex].ValuesInfo = append(SourceInfo.Enums[msgIndex].ValuesInfo, fi)
		}

		// 4 111 2 222 => 4 MessageIndex 2 FieldIndex
		// for field with index 222 in message with index 111
		// location info for descriptor.MessageType[111]Field[222]
		if len(location.GetPath()) == 2 && location.Path[0] == 4 {
			msgIndex := location.Path[1]
			SourceInfo.Messages = append(SourceInfo.Messages, MessageInfo{
				Name:       *descr.MessageType[msgIndex].Name,
				Message:    *descr.MessageType[msgIndex],
				Info:       location,
				FieldInfos: []FieldInfo{},
			})
		}

		// 4 111 2 222 =>	 4 MessageIndex 2 FieldIndex
		// for field with index 222 in message with index 111
		// location info for descriptor.MessageType[111]Field[222]
		if len(location.GetPath()) == 4 && location.Path[0] == 4 && location.Path[2] == 2 {
			msgIndex := location.Path[1]
			fieldIndex := location.Path[3]
			fi := FieldInfo{
				Name:    *descr.MessageType[msgIndex].Field[fieldIndex].Name,
				Info:    location,
				Field:   descr.MessageType[msgIndex].Field[fieldIndex],
				Message: *descr.MessageType[msgIndex],
			}
			SourceInfo.Messages[msgIndex].FieldInfos = append(SourceInfo.Messages[msgIndex].FieldInfos, fi)
		}

	}
	return SourceInfo
}
