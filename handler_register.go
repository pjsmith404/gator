package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pjsmith404/gator/internal/database"
	"time"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username provided")
	}

	user := cmd.args[0]
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	createdUser, err := s.db.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}

	err = s.config.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Println("User registered:", createdUser)

	return nil
}
