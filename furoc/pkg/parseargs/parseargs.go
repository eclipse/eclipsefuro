package parseargs

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

type Arglist struct {
	Commands []Command `yaml:"Commands"`
	Inputs   []string  `yaml:"Inputs"` // todo: at the moment single input is supported only, the other inputs will be ignored
	Binary   string    `yaml:"Binary"`
}

type Command struct {
	Args       []string `yaml:"Args"`
	Plugin     string   `yaml:"Plugin"`
	PluginName string   `yaml:"PluginName"` // binary
	OutputDir  string   `yaml:"OutputDir"`
}

func Parse() Arglist {
	a := Arglist{
		Commands: []Command{},
		Inputs:   []string{},
		Binary:   "",
	}
	a.Binary = os.Args[0]

	plugins := map[string]string{}

	// triage for the commands
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-I") {
			a.Inputs = append(a.Inputs, arg[2:])
		}
		// plugins
		if strings.HasPrefix(arg, "--plugin") {
			p := path.Base(arg[9:])
			if strings.HasPrefix(p, "furoc-gen-") {
				pth := strings.Split(arg, "=")
				plugins[p[10:]] = pth[1]
			} else {
				log.Fatal("Malformed arg: ", arg)
			}
		}

	}
	// command options
	// can be like this
	// --u33e_out= \
	// Sreference-search,   \
	// Scollection-dropdown,\
	//:outputBaseDirectoryForU33e
	//
	// Or
	// --u33e_out=Sreference-search, Scollection-dropdown :outputBaseDirectoryForU33e
	argline := strings.Join(os.Args, " ")
	var regex = regexp.MustCompile(`--([^=]*)_out=([^:]*):([^\s]*)`)
	matches := regex.FindAllStringSubmatch(argline, -1)
	if len(matches) == 0 {
		fmt.Println("typeline not parseable", regex)
	}
	for _, m := range matches {
		pluginshortname := m[1]
		plugin := "furoc-gen-" + pluginshortname
		binary, ok := plugins[pluginshortname]
		if ok {
			plugin = binary
		}
		cmd := Command{
			Args:       []string{},
			Plugin:     plugin,
			PluginName: pluginshortname,
			OutputDir:  m[3],
		}

		// trim
		subargs := strings.Split(m[2], ",")
		for i, _ := range subargs {
			subargs[i] = strings.TrimSpace(subargs[i])
			if subargs[i] != "" {
				cmd.Args = append(cmd.Args, subargs[i])
			}
		}
		a.Commands = append(a.Commands, cmd)
	}
	return a
}
