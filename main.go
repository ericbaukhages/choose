package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var (
	session string
)

func main() {
	CallClear()

	logFiles := "/Users/ebaukhages/Documents/scripts/tmux.sessions.log"

	config := ParseConfig(logFiles)
	fmt.Printf("%s", config)
	flag.Parse()

	if flag.NFlag() == 0 {
		PrintUsage()
		os.Exit(1)
	}

	fmt.Printf("Loading session: %s\n", session)
}

func init() {
	CallClearInit()
	flag.StringVarP(&session, "session", "s", "", "Session")
}
