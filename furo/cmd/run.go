/*
Copyright © 2020 Veith Zäch <veithz@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/eclipse/eclipsefuro/furo/internal/cmd/runner"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a configured flow.",
	Long: `Runs a configured flow of spectools commands.

To configure a flow of commands just add a "flow" in the flows section of your .furo config.
A flow is just a list of commands which gets executed in order

Example Config:

	[.furo]
	commands:
	  publish_npm: "./scripts/test.sh"
	flows:
	  type:
		- cleanTypeProtoDir
		- muSpec2Spec
		- TypeSpec2Proto
		- publish_npm

Command:

This config will run "cleanTypeProtoDir",  muSpec2Spec"" and "TypeSpec2Proto" in sequence and calling the command publish_npm

Tipp: If you need the types and services in your command, just call spectools again. 

Like:
    #!/bin/bash

    # generate the type documentation...
    spectools exportAsYaml | simple-generator -t scripts/typedoc.tpl > dist/typedoc.md

[example](../samples/typedoc/readme.md)
`,
	Run: runner.Run,
}

// needed for the documentation generator
var RunCmd = runCmd

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runCmd.Flags().StringP("flow", "f", "default", "A configured flow from the .furo config")
}
