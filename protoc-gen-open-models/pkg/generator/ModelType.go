package generator

import (
	"bytes"
	"github.com/bufbuild/protoplugin"
	"protoc-gen-open-models/pkg/sourceinfo"
	"text/template"
)

type ModelType struct {
	Name            string
	Fields          []ModelFields
	LeadingComments []string
	MetaTypeName    string
}

type ModelFields struct {
	LeadingComments     []string
	TrailingComment     string
	FieldName           string // name of the field in interface types and models
	FieldTransportName  string // name of the field in transport types
	ModelType           string // like STRING, ENUM<SomeTHING>, Int32Value, FuroTypeMessage
	SetterCommand       string // for primitive or typesetter
	SetterType          string // string or LFuroTypeMessage
	GetterType          string // string or FuroTypeMessage
	Kind                string // proto field type like TYPE_MESSAGE
	EnumDefault         string // Name of the first enum option
	MAPValueConstructor string // value constructor for MAP
}

var ModelTypeTemplate = `{{if .LeadingComments}}{{range $i, $commentLine := .LeadingComments}}
// {{$commentLine}}{{end}}{{end}}
export class {{.Name}} extends FieldNode {
{{range .Fields}}
  {{if .LeadingComments}}{{range $i, $commentLine := .LeadingComments}}
  // {{$commentLine}}{{end}}{{end}}
  private _{{.FieldName}}: {{.ModelType}}; // trailing comment
{{end}}
 public __defaultValues: L{{.Name}};

  constructor(
    initData?: L{{.Name}},
    parent?: FieldNode,
    parentAttributeName?: string,
  ) {
    super(undefined, parent, parentAttributeName);
    this.__meta.typeName = '{{.MetaTypeName}}';

    this.__meta.nodeFields = [
{{- $first := true }}{{range .Fields}}{{if not $first}}, {{else}}{{$first = false}}{{end}}
      {
        fieldName: '{{.FieldName}}',
        jsonName: '{{.FieldTransportName}}',
        FieldConstructor: {{.ModelType}},
        {{- if ne .MAPValueConstructor ""}}
        ValueConstructor: {{.MAPValueConstructor}},{{end}}
        // constraints: { max_items: 4, required: true },
      }{{end}}
    ];

    // Initialize the fields
{{- range .Fields}}{{if .LeadingComments}}{{range $i, $commentLine := .LeadingComments}}
    // {{$commentLine}}{{end}}{{end}}
    this._{{.FieldName}} = new {{.ModelType}}(undefined,{{if eq .Kind "TYPE_ENUM"}}{{.SetterType}}, {{.SetterType}}.{{.EnumDefault}}, {{end}} this, '{{.FieldName}}'); 
{{end}}

    // Default values from openAPI annotations
    this.__defaultValues = {};

    // Initialize the fields with init data
    if (initData !== undefined) {
      this.__fromLiteral({ ...this.__defaultValues, ...initData });
    } else {
      this.__fromLiteral(this.__defaultValues);
    }
  }
{{range .Fields}}
{{if .LeadingComments}}{{range $i, $commentLine := .LeadingComments}}
  // {{$commentLine}}{{end}}{{end}}
  public get {{.FieldName}}(): {{.GetterType}} {
    return this._{{.FieldName}};
  }

  public set {{.FieldName}}(v: {{.SetterType}}) {
    this.{{.SetterCommand}}(this._{{.FieldName}}, v);
  }
{{end}}

  fromLiteral(data: L{{.Name}}) {
    super.__fromLiteral(data);
  }

  toLiteral(): L{{.Name}} {
    return super.__toLiteral();
  }
}

Registry.register('{{.MetaTypeName}}', {{.Name}});
`

func (r *ModelType) Render() string {

	t, err := template.New("ModelType").Parse(ModelTypeTemplate)
	if err != nil {
		panic(err)
	}

	var res bytes.Buffer
	err = t.Execute(&res, r)
	if err != nil {
		panic(err)
	}
	return res.String()
}

func prepareModelType(message *sourceinfo.MessageInfo, imports ImportMap, si sourceinfo.SourceInfo, request protoplugin.Request) ModelType {
	modelType := ModelType{
		Name:            fullQualifiedName(message.Package, message.Name),
		Fields:          nil,
		LeadingComments: multilineComment(message.Info.GetLeadingComments()),
		MetaTypeName:    message.Package + "." + message.Name,
	}
	for _, field := range message.FieldInfos {
		enumDefault := ""
		if field.Field.Type.String() == "TYPE_ENUM" {
			enumDefault = resolveFirstEnumOptionForField(field, si, request)
		}
		m, sc, st, gt, mapValueConstructor := resolveModelType(imports, field)

		modelType.Fields = append(modelType.Fields, ModelFields{
			LeadingComments:     multilineComment(field.Info.GetLeadingComments()),
			TrailingComment:     field.Info.GetTrailingComments(),
			FieldName:           field.Field.GetJsonName(), // todo: check preserve proto names
			FieldTransportName:  field.Field.GetName(),     // todo: check  preserve proto names
			ModelType:           m,
			SetterCommand:       sc,
			SetterType:          st,
			GetterType:          gt,
			Kind:                field.Field.Type.String(),
			EnumDefault:         enumDefault,
			MAPValueConstructor: mapValueConstructor,
		})
	}
	return modelType
}

func resolveFirstEnumOptionForField(field sourceinfo.FieldInfo, si sourceinfo.SourceInfo, request protoplugin.Request) string {
	field.Field.Label.String()
	var enumSi []sourceinfo.EnumInfo
	// find the correct descriptor
Exit:
	for _, proto := range request.AllFileDescriptorProtos() {
		for _, descriptorProto := range proto.EnumType {
			if "."+proto.GetPackage()+"."+descriptorProto.GetName() == field.Field.GetTypeName() {
				enumSi = sourceinfo.GetSourceInfo(proto).Enums
				continue Exit
			}

		}
	}

	for _, info := range enumSi {
		if "."+field.Package+"."+info.Name == field.Field.GetTypeName() {
			return info.ValuesInfo[0].Name
		}
	}
	return "could not resolve " + field.Name
}
