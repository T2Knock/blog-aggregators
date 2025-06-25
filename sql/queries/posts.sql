-- name: CreatePost :exec
INSERT INTO posts (
    post_id, title, description, url, feed_id, published_at
) VALUES (
    $1, $2, $3, $4, $5, $6
);

-- name: GetPostForUser :many
SELECT * FROM posts
WHERE feed_id = ANY($1::varchar [])
ORDER BY published_at DESC
LIMIT $2;

-- name: GetPostByURL :one
SELECT post_id FROM posts
WHERE url = $1;
