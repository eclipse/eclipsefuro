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
	"github.com/eclipse/eclipsefuro/furo/internal/cmd/muSpec2Spec"
	"github.com/spf13/cobra"
)

// muSpec2SpecCmd represents the muSpec2Spec command
var muSpec2SpecCmd = &cobra.Command{
	Use:   "muSpec2Spec",
	Short: "Updates the type specs with the definitions from the type µSpecs.",
	Long: `The converter will update your type specs and also delete specs and fields if they are not in the µSpec file anymore.

Do not forget to set your µSpec folder in the .furo config.`,
	Run: muSpec2Spec.Run,
}

// needed for the documentation generator
var MuSpec2SpecCmd = muSpec2SpecCmd

func init() {
	rootCmd.AddCommand(muSpec2SpecCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// muSpec2SpecCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	muSpec2SpecCmd.Flags().BoolP("delete", "d", false, "Delete specs which are not in muTypes")
	muSpec2SpecCmd.Flags().BoolP("overwrite-spec-options", "", false, "Overwrite the proto options section in the spec files")
}
