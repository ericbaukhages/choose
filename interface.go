package main

import (
	"strings"

	"github.com/manifoldco/promptui"
)

// Interface manages the UI interface
type Interface struct {
	config Config
	prompt promptui.Select
}

// Item stores data used by prompt
type Item struct {
	Name string
	Path string
}

// Run starts the interface
func (i *Interface) Run() (string, error) {

	templates := &promptui.SelectTemplates{
		Active:   "* {{ .Name | yellow }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "* {{ .Name | red }}",

		Details: "Path: {{ .Path }}",
	}

	items := make([]Item, len(i.config.keys))
	for index, key := range i.config.keys {
		items[index] = Item{
			Name: key,
			Path: i.config.values[key],
		}
	}

	searcher := func(input string, index int) bool {
		item := items[index]
		name := strings.Replace(strings.ToLower(item.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:             "Select Session",
		Items:             items,
		Templates:         templates,
		Size:              4,
		Searcher:          searcher,
		StartInSearchMode: true,
	}
	i.prompt = prompt

	_, name, err := i.prompt.Run()

	if err != nil {
		return "", err
	}
	return name, nil
}
