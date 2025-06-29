// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: feeds.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (feed_id, name, url, created_by) VALUES (
    $1, $2, $3, $4
) RETURNING feed_id, name, url, created_by, created_at, updated_at, last_fetched_at
`

type CreateFeedParams struct {
	FeedID    string
	Name      string
	Url       string
	CreatedBy string
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.FeedID,
		arg.Name,
		arg.Url,
		arg.CreatedBy,
	)
	var i Feed
	err := row.Scan(
		&i.FeedID,
		&i.Name,
		&i.Url,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastFetchedAt,
	)
	return i, err
}

const getFeedByURL = `-- name: GetFeedByURL :one
SELECT
    feed_id,
    name,
    url
FROM feeds
WHERE url = $1
`

type GetFeedByURLRow struct {
	FeedID string
	Name   string
	Url    string
}

func (q *Queries) GetFeedByURL(ctx context.Context, url string) (GetFeedByURLRow, error) {
	row := q.db.QueryRowContext(ctx, getFeedByURL, url)
	var i GetFeedByURLRow
	err := row.Scan(&i.FeedID, &i.Name, &i.Url)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT
    feeds.name AS feed_name,
    url,
    users.name AS user_name
FROM feeds INNER JOIN users ON feeds.created_by = users.user_id
`

type GetFeedsRow struct {
	FeedName string
	Url      string
	UserName string
}

func (q *Queries) GetFeeds(ctx context.Context) ([]GetFeedsRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedsRow
	for rows.Next() {
		var i GetFeedsRow
		if err := rows.Scan(&i.FeedName, &i.Url, &i.UserName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNextFeedToFetch = `-- name: GetNextFeedToFetch :one
SELECT
    feed_id,
    name,
    url,
    last_fetched_at,
    updated_at
FROM feeds
ORDER BY
    last_fetched_at
    ASC NULLS FIRST
`

type GetNextFeedToFetchRow struct {
	FeedID        string
	Name          string
	Url           string
	LastFetchedAt sql.NullTime
	UpdatedAt     time.Time
}

func (q *Queries) GetNextFeedToFetch(ctx context.Context) (GetNextFeedToFetchRow, error) {
	row := q.db.QueryRowContext(ctx, getNextFeedToFetch)
	var i GetNextFeedToFetchRow
	err := row.Scan(
		&i.FeedID,
		&i.Name,
		&i.Url,
		&i.LastFetchedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const markFeedFetched = `-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $2
WHERE feed_id = $3
`

type MarkFeedFetchedParams struct {
	LastFetchedAt sql.NullTime
	UpdatedAt     time.Time
	FeedID        string
}

func (q *Queries) MarkFeedFetched(ctx context.Context, arg MarkFeedFetchedParams) error {
	_, err := q.db.ExecContext(ctx, markFeedFetched, arg.LastFetchedAt, arg.UpdatedAt, arg.FeedID)
	return err
}
