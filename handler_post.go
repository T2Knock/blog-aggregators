package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/T2Knock/blog-aggregators/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	followFeed, err := s.db.GetFeedFollowForUser(ctx, user.Name)
	if err != nil {
		return fmt.Errorf("failed to fetch user follow feed: %w", err)
	}

	var feedIDs []string
	for _, following := range followFeed {
		feedIDs = append(feedIDs, following.FeedID)
	}

	limit := 2
	if len(cmd.Arguments) > 0 {
		if parsed, err := strconv.Atoi(cmd.Arguments[0]); err == nil {
			limit = parsed
		}
	}

	posts, err := s.db.GetPostForUser(ctx, database.GetPostForUserParams{Column1: feedIDs, Limit: int32(limit)})
	if err != nil {
		return fmt.Errorf("failed to fetch user follow posts: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title.String)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.PostUrl)
		fmt.Println("=====================================")
	}

	return nil
}
