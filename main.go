package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	CallClear()

	config := Config{
		location: "/Users/ebaukhages/Documents/scripts/tmux.sessions.log",
	}
	config.Parse()

	prompt := promptui.Select{
		Label: "Select Session",
		Items: config.keys,
	}

	_, name, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	path := config.values[name]

	session := Session{
		path:    path,
		session: name,
		config:  config,
	}

	_, err = session.Start()

	if err != nil {
		fmt.Printf("Session failed %v\n", err)
		return
	}
}

func init() {
	CallClearInit()
}
