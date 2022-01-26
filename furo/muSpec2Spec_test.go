package main_test

import (
	"github.com/eclipse/eclipsefuro/furo/cmd"
	"github.com/eclipse/eclipsefuro/furo/test"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestMuSpec2SpecCommand(t *testing.T) {
	dir, _ := test.CwdTestDir()

	defer test.RemoveTestDir(dir)

	// init
	os.Args = []string{"cmd", "init", "--repository=github.com/yourname/furo-test"}
	rco := cmd.RootCmd
	rco.Execute()

	// install
	os.Args = []string{"cmd", "install"}
	rco.Execute()

	os.Args = []string{"cmd", "muSpec2Spec"}
	rco.Execute()

	require.Equal(t, true, test.FileExist(path.Join(dir, "specs", "sample")))
	require.Equal(t, "f0fab8c3b2c45792a676b418352aa9c6", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Sample.type.spec")))
	require.Equal(t, "1a6284ebd7153476d831557f253d0757", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Samples.service.spec")))

	os.Args = []string{"cmd", "checkImports"}
	rco.Execute()

	require.Equal(t, true, test.FileExist(path.Join(dir, "specs", "sample")))
	require.Equal(t, "f0fab8c3b2c45792a676b418352aa9c6", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Sample.type.spec")))
	require.Equal(t, "6d172492e249b527849e3f04a3c4534d", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Samples.service.spec")))

}
