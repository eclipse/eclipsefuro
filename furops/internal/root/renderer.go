package root

import (
	"bufio"
	"github.com/eclipse/eclipsefuro/furops/internal/root/expressions"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"text/template"
)

type templatedata struct {
	Var  map[string]interface{}
	Conf FPS
}

func RenderTemplates(vars map[string]interface{}, pattern FPS) {

	for _, structure := range pattern.Structure {

		// execute the templates with the collected data
		tdata := templatedata{
			Var:  vars,
			Conf: pattern,
		}

		// check for repeats and expressions and prepare the variables
		if structure.RepeatBy != "" {

			for _, stringitem := range vars[structure.RepeatBy].([]string) {
				vars[structure.RepeatAs] = stringitem

				// evaluate structure expressions
				for varname, expression := range structure.Expressions {
					vars[varname] = expressions.EvaluateExpression(vars, expression)
				}

				// if condition was set and results to false, skip the item
				if structure.Condition != "" {
					if expressions.EvaluateExpression(vars, structure.Condition) == false {
						continue
					}
				}
				RenderToFile(vars, pattern.Path, structure, tdata)
			}
		} else {

			// if condition was set and results to false, skip the var
			if structure.Condition != "" {
				if expressions.EvaluateExpression(vars, structure.Condition) == false {
					continue
				}
			}

			RenderToFile(vars, pattern.Path, structure, tdata)
		}

	}
}

func RenderToFile(vars map[string]interface{}, templatedir string, structure Structure, templatedata templatedata) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(path.Dir(viper.ConfigFileUsed()))

	templatefile := path.Join(templatedir, structure.Template)
	target := expressions.EvaluateExpression(vars, structure.Target)
	tmpl, err := template.New(path.Base(structure.Template)).ParseFiles(templatefile)
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(cwd)

	if target == "STDOUT" {
		err = tmpl.Execute(os.Stdout, templatedata)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	tpath := path.Dir(target.(string))
	err = os.MkdirAll(tpath, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	f, err := os.Create(target.(string))
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)

	err = tmpl.Execute(w, templatedata)
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()
}
