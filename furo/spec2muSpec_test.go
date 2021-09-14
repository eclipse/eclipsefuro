package main_test

import (
	"github.com/eclipse/eclipsefuro/furo/cmd"
	"github.com/eclipse/eclipsefuro/furo/test"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestSpec2muSpecCommand(t *testing.T) {
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
	require.Equal(t, "4db69d8f0d25911b8035e1bdacd1952e", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Sample.type.spec")))
	require.Equal(t, "1a6284ebd7153476d831557f253d0757", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Samples.service.spec")))

	os.Args = []string{"cmd", "checkImports"}
	rco.Execute()

	require.Equal(t, true, test.FileExist(path.Join(dir, "specs", "sample")))
	require.Equal(t, "4db69d8f0d25911b8035e1bdacd1952e", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Sample.type.spec")))
	require.Equal(t, "6d172492e249b527849e3f04a3c4534d", test.MustMd5Sum(path.Join(dir, "specs", "sample", "Samples.service.spec")))

	// delete the muspecs
	test.RemoveTestDir(path.Join(dir, "muspecs"))

	require.Equal(t, false, test.FileExist(path.Join(dir, "muspecs")))

	os.Args = []string{"cmd", "spec2muSpec"}
	rco.Execute()
	require.Equal(t, true, test.FileExist(path.Join(dir, "muspecs")))

	// initial were 2 files in muspec, spec2Muspec creates one file per spec, so we have 4 files now
	require.Equal(t, "83f97c7fd66d5b9fd34069e075810a80", test.MustMd5Sum(path.Join(dir, "muspecs", "sample", "Sample.types.yaml")))
	require.Equal(t, "102b3562851cc4bac750b6d409344f6e", test.MustMd5Sum(path.Join(dir, "muspecs", "sample", "SampleCollection.types.yaml")))
	require.Equal(t, "cecdceb48471dc1b2ebc958ce30e97bb", test.MustMd5Sum(path.Join(dir, "muspecs", "sample", "SampleEntity.types.yaml")))
	require.Equal(t, "c7e54a185c29003ba0aad3ccd9bdb684", test.MustMd5Sum(path.Join(dir, "muspecs", "sample", "Samples.services.yaml")))

}
