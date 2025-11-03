package main

import (
	"context"
	"fmt"
	"github.com/pjsmith404/gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		newLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = newLimit
	}

	feedsFollowing, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feedsFollowing {
		posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
			FeedID: feed.FeedID,
			Limit:  int32(limit),
		})
		if err != nil {
			return err
		}

		for _, post := range posts {
			fmt.Println(post.Title, post.Url, post.PublishedAt)
		}
	}

	return nil
}
