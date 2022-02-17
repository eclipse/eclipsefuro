/*
Copyright © 2020

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/eclipse/eclipsefuro/furo/internal/cmd/runner"
	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
	"log"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string

// needed for the documentation generator
var RootCmd = rootCmd

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "furo",
	Short: "The furo spec toolkit",
	Long: `Furo furo contains helpful generators, converters, sanitizer for the furo specs.
Read more about the single commands in the see also section below.

Calling furo without any arguments and flags will run the flow runner with the default flow. 
Modify your default flow in the .furo config file to your needs. You can set any of the sub commands as default.

> Note: Environment variables are prefixed with **FST**. 
>
> To set the specformat with the environment variable use **FST_SPECFORMAT=value**
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		runner.Run(cmd, args)
	},
	Version: "1.34.2",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is CWD/.furo)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in home directory with name ".furo" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
		viper.SetConfigName(".furo")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("FST")
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	// read the furo version, abort if the project version is higher
	projectVersion := ("v" + viper.GetString("furo"))

	toolVersion := ("v" + rootCmd.Version)
	if semver.Max(projectVersion, toolVersion) != toolVersion {
		log.Fatal("The project requires a newer version of furo. \n Furo ", toolVersion, " is installed, ", projectVersion, " is required")
	}

}
