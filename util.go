package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
	clear map[string]func()
)

// ParseConfig function reads lines from a log file
func ParseConfig(logFile string) {
	f, err := os.OpenFile(logFile, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("open file error: %v", err)
	}

	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}

		values := strings.Split(line, " ")
		sessionName, path := values[0], values[1]
		fmt.Printf("sessionName: %s, path: %s", sessionName, path)
	}
}

// CallClear function clears the terminal screen
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported, I can't clear the screen!")
	}
}

// initialzes the CallClear necessary components
func init() {
	// set up the clear func
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = clear["linux"]
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// PrintUsage function prints command usage
func PrintUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
}
