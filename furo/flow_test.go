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
	require.Equal(t, "d8a9dbdca98179d9a7d3ebbdcb2c3fa8", test.MustMd5Sum(path.Join(dir, ".furo")))

	err := test.CopyTestFile("test/testdata/.furo", path.Join(dir, ".furo"))
	require.NoError(t, err, "Must Copy .furo file")
	require.Equal(t, "8493075c067f312d3ffd760363e3e0f6", test.MustMd5Sum(path.Join(dir, ".furo")))

	os.Args = []string{"cmd", "install"}
	rco.Execute()

	// execute the default flow
	os.Args = []string{"cmd", "run", "default"}
	rco.Execute()

	require.Equal(t, "4db69d8f0d25911b8035e1bdacd1952e", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Sample.type.spec")))
	require.Equal(t, "6d172492e249b527849e3f04a3c4534d", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Samples.service.spec")))

	require.Equal(t, true, test.FileExist(path.Join(dir, "dist", "protos")))
	require.Equal(t, true, test.FileExist(path.Join(dir, "dist", "env.js")))

}