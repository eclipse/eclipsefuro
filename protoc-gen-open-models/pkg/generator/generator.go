package generator

import (
	"github.com/bufbuild/protoplugin"
	"path"
	"protoc-gen-open-models/pkg/sourceinfo"
	"strings"
)

var projectFiles = map[string]string{}

func GenerateAll(responseWriter protoplugin.ResponseWriter, request protoplugin.Request) {

	for _, fileDescriptorProto := range request.AllFileDescriptorProtos() {
		if strings.HasPrefix(fileDescriptorProto.GetPackage(), "openapi.v3") {
			continue
		}
		// collect source infos like messages,enums,services,... including comments
		sourceInfo := sourceinfo.GetSourceInfo(fileDescriptorProto)
		// build a list of project files
		for _, enum := range sourceInfo.Enums {
			projectFiles[path.Join(sourceInfo.Path, enum.Name)] = "ENUM"
		}
		for _, service := range sourceInfo.Services {
			projectFiles[path.Join(sourceInfo.Path, service.Name)] = "SERVICE"

		}
		for _, message := range sourceInfo.Messages {
			projectFiles[path.Join(sourceInfo.Path, message.Name)] = "MESSAGE"
		}
	}

	for _, fileDescriptorProto := range request.FileDescriptorProtosToGenerate() {
		if strings.HasPrefix(fileDescriptorProto.GetPackage(), "openapi.v3") {
			continue
		}
		// collect source infos like messages,enums,services,... including comments
		si := sourceinfo.GetSourceInfo(fileDescriptorProto)
		Generate(si, responseWriter, request)
	}
}

func Generate(sourceInfo sourceinfo.SourceInfo, responseWriter protoplugin.ResponseWriter, request protoplugin.Request) {

	// build enum types
	// the enums are used by T(ransport), L(iteral) and Models
	for _, enum := range sourceInfo.Enums {
		// Add the response file to the response.
		responseWriter.AddFile(
			path.Join(sourceInfo.Path, enum.Name+".ts"),
			createEnum(sourceInfo, &enum),
		)
	}

	// build up the services
	for _, service := range sourceInfo.Services {
		// Add the response file to the response.
		responseWriter.AddFile(
			path.Join(sourceInfo.Path, service.Name+".ts"),
			createOpenModelService(sourceInfo, service),
		)
	}

	// build model files, they include the Literal, Transport and Model
	for _, message := range sourceInfo.Messages {
		// Add the response file to the response.
		responseWriter.AddFile(
			path.Join(sourceInfo.Path, message.Name+".ts"),
			createOpenModel(sourceInfo, &message, request),
		)
	}
}

func createEnum(si sourceinfo.SourceInfo, enum *sourceinfo.EnumInfo) string {
	options := []string{}

	for _, en := range enum.ValuesInfo {
		if en.Info.GetLeadingComments() != "" {
			options = append(options, multilineCommentString(en.Info.GetLeadingComments()))
		}
		o := "  " + en.Name + " = \"" + en.Name + "\","
		if en.Info.GetTrailingComments() != "" {
			o = o + " //" + en.Info.GetTrailingComments()
			if strings.HasSuffix(o, "\n") {
				o = o[:len(o)-1]
			}
		}
		options = append(options, o)
	}

	content := "export enum " + fullQualifiedName(enum.Package, enum.Name) + " {\n" + strings.Join(options, "\n") + "\n}"
	parts := []string{
		"// Code generated by furo protoc-gen-open-models. DO NOT EDIT.",
		"// protoc-gen-open-models version: ????",
		"",
		multilineCommentString(enum.Info.GetLeadingComments()),
		content,
	}
	return string(strings.Join(parts, "\n"))
}

func createOpenModel(si sourceinfo.SourceInfo, message *sourceinfo.MessageInfo, request protoplugin.Request) string {

	imports := ImportMap{Imports: make(map[string]map[string]bool)}
	imports.AddImport("@furo/open-models/dist/index", "FieldNode")
	imports.AddImport("@furo/open-models/dist/index", "Registry")

	literalType := prepareLiteralType(message, imports)
	transportType := prepareTransportType(message, imports)
	modelType := prepareModelType(message, imports, si, request)

	parts := []string{
		"// Code generated by furo protoc-gen-open-models. DO NOT EDIT.",
		"// protoc-gen-open-models version: ????",
		imports.Render(),
		literalType.Render(),
		transportType.Render(),
		modelType.Render(),
	}
	return string(strings.Join(parts, "\n"))

}

func createOpenModelService(si sourceinfo.SourceInfo, service sourceinfo.ServiceInfo) string {
	parts := []string{
		"// Code generated by furo protoc-gen-open-models. DO NOT EDIT.",
		"// protoc-gen-open-models version: ????",
		service.Name,
	}
	return string(strings.Join(parts, "\n"))
}
