package root

import "github.com/c-bata/go-prompt"

func applyTheme() []prompt.Option {
	opts := []prompt.Option{
		//prompt.OptionInputTextColor(prompt.DarkGray),

		prompt.OptionSuggestionTextColor(prompt.Black),
		prompt.OptionSuggestionBGColor(prompt.LightGray),

		prompt.OptionDescriptionTextColor(prompt.Black),
		prompt.OptionDescriptionBGColor(prompt.LightGray),

		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
		prompt.OptionSelectedSuggestionBGColor(prompt.DarkGray),

		prompt.OptionSelectedDescriptionTextColor(prompt.Black),
		prompt.OptionSelectedDescriptionBGColor(prompt.DarkGray),
	}
	return opts
}
