package main

import (
	"errors"
	"os"
	"os/exec"
	"syscall"
)

// Session holds all the stuff to start these sessions
type Session struct {
	config  Config
	path    string
	session string
}

// Start function starts the session
func (s *Session) Start() (string, error) {
	// TODO: convert potential partial path to full path

	isValid, err := s.valid()
	if !isValid {
		return "Session was invalid", err
	}

	// TODO: create session in `s.path`
	_, _ = runCommand([]string{
		"tmux",
		"new-session",
		"-d",
		"-s",
		s.session,
	})

	// TODO: Properly handle errors from session creation
	// if err != nil {
	// 	return "create session failed", err
	// }

	executeCommand([]string{
		"tmux",
		"attach",
	})

	return "Session started successfully", nil
}

func (s *Session) valid() (bool, error) {

	// check if is valid directory
	stat, err := os.Stat(s.path)
	if !(err == nil && stat.IsDir()) {
		return false, errors.New("path is not valid directory")
	}

	// TODO: check if tmux session isn't already open
	// runArgs := []string{
	// 	"tmux",
	// 	"list-sessions",
	// 	"-F",
	// 	"#S",
	// }

	// out, err := runCommand(runArgs)
	// if err != nil {
	// 	return false, errors.New("tmux list-sessions command failed")
	// }

	return true, nil
}

func runCommand(args []string) ([]byte, error) {
	c := exec.Command(args[0], args[1:]...)
	out, err := c.Output()

	return out, err
}

func executeCommand(args []string) (string, error) {
	binary, lookErr := exec.LookPath("tmux")
	if lookErr != nil {
		return "Unable to find tmux", lookErr
	}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		return "Unable to run command", execErr
	}

	return "Successfully attached", nil
}
