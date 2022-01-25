package main_test

import (
	"github.com/eclipse/eclipsefuro/furo/cmd"
	"github.com/eclipse/eclipsefuro/furo/test"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestRunCommand(t *testing.T) {
	dir, _ := test.CwdTestDir()
	defer test.RemoveTestDir(dir)

	os.Args = []string{"cmd", "init", "--repository=github.com/yourname/furo-test"}
	rco := cmd.RootCmd
	rco.Execute()

	require.Equal(t, true, test.FileExist(path.Join(dir, ".furo")))
	require.Equal(t, "427e3452ab2c1c6b0ab1d147564b783c", test.MustMd5Sum(path.Join(dir, ".furo")))

	os.Args = []string{"cmd", "install"}
	rco.Execute()

	// execute the default flow
	os.Args = []string{"cmd", "run", "default"}
	rco.Execute()

	require.Equal(t, "112058c1b1d8cb64821d30e6f6864570", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Sample.type.spec")))
	require.Equal(t, "6d172492e249b527849e3f04a3c4534d", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Samples.service.spec")))

	require.Equal(t, true, test.FileExist(path.Join(dir, "dist", "protos")))
	require.Equal(t, true, test.FileExist(path.Join(dir, "dist", "env.js")))

}
