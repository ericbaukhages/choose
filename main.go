package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

var (
	session string
)

func main() {
	CallClear()

	configFile := "/Users/ebaukhages/Documents/scripts/tmux.sessions.log"

	config := ParseConfig(configFile)
	keys := MakeKeys(config)

	prompt := promptui.Select{
		Label: "Select Session",
		Items: keys,
	}

	_, name, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	path := config[name]

	session := Session{
		path:    path,
		session: name,
		config:  configFile,
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
