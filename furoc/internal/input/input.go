package input

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

func GetInputYaml(specDir string, command *exec.Cmd) (error, []byte) {
	furo := command
	furo.Stderr = os.Stderr
	furo.Dir = specDir

	var b bytes.Buffer // buffer the furo output here
	furo.Stdout = &b

	err := furo.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = furo.Wait()
	if err != nil {
		log.Fatal(err)
	}

	specYaml := b.Bytes()
	return err, specYaml
}
