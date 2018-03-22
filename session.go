package main

import (
	"errors"
	"fmt"
	"os/exec"
)

// Session holds all the stuff to start these sessions
type Session struct {
	log     string
	path    string
	script  string
	session string
}

// Start function starts the session
func (s *Session) Start() (string, error) {
	c := exec.Command(s.script, s.session, s.path, "1")
	out, err := c.Output()

	if err != nil {
		fmt.Printf("Command failed: %s", out)
		fmt.Printf("         error: %s\n", err.Error())
		return "", errors.New("command failed")
	}

	return "Session started successfully", nil
}
