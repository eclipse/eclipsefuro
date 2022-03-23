package root

import (
	"bufio"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/eclipse/eclipsefuro/furops/internal/root/expressions"
	"github.com/eclipse/eclipsefuro/furops/internal/root/specs"
	"github.com/spf13/cobra"
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

func Run(cmd *cobra.Command, args []string) {
	ResolvePatterns()

	// only accept valid pattern as input
	var patternFound bool
	var pattern FPS

	for !patternFound {
		Clear()
		fmt.Println("Please SELECT a pattern. (↓)")
		opts := applyTheme()
		opts = append(opts, prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn:  exit,
		}))
		p := prompt.Input("> ", patterncompleter, opts...)
		pattern, patternFound = patterns[p]
	}

	specs.ReadSpecs()

	// collect the variables
	vars := queryVariables(pattern)

	Clear()

	// execute the templates with the collected data
	tdata := templatedata{
		Var:  vars,
		Conf: pattern,
	}

	for _, patternconfig := range pattern.Structure {

		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		os.Chdir(path.Dir(viper.ConfigFileUsed()))

		templatefile := path.Join(pattern.Path, patternconfig.Template)
		target := expressions.EvaluateExpression(vars, patternconfig.Target)
		tmpl, err := template.New(patternconfig.Template).ParseFiles(templatefile)
		if err != nil {
			log.Fatal(err)
		}
		os.Chdir(cwd)

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
		err = tmpl.Execute(w, tdata)
		if err != nil {
			log.Fatal(err)
		}
		w.Flush()

	}
}

func exit(buffer *prompt.Buffer) {
	os.Exit(0)
}
