package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

var (
	session string
)

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		printUsage()
		os.Exit(1)
	}

	fmt.Printf("Loading session: %s\n", session)
}

func init() {
	flag.StringVarP(&session, "session", "s", "", "Session")
}
