package main_test

import (
	"github.com/eclipse/eclipsefuro/furo/cmd"
	"github.com/eclipse/eclipsefuro/furo/test"
	"os"
	"testing"
)

func TestExportAsYamlCommand(t *testing.T) {
	dir, _ := test.CwdTestDir()

	defer test.RemoveTestDir(dir)

	// init
	os.Args = []string{"cmd", "init", "--repository=github.com/yourname/furo-test"}
	rco := cmd.RootCmd
	rco.Execute()

	// install
	os.Args = []string{"cmd", "install"}
	rco.Execute()

	os.Args = []string{"cmd", "exportAsYaml"}
	rco.Execute()

	// TODO: capture the output

}
