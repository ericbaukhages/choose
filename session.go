package main

import (
	"errors"
	"os"
)

// Session holds all the stuff to start these sessions
type Session struct {
	config  Config
	path    string
	session string
}

// Start starts the session
func (s *Session) Start() (string, error) {
	// TODO: convert potential partial path to full path

	isValid, err := s.valid()
	if !isValid {
		return "Session was invalid", err
	}

	create := Command{
		args: []string{
			"tmux",
			"new-session",
			"-d",
			"-s",
			s.session,
		},
		env: map[string]string{
			"ORIG_PWD_FOR_TMUX": s.path,
		},
	}
	_, _ = create.run()

	// TODO: Properly handle errors from session creation
	// if err != nil {
	// 	return "create session failed", err
	// }

	attach := Command{
		args: []string{
			"tmux",
			"attach",
		},
	}
	attach.exec()

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
