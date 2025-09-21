package main

import (
	"fmt"
	"github.com/pjsmith404/gator/internal/config"
	"github.com/pjsmith404/gator/internal/database"
)

type state struct {
	db     *database.Queries
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
	command, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("Command not found: %v", cmd.name)
	}

	return command(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
