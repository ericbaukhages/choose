package choose

import (
	"fmt"
	"strings"
)

// Session holds all the stuff to start these sessions
type Session struct {
	Path    string
	Session string
}

// Start starts the session
func (s *Session) Start() (string, error) {
	// TODO: convert potential partial path to full path

	err := s.valid()
	if err != nil {
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

func (s *Session) valid() error {
	// TODO: make sure path is valid

	listSessions := Command{
		args: []string{
			"tmux",
			"list-sessions",
			"-F",
			"#S",
		},
	}

	sessions, _ := listSessions.run()
	lines := strings.Split(string(sessions), "\n")

	if StringInSlice(s.Session, lines) {
		return fmt.Errorf("%s is already open", s.Session)
	}

	return nil
}
