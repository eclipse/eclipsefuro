package suggester

import (
	"github.com/c-bata/go-prompt"
	"github.com/eclipse/eclipsefuro/furops/internal/root/specs"
	"sort"
	"strings"
)

func Servicecompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	for typename, spec := range specs.Specs.Services {
		sp := spec.(map[string]interface{})
		if strings.HasPrefix(typename, d.Text) {
			s = append(s, prompt.Suggest{
				Text:        typename,
				Description: sp["description"].(string),
			})
		}
	}
	for typename, spec := range specs.Specs.InstalledServices {
		sp := spec.(map[string]interface{})
		if strings.HasPrefix(typename, d.Text) {
			s = append(s, prompt.Suggest{
				Text:        typename,
				Description: sp["description"].(string),
			})
		}
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Text < s[j].Text
	})
	return s
}
