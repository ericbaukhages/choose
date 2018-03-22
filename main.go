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

	logFile := "/Users/ebaukhages/Documents/scripts/tmux.sessions.log"
	sessionScript := "/Users/ebaukhages/Documents/scripts/session.sh"

	config := ParseConfig(logFile)
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
		log:     sessionScript,
		script:  logFile,
	}

	session.Start()
}

func init() {
	CallClearInit()
}
