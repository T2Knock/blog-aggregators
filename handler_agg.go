package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/T2Knock/blog-aggregators/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("missing arguments on command %s <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("failed parsing duration: %w", err)
	}

	log.Printf("Collecting feeds every %q \n", cmd.Arguments[0])

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch next feed: %w", err)
	}

	log.Printf("Fetching %q at %q\n", feed.Name, feed.Url)

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	if err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{FeedID: feed.FeedID, LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}, UpdatedAt: time.Now()}); err != nil {
		return fmt.Errorf("failed to mark feed fetched: %w", err)
	}

	log.Println("Feed Items")

	for _, item := range rssFeed.Channel.Item {
		log.Printf("* %s\n", item.Title)
	}

	return nil
}
