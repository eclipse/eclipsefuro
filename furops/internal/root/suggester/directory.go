package suggester

import (
	"github.com/c-bata/go-prompt"
	"io/ioutil"
)

func Directory(d prompt.Document) []prompt.Suggest {
	sug := []prompt.Suggest{}

	files, err := ioutil.ReadDir(d.Text)

	if err != nil {
		return []prompt.Suggest{}
	}

	for _, f := range files {
		if f.IsDir() {
			sug = append(sug, prompt.Suggest{
				Text: f.Name(),
			})
		}
	}
	return sug
}
