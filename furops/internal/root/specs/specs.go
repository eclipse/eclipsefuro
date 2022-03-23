package specs

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var Specs struct {
	Services          map[string]interface{} `yaml:"services"`
	Types             map[string]interface{} `yaml:"types"`
	InstalledServices map[string]interface{} `yaml:"installedServices"`
	InstalledTypes    map[string]interface{} `yaml:"installedTypes"`
	Config            map[string]interface{} `yaml:"config"`
	Enums             map[string]interface{} `yaml:"enums"`
}

func ReadSpecs() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(path.Dir(viper.ConfigFileUsed()))

	specFile, err := ioutil.ReadFile(viper.GetString("specExportFile"))
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		os.Exit(1)
		return
	}
	err = yaml.Unmarshal(specFile, &Specs)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		os.Exit(1)
	}
	os.Chdir(cwd)
}
