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
	logFiles := "/Users/ebaukhages/Documents/scripts/tmux.sessions.log"

	CallClear()
	ParseConfig(logFiles)
	flag.Parse()

	if flag.NFlag() == 0 {
		PrintUsage()
		os.Exit(1)
	}

	fmt.Printf("Loading session: %s\n", session)
}

func init() {
	flag.StringVarP(&session, "session", "s", "", "Session")
}
