package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pjsmith404/gator/internal/database"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Usage: addfeed <name> <url>")
	}

	name, url := cmd.args[0], cmd.args[1]

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), newFeed)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
