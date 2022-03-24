package root

// FPS contains the config definition for a pattern
type FPS struct {
	Name        string      `yaml:"name"`
	Description string      `yaml:"description"`
	Variables   []Variable  `yaml:"variables"`
	Structure   []Structure `yaml:"structure"`
	Data        interface{} `yaml:"data"`
	Path        string
}

// Variable config for a single variable, if Expression is used, the prompt will not be used
type Variable struct {
	VarName    string `yaml:"varname"`
	InputKind  string `yaml:"inputKind"`
	Default    string `yaml:"default"`
	Prompt     string `yaml:"prompt"`
	Regexp     string `yaml:"regexp"`     // use this to validate the input
	RegexpText string `yaml:"regexpText"` // Textual description of regexp
	Expression string `yaml:"expression"` // expression to construct variable contents
	Condition  string `yaml:"condition"`  // bool expression to define if var is prompted

}

// Structure contains the templates and filename to produce (Target)
type Structure struct {
	Template  string `yaml:"template"`
	Notes     string `yaml:"notes"`
	Target    string `yaml:"target"`
	Condition string `yaml:"condition"` // bool expression to define if file is generated
}
