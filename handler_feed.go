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

	newFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{FeedID: ulid.Make().String(), Name: cmd.Arguments[0], Url: cmd.Arguments[1], CreatedBy: user.UserID})
	if err != nil {
		return fmt.Errorf("failed creating new feed: %w", err)
	}

	if _, err = s.db.CreateFeedFollows(ctx, database.CreateFeedFollowsParams{FeedFollowID: ulid.Make().String(), FeedID: newFeed.FeedID, FollowerID: user.UserID}); err != nil {
		return fmt.Errorf("failed following new feed: %w", err)
	}

	fmt.Println(newFeed)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("failed fetching user: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("name: %s, url: %s, user: %s \n", feed.FeedName, feed.Url, feed.UserName)
	}

	return nil
}
