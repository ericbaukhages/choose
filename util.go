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
func ParseConfig(logFile string) map[string]string {
	f, err := os.OpenFile(logFile, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("open file error: %v", err)
	}

	defer f.Close()

	config := make(map[string]string)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
		}

		values := strings.Split(line, " ")
		name, path := values[0], strings.TrimSpace(values[1])
		config[name] = path
	}

	return config
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

// CallClearInit function initialzes the CallClear necessary components
func CallClearInit() {
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

// MakeKeys function makes keys from string map
func MakeKeys(config map[string]string) []string {
	keys := make([]string, len(config))

	i := 0
	for k := range config {
		keys[i] = k
		i++
	}

	return keys
}
