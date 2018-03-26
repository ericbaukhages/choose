package main

import (
	"github.com/manifoldco/promptui"
)

// Interface manages the UI interface
type Interface struct {
	config Config
	prompt promptui.Select
}

// Run starts the interface
func (i *Interface) Run() (string, error) {
	prompt := promptui.Select{
		Label: "Select Session",
		Items: i.config.keys,
	}
	i.prompt = prompt

	_, name, err := i.prompt.Run()

	if err != nil {
		return "", err
	}
	return name, nil
}
