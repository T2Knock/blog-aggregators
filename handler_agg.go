package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/T2Knock/blog-aggregators/internal/database"
	"github.com/oklog/ulid/v2"
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

	for _, post := range rssFeed.Channel.Item {
		publishedAt, err := time.Parse(time.RFC1123, post.PubDate)
		if err != nil {
			publishedAt = time.Time{}
		}

		postID, err := s.db.GetPostByURL(ctx, post.Link)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.Printf("Failed get post by url %v", err)
		}

		if postID != "" {
			continue
		}

		if err = s.db.CreatePost(ctx, database.CreatePostParams{PostID: ulid.Make().String(), Title: sql.NullString{String: post.Title, Valid: post.Title != ""}, Description: sql.NullString{String: post.Description, Valid: post.Description != ""}, Url: post.Link, PublishedAt: sql.NullTime{Time: publishedAt, Valid: !publishedAt.IsZero()}, FeedID: feed.FeedID}); err != nil {
			log.Printf("Failed to save post %q: %v", post.Title, err)
		}
	}

	return nil
}
