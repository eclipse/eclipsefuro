package test

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

var Sourcedir = ""

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// CwdTestDir Creates a temp directory and sets it as working directory
func CwdTestDir() (dir string, err error) {
	dir, err = ioutil.TempDir("", "furo-test")
	Sourcedir, _ = os.Getwd() // store sourcedir for copy mechanism
	os.Chdir(dir)
	return dir, err
}

func CopyTestFile(relativesourcepath string, absolutetargetpath string) error {
	src := path.Join(Sourcedir, relativesourcepath)

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(absolutetargetpath)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)

	return err
}

func RemoveTestDir(dir string) {
	if err := os.RemoveAll(dir); err != nil {
		panic(err)
	}
}

func MustMd5Sum(path string) string {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return returnMD5String
	}

	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String
	}

	//Get the hash
	hashInBytes := hash.Sum(nil)[:16]

	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String

}
