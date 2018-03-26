package main

import (
	"fmt"
	"os"
)

func main() {
	CallClear()

	config := Config{
		location: "/Users/ebaukhages/Documents/scripts/tmux.sessions.log",
	}
	config.Parse()

	var (
		name string
		err  error
	)

	if len(os.Args) == 1 {
		ui := Interface{
			config: config,
		}
		name, err = ui.Run()

		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
		}
	} else {
		name = os.Args[1]
	}

	path := config.values[name]

	if path == "" {
		fmt.Printf("No session called: %s. Exiting.\n", name)
		return
	}

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
