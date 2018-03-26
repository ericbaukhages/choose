package main

import (
	"fmt"
)

func main() {
	CallClear()

	config := Config{
		location: "/Users/ebaukhages/Documents/scripts/tmux.sessions.log",
	}
	config.Parse()

	ui := Interface{
		config: config,
	}
	name, err := ui.Run()

	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
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
