package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username provided")
	}

	user := cmd.args[0]

	s.config.SetUser(user)

	fmt.Println("User set to", user)

	return nil
}

