package subcommand

import (
	"bufio"
	"bytes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"log"
	"os"
	"os/exec"
)

func ExecuteSubcommand(command string, specYaml []byte, params []string) (files *pluginpb.CodeGeneratorResponse, err error) {
	if command == "" {
		log.Fatal("plugin not defined")
	}
	subProcess := exec.Command(command, params...)
	stdin, err := subProcess.StdinPipe()
	writer := bufio.NewWriter(stdin)
	if err != nil {
		log.Fatal(command, err)
	}

	var b bytes.Buffer
	subProcess.Stdout = &b
	subProcess.Stderr = os.Stderr

	if err = subProcess.Start(); err != nil { //Use start, not run
		log.Fatal(command, err)
	}

	writer.Write(specYaml)

	stdin.Close()
	subProcess.Wait()

	inp := b.Bytes()

	r := &pluginpb.CodeGeneratorResponse{}
	err = proto.Unmarshal(inp, r)
	if err != nil {
		log.Fatal(err)
	}

	return r, nil
}
