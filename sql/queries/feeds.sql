-- name: CreateFeed :one
INSERT INTO feeds (feed_id, name, url, created_by) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetFeeds :many
SELECT
    feeds.name AS feed_name,
    url,
    users.name AS user_name
FROM feeds INNER JOIN users ON feeds.created_by = users.user_id;

-- name: GetFeedByURL :one
SELECT
    feed_id,
    name,
    url
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $2
WHERE feed_id = $3;

-- name: GetNextFeedToFetch :one
SELECT
    feed_id,
    name,
    url,
    last_fetched_at,
    updated_at
FROM feeds
ORDER BY
    last_fetched_at
    ASC NULLS FIRST;
