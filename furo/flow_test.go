package main_test

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/cmd"
	"github.com/eclipse/eclipsefuro/furo/test"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestRunCommand(t *testing.T) {
	dir, _ := test.CwdTestDir()
	// defer test.RemoveTestDir(dir)
	fmt.Println(dir)
	os.Args = []string{"cmd", "init", "--repository=github.com/yourname/furo-test"}
	rco := cmd.RootCmd
	rco.Execute()

	require.Equal(t, true, test.FileExist(path.Join(dir, ".furo")))
	require.Equal(t, "f4ba32ae7424570c4b7520432ce89e94", test.MustMd5Sum(path.Join(dir, ".furo")))

	err := test.CopyTestFile("test/testdata/.furo", path.Join(dir, ".furo"))
	require.NoError(t, err, "Must Copy .furo file")
	require.Equal(t, "e7b51c770c547d672201af43e001a61f", test.MustMd5Sum(path.Join(dir, ".furo")))

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
