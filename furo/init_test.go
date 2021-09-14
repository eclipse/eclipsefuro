package main_test

import (
	"github.com/eclipse/eclipsefuro/furo/cmd"
	"github.com/eclipse/eclipsefuro/furo/test"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestInitCommand(t *testing.T) {
	dir, _ := test.CwdTestDir()
	defer test.RemoveTestDir(dir)

	os.Args = []string{"cmd", "init", "--repository=github.com/yourname/furo-test"}
	rco := cmd.RootCmd
	rco.Execute()

	require.Equal(t, true, test.FileExist(path.Join(dir, ".furo")))
	require.Equal(t, "f4ba32ae7424570c4b7520432ce89e94", test.MustMd5Sum(path.Join(dir, ".furo")))

}
