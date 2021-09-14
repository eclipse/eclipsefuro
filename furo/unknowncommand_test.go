package main_test

import (
	"github.com/eclipse/eclipsefuro/furo/cmd"
	"github.com/eclipse/eclipsefuro/furo/test"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestUnknownCommand(t *testing.T) {
	dir, _ := test.CwdTestDir()
	defer test.RemoveTestDir(dir)

	os.Args = []string{"cmd", "something"}
	rco := cmd.RootCmd

	require.Error(t, rco.Execute())

}
