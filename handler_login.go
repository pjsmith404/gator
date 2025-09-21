package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username provided")
	}

	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}

	err = s.config.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Println("User set to", name)

	return nil
}
