package choose

import (
	"fmt"
	"os"
)

func run() {
	CallClear()

	config := Config{
		Location: "/Users/ebaukhages/Documents/scripts/tmux.sessions.log",
	}
	config.Parse()

	var (
		name string
		err  error
	)

	if len(os.Args) == 1 {
		ui := Interface{
			Config: config,
		}
		name, err = ui.Run()

		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
		}
	} else {
		name = os.Args[1]
	}

	path := config.Values[name]

	if path == "" {
		fmt.Printf("No session called: %s. Exiting.\n", name)
		return
	}

	session := Session{
		Path:    path,
		Session: name,
		Config:  config,
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
