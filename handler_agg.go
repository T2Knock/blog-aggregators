package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/T2Knock/blog-aggregators/internal/database"
)

func scrapeFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch next feed: %w", err)
	}

	fmt.Printf("Fetching %q at %q\n", feed.Name, feed.Url)

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	if err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{FeedID: feed.FeedID, LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}, UpdatedAt: time.Now()}); err != nil {
		return fmt.Errorf("failed to mark feed fetched: %w", err)
	}

	fmt.Println("Feed Items")

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}

	return nil
}
