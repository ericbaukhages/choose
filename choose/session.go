package choose

import (
	"errors"
)

// Session holds all the stuff to start these sessions
type Session struct {
	Config  Config
	Path    string
	Session string
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
			s.Session,
			"-c",
			s.Path,
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
			"-t",
			s.Session,
		},
	}
	attach.exec()

	return "Session started successfully", nil
}

func (s *Session) valid() (bool, error) {

	// check if is valid directory
	err := ValidPath(s.Path)
	if err != nil {
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
