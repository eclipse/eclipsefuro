package suggester

import "github.com/c-bata/go-prompt"

func Bool(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "true", Description: "Yes"},
		{Text: "false", Description: "No"},
	}
}
