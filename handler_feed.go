package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/T2Knock/blog-aggregators/internal/database"
	"github.com/oklog/ulid/v2"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("missing arguments: usage %s <name> <url>", cmd.Name)
	}

	if _, err := url.ParseRequestURI(cmd.Arguments[1]); err != nil {
		return errors.New("invalid URL")
	}

	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed fetching user: %w", err)
	}

	newFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{FeedID: ulid.Make().String(), Name: cmd.Arguments[0], Url: cmd.Arguments[1], UserID: user.UserID})
	if err != nil {
		return fmt.Errorf("failed creating new feed: %w", err)
	}

	fmt.Println(newFeed)

	return nil
}
