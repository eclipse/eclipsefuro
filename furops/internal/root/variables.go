package root

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/eclipse/eclipsefuro/furops/internal/root/expressions"
	"github.com/eclipse/eclipsefuro/furops/internal/root/suggester"
	"log"
	"regexp"
	"strconv"
)

//func (d prompt.Document) []prompt.Suggest{}

func queryVariables(fps FPS) map[string]interface{} {

	prompters := map[string]func(d prompt.Document) []prompt.Suggest{
		"string":    suggester.Stringinput,
		"type":      suggester.Typecompleter,
		"service":   suggester.Servicecompleter,
		"directory": suggester.Directory,
		"number":    suggester.Number,
		"bool":      suggester.Bool,
	}

	res := map[string]interface{}{}

	for _, conf := range fps.Variables {
		// if condition was set and results to false, skip the var
		if conf.Condition != "" {
			if expressions.EvaluateExpression(res, conf.Condition) == false {
				continue
			}
		}

		if conf.Expression != "" {
			res[conf.VarName] = expressions.EvaluateExpression(res, conf.Expression)
			continue
		}

		done := false
		p, av := prompters[conf.InputKind]
		if !av {
			fmt.Println(conf.VarName, " input kind ", conf.InputKind, " not supported")
			continue
		}

		initialText := ""
		if conf.Default != "" {
			it := expressions.EvaluateExpression(res, conf.Default)
			initialText = it.(string)
		}
		Clear()

		for !done {

			opts := applyTheme()
			opts = append(opts, prompt.OptionInitialBufferText(initialText))
			opts = append(opts, prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.ControlC,
				Fn:  exit,
			}))
			if conf.InputKind == "directory" {
				opts = append(opts, prompt.OptionCompletionWordSeparator("/"))
			}

			fmt.Println(conf.Prompt)
			input := prompt.Input(conf.VarName+": ", p, opts...)

			// check for regexp and re query if it is not fulfilled
			if conf.Regexp != "" {
				matched, err := regexp.MatchString(conf.Regexp, input)
				if err != nil {
					log.Fatal("regexp error for ", conf.VarName, err)
				}
				done = matched
				if !matched {
					fmt.Println("Input ", input, " does match pattern ", conf.RegexpText)
					initialText = input
				}
			} else {
				done = true
			}
			switch conf.InputKind {
			case "number":
				val, err := strconv.ParseFloat(input, 64)
				if err == nil {
					res[conf.VarName] = val
				} else {
					res[conf.VarName] = 0
				}

				break
			case "bool":
				if input == "true" {
					res[conf.VarName] = true
				} else {
					res[conf.VarName] = false
				}

				break
			default:
				res[conf.VarName] = input

			}

		}
	}

	return res
}
