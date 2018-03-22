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

	keys := MakeKeys(config)
	fmt.Printf("%s", keys)

}

func init() {
	CallClearInit()
}
