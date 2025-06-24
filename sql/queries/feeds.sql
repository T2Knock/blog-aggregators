-- name: CreateFeed :one
INSERT INTO feeds (feed_id, name, url, user_id) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetFeeds :many
SELECT
    feeds.name AS feed_name,
    url,
    users.name AS user_name
FROM feeds INNER JOIN users ON feeds.user_id = users.user_id;
