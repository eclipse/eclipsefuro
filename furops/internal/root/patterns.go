package root

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var suggestions = []prompt.Suggest{}

var patterns = map[string]FPS{}

func ResolvePatterns() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(path.Dir(viper.ConfigFileUsed()))

	patternDirs := viper.GetStringSlice("patternDirs")

	for _, patternDir := range patternDirs {

		files, err := ioutil.ReadDir(patternDir)

		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			if f.IsDir() {
				pdir := f.Name()
				// look for fps.yaml
				fpsFile, err := ioutil.ReadFile(path.Join(patternDir, pdir, "/fps.yaml"))
				if err != nil {
					fmt.Printf("Error reading YAML file: %s\n", err)
					os.Exit(1)
					return
				}

				var fps FPS
				err = yaml.Unmarshal(fpsFile, &fps)
				if err != nil {
					fmt.Println(path.Join(patternDir, pdir, "/fps.yaml"))
					fmt.Printf("Error parsing YAML file: %s\n", err)
					os.Exit(1)
				}
				fps.Path = path.Join(patternDir, pdir)
				patterns[pdir] = fps
				// add patterns with descriptions to suggestion list
				suggestions = append(suggestions, prompt.Suggest{
					Text:        pdir,
					Description: fps.Description,
				})
			}
		}
	}
	os.Chdir(cwd)
}

func patterncompleter(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
