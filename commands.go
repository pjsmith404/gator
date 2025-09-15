package main

import (
	"fmt"
	"github.com/pjsmith404/gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	command := c.commands[cmd.name]

	if command == nil {
		return fmt.Errorf("Command not found: %v", cmd.name)
	}

	err := command(s, cmd)
	if err != nil {
		return fmt.Errorf("Running command %v: %v", cmd.name, err)
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

