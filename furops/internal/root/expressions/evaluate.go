package expressions

import (
	"github.com/eclipse/eclipsefuro/furops/internal/root/specs"
	"github.com/iancoleman/strcase"
	"gopkg.in/Knetic/govaluate.v3"
	"log"
)

var functions = map[string]govaluate.ExpressionFunction{
	"Strlen": func(args ...interface{}) (interface{}, error) {
		length := len(args[0].(string))
		return length, nil
	},
	"ToCamel": func(args ...interface{}) (interface{}, error) {
		return strcase.ToCamel(args[0].(string)), nil
	},
	"ToSnake": func(args ...interface{}) (interface{}, error) {
		return strcase.ToSnake(args[0].(string)), nil
	},
	"ToLowerCamel": func(args ...interface{}) (interface{}, error) {
		return strcase.ToLowerCamel(args[0].(string)), nil
	},
	"ToSnakeWithIgnore": func(args ...interface{}) (interface{}, error) {
		return strcase.ToSnakeWithIgnore(args[0].(string), args[1].(string)), nil
	},
	"ToScreamingSnake": func(args ...interface{}) (interface{}, error) {
		return strcase.ToScreamingSnake(args[0].(string)), nil
	},
	"ToKebab": func(args ...interface{}) (interface{}, error) {
		return strcase.ToKebab(args[0].(string)), nil
	},
	"ToScreamingKebab": func(args ...interface{}) (interface{}, error) {
		return strcase.ToScreamingKebab(args[0].(string)), nil
	},
	"GetService": func(args ...interface{}) (interface{}, error) {
		var servicespec interface{}
		var found bool
		servicespec, found = specs.Specs.Services[args[0].(string)]
		if !found {
			servicespec = specs.Specs.InstalledServices[args[0].(string)]
		}
		return servicespec, nil
	},
	"GetType": func(args ...interface{}) (interface{}, error) {
		var servicespec interface{}
		var found bool
		servicespec, found = specs.Specs.Types[args[0].(string)]
		if !found {
			servicespec = specs.Specs.InstalledTypes[args[0].(string)]
		}
		return servicespec, nil
	},
}

func EvaluateExpression(parameters map[string]interface{}, expressionstring string) interface{} {
	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expressionstring, functions)
	if err != nil {
		log.Fatalln(err, expressionstring)
	}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}
