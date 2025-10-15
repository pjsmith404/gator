package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pjsmith404/gator/internal/database"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Usage: follow <url>")
	}

	feedUrl := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Feed not found: %v", err)
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("User not found: %v", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%v is now following %v\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
