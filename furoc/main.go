package main

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/pkg/util"
	"github.com/eclipse/eclipsefuro/furoc/internal/input"
	"github.com/eclipse/eclipsefuro/furoc/internal/subcommand"
	"github.com/eclipse/eclipsefuro/furoc/pkg/parseargs"
	"google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {
	var arglist parseargs.Arglist
	fmt.Println("furoc 1.31.0")
	if len(os.Args) == 1 {
		// look for a .furoc config in cwd
		// if we are in a spec project and have furoc instructions there, follow them
		if !util.FileExists(".furoc") {
			log.Fatal("Config file .furoc not found. You must have a .furoc config when no arguments are given")
		} else {
			// read config
			arglist = parseargs.FromFurocConfig()
		}

	} else {
		// parse furoc command arguments
		arglist = parseargs.Parse()
	}

	if len(arglist.Inputs) == 0 {
		log.Fatal("No input given.")
	}

	specDir := arglist.Inputs[0]
	err, specYaml := input.GetInputYaml(specDir,
		exec.Command("furo", "exportAsYaml", "-f"))
	if err != nil {
		log.Fatal(err)
	}

	type CMDResponse struct {
		response      *pluginpb.CodeGeneratorResponse
		baseTargetDir string
	}
	allResponses := []CMDResponse{}
	//  for duplicate file check
	fullFilelist := map[string]bool{}

	for _, cmd := range arglist.Commands {
		r, err := subcommand.ExecuteSubcommand(cmd.Plugin, specYaml, cmd.Args)
		if err != nil {
			log.Fatal(err)
		}
		allResponses = append(allResponses, CMDResponse{
			response:      r,
			baseTargetDir: cmd.OutputDir,
		})
		// check for duplicate files
		for _, f := range r.File {
			fname := cmd.OutputDir + "/" + *f.Name
			_, alreadyRagistred := fullFilelist[fname]
			if alreadyRagistred {
				log.Fatal(fname, " try to write same file twice")
			} else {
				fullFilelist[fname] = true
			}
		}
	}

	// Writer:

	for _, responseSet := range allResponses {
		for _, file := range responseSet.response.File {
			if util.DirExists(responseSet.baseTargetDir) {
				fname := path.Join(responseSet.baseTargetDir, *file.Name)
				util.MkdirRelative(path.Dir(fname))
				ioutil.WriteFile(fname, []byte(*file.Content), 0644)
			} else {
				log.Fatal("Target directory does not exist: ", responseSet.baseTargetDir)
			}

		}
	}

}
