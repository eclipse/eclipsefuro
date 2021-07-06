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
	"github.com/eclipse/eclipsefuro/furo/internal/cmd/exportAsYaml"

	"github.com/spf13/cobra"
)

// exportAsYamlCmd represents the exportAsYaml command
var exportAsYamlCmd = &cobra.Command{
	Use:   "exportAsYaml",
	Short: "Exports all specs and the current config in one yaml file to stdout",
	Long: `Use this for your chain of generators...
	
You will get a yaml with all types and services and the config.
Feel free to add custom sections in the config to use them in custom commands or scripts.

    services:
       your.Service: ...
    types:
       your.type: ...
    config:
       module: mod
       custom:
          remoteDir: "path/to/somewhere"
          otherCustomSetting: true
    

`,
	Run: exportAsYaml.Run,
}

func init() {
	rootCmd.AddCommand(exportAsYamlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportAsYamlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportAsYamlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	exportAsYamlCmd.Flags().BoolP("full", "f", false, "Include the ast info")
}
