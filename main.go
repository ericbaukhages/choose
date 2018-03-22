package main

import (
	"fmt"
	"os/exec"

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

	_, session, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	path := config[session]
	// fmt.Printf("You choose %s and will go to %s\n", session, path)

	c := exec.Command(sessionScript, session, path, "1")
	out, err := c.Output()

	if err != nil {
		fmt.Printf("Command failed: %s", out)
		fmt.Printf("         error: %s\n", err.Error())
		return
	}
}

func init() {
	CallClearInit()
}
