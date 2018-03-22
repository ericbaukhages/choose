package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
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
