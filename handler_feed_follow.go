package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/T2Knock/blog-aggregators/internal/database"
	"github.com/oklog/ulid/v2"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("missing arguments on command %s <url>")
	}

	ctx := context.Background()

	feed, err := s.db.GetFeedByURL(ctx, cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("failed to fetch feed by url: %w", err)
	}

	user, err := s.db.GetUser(ctx, s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to fetch current user: %w", err)
	}

	if _, err = s.db.CreateFeedFollows(ctx, database.CreateFeedFollowsParams{
		FeedFollowID: ulid.Make().String(),
		FeedID:       feed.FeedID,
		FollowerID:   user.UserID,
	}); err != nil {
		return fmt.Errorf("failed to follow the feed: %w", err)
	}

	fmt.Printf("User %q just follow %q \n", user.Name, feed.Name)

	return nil
}
