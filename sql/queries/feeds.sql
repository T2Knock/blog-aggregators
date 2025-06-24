-- name: CreateFeed :one
INSERT INTO feeds (feed_id, name, url, user_id) VALUES ($1, $2, $3, $4) RETURNING *;
