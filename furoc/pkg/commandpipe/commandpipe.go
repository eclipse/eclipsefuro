package commandpipe

// use this to easily call external programms and pass data to the stdin of them
//
// usage without stdin:
//    cmdLs := commandpipe.NewCommand("ls","-al")
//    content,err := cmdLs.Execute()
//
// usage with stdin:
//    cmdSimpleGenerator := commandpipe.NewCommand("simple-generator","-t=test.tpl")
//    d, err := cmdSimpleGenerator.WriteToStdin(c)
//
// Keep in mind that a command can only be used once. You have to create it again and again if you want to use it in a loop.
import (
	"bufio"
	"bytes"

	"io"
	"log"
	"os"
	"os/exec"
)

type Command struct {
	Buffer  bytes.Buffer
	process *exec.Cmd
	writer  *bufio.Writer
	stdin   io.WriteCloser
}

func NewCommand(command string, args ...string) *Command {
	c := &Command{
		process: exec.Command(command, args...),
	}
	var err error
	c.stdin, err = c.process.StdinPipe()
	if err != nil {
		log.Fatal(command, err)
	}

	c.process.Stdout = &c.Buffer
	c.process.Stderr = os.Stderr

	c.writer = bufio.NewWriter(c.stdin)
	return c
}

// run the process without using stdin and receive the response
func (c *Command) Execute() ([]byte, error) {

	if err := c.process.Run(); err != nil { //Use start, not run
		return nil, err
	}

	// return the response
	return c.Buffer.Bytes(), nil
}

// write data to stdin of process and receive the response
func (c *Command) WriteToStdin(data []byte) ([]byte, error) {

	if err := c.process.Start(); err != nil { //Use start, not run
		return nil, err
	}

	_, err := c.writer.Write(data)
	if err != nil {
		return nil, err
	}
	c.stdin.Close()
	err = c.process.Wait()
	if err != nil {
		return nil, err
	}
	// return the response
	return c.Buffer.Bytes(), nil
}
