package runner

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Run(cmd *cobra.Command, args []string) {
	flow := "default"

	f := cmd.Flag("flow")
	if f != nil {
		// override with flag
		flow = f.Value.String()
	}

	// use argument as flow if set
	if len(args) == 1 && flow == "default" {
		flow = args[0]
	}

	listOfConfiguredFlows := collectCommands(cmd)

	fmt.Println("START FLOW: " + flow)
	seq := viper.GetStringSlice("flows." + flow)
	// flow not found
	if len(seq) == 0 {
		log.Fatal("flow is not defined or has no sequence --flow=", flow)
	}
	fmt.Println("sequence", seq)
	for _, step := range seq {
		fmt.Println("RUNNING STEP: ", step)
		if listOfConfiguredFlows[step] != nil {
			// configured flows go first
			listOfConfiguredFlows[step](cmd, args)
		} else {
			// try commands
			commandList := viper.GetStringMapString("commands")
			if commandList[strings.ToLower(step)] != "" {
				// exec the command...
				cmd := exec.Command(commandList[strings.ToLower(step)], args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatalf("cmd.Run() failed with %s\n", err)
				}
			}
		}

	}

}

// commands are in parent
func collectCommands(cmd *cobra.Command) (commandList map[string]func(cmd *cobra.Command, args []string)) {
	if cmd.Commands() == nil {
		cmd = cmd.Parent()
	}
	commandList = map[string]func(cmd *cobra.Command, args []string){}
	for _, c := range cmd.Commands() {
		commandList[c.Use] = c.Run
	}
	return commandList
}
