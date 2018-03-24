package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Command provides useful wrapper for system calls
type Command struct {
	args []string
	dir  string
	env  map[string]string
}

func (c *Command) run() ([]byte, error) {
	command := create(c.args, c.dir)

	if c.env != nil {
		for k, v := range c.env {
			addEnv(command, k, v)
		}
	}

	out, err := command.Output()

	return out, err
}

func (c *Command) exec() (string, error) {
	binary, lookErr := exec.LookPath(c.args[0])
	if lookErr != nil {
		return fmt.Sprintf("Unable to find %s", c.args[0]), lookErr
	}

	env := os.Environ()

	execErr := syscall.Exec(binary, c.args, env)
	if execErr != nil {
		return "Unable to run command", execErr
	}

	return "Successfully attached", nil
}

func addEnv(cmd *exec.Cmd, key string, value string) {
	env := os.Environ()
	env = append(env, fmt.Sprintf("%s=%s", key, value))
	cmd.Env = env
}

func create(args []string, dir string) *exec.Cmd {
	var command *exec.Cmd

	if len(args) > 1 {
		command = exec.Command(args[0], args[1:]...)
	} else {
		command = exec.Command(args[0])
	}

	if dir != "" {
		command.Dir = dir
	}

	return command
}
