package generator

import (
	"bytes"
	"encoding/json"
	"github.com/bufbuild/protoplugin"
	openapi_v3 "github.com/google/gnostic/openapiv3"
	"protoc-gen-open-models/pkg/sourceinfo"
	"slices"
	"strings"
	"text/template"
)

type ModelType struct {
	Name            string
	Fields          []ModelFields
	LeadingComments []string
	MetaTypeName    string
	RequiredFields  string
	ReadonlyFields  string
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
	FieldConstructor    string // constructor for the field / usually it is the same as ModelType
	Constraints         string // Openapi Constraints as json literal
}

type FieldConstraints struct {
	Nullable         bool    `json:"nullable,omitempty"`
	ReadOnly         bool    `json:"read_only,omitempty"`
	WriteOnly        bool    `json:"write_only,omitempty"`
	Deprecated       bool    `json:"deprecated,omitempty"`
	Title            string  `json:"title,omitempty"`
	MultipleOf       float64 `json:"multiple_of,omitempty"`
	Maximum          float64 `json:"maximum,omitempty"`
	ExclusiveMaximum bool    `json:"exclusive_maximum,omitempty"`
	Minimum          float64 `json:"minimum,omitempty"`
	ExclusiveMinimum bool    `json:"exclusive_minimum,omitempty"`
	MaxLength        int64   `json:"max_length,omitempty"`
	MinLength        int64   `json:"min_length,omitempty"`
	Pattern          string  `json:"pattern,omitempty"`
	MaxItems         int64   `json:"max_items,omitempty"`
	MinItems         int64   `json:"min_items,omitempty"`
	UniqueItems      bool    `json:"unique_items,omitempty"`
	MaxProperties    int64   `json:"max_properties,omitempty"`
	MinProperties    int64   `json:"min_properties,omitempty"`
	Required         bool    `json:"required,omitempty"`
	Type             string  `json:"type,omitempty"`
	Description      string  `json:"description,omitempty"`
	Format           string  `json:"format,omitempty"`
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
        FieldConstructor: {{.FieldConstructor}},
        {{- if ne .MAPValueConstructor ""}}
        ValueConstructor: {{.MAPValueConstructor}},{{end}}{{if .Constraints}}
        constraints: {{.Constraints}},{{end}}
      }{{end}}
    ];

    // Initialize the fields
{{- range .Fields}}{{if .LeadingComments}}{{range $i, $commentLine := .LeadingComments}}
    // {{$commentLine}}{{end}}{{end}}
    this._{{.FieldName}} = new {{.ModelType}}(undefined,{{if eq .Kind "TYPE_ENUM"}}{{.SetterType}}, {{.SetterType}}.{{.EnumDefault}}, {{end}} this, '{{.FieldName}}'); 
{{end}}


    // Set required fields
    [{{.RequiredFields}}].forEach(fieldName => {
      (
        this[fieldName as keyof {{.Name}}] as FieldNode
      ).__meta.required = true;
    });

    // Set readonly fields
    [{{.ReadonlyFields}}].forEach(fieldName => {
      (
        this[fieldName as keyof {{.Name}}] as FieldNode
      ).__readonly = true;
    });

    // Default values from openAPI annotations
    this.__defaultValues = {};

    // Initialize the fields with init data
    if (initData !== undefined) {
      this.__fromLiteral({ ...this.__defaultValues, ...initData });
    } else {
      this.__fromLiteral(this.__defaultValues);
    }

    this.__meta.isPristine = true;
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
	reqFields := []string{}
	if message.OpenApiSchema != nil {
		reqFields = message.OpenApiSchema.Required
	}
	readonlyFields := []string{}

	modelType := ModelType{
		Name:            fullQualifiedName(message.Package, message.Name),
		Fields:          nil,
		LeadingComments: multilineComment(message.Info.GetLeadingComments()),
		MetaTypeName:    message.Package + "." + message.Name,
	}

	defaults := map[string]openapi_v3.DefaultType{}

	if len(reqFields) > 0 {
		modelType.RequiredFields = "'" + strings.Join(reqFields, "', '") + "'"
	}

	for _, field := range message.FieldInfos {
		enumDefault := ""
		if field.Field.Type.String() == "TYPE_ENUM" {
			enumDefault = resolveFirstEnumOptionForField(field, si, request)
		}
		m, sc, st, gt, mapValueConstructor, fc := resolveModelType(imports, field)
		var constraints string
		if field.OpenApiProperties != nil {
			// check for readonly fields
			if field.OpenApiProperties.ReadOnly {
				readonlyFields = append(readonlyFields, field.Field.GetJsonName())
			}
			if field.OpenApiProperties.Default != nil {
				// collect the defaults
				defaults[field.Field.GetJsonName()] = *field.OpenApiProperties.Default
				// do not put the defaults in to the constraints
				field.OpenApiProperties.Default = nil
			}

			c, err := json.Marshal(field.OpenApiProperties)
			if err == nil {
				fieldConstraints := FieldConstraints{}
				json.Unmarshal(c, &fieldConstraints)
				if slices.Contains(reqFields, field.Field.GetName()) {
					fieldConstraints.Required = true
				}
				fieldConstraintsJson, _ := json.Marshal(fieldConstraints)
				constraints = string(fieldConstraintsJson)
			}

		}
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
			FieldConstructor:    fc,
			Constraints:         constraints,
		})
	}
	if len(readonlyFields) > 0 {
		modelType.ReadonlyFields = "'" + strings.Join(readonlyFields, "', '") + "'"
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
