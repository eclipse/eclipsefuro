package parseargs

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func FromFurocConfig() Arglist {
	configFile, err := ioutil.ReadFile(".furoc")
	if err != nil {
		log.Fatal(err)
	}

	type fc struct {
		Furoc Arglist `yaml:"furoc"`
	}

	config := fc{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	a := config.Furoc

	if len(a.Inputs) == 0 {
		a.Inputs = append(a.Inputs, "./")
	}

	// informative, to show which furoc bin was used
	a.Binary = os.Args[0]

	return a
}
