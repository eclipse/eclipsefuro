package root

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/eclipse/eclipsefuro/furops/internal/root/specs"
	"github.com/spf13/cobra"
	"os"
)

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

	RenderTemplates(vars, pattern)
}

func exit(buffer *prompt.Buffer) {
	os.Exit(0)
}
