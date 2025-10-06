package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pjsmith404/gator/internal/database"
	"time"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return err
	}

	currentUser := s.config.CurrentUserName

	for _, user := range users {
		if user.Name == currentUser {
			fmt.Println("*", user.Name, "(current)")
		} else {
			fmt.Println("*", user.Name)
		}
	}

	return nil
}

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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username provided")
	}

	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("User not regstered: %v", err)
	}

	err = s.config.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Println("User set to", name)

	return nil
}
