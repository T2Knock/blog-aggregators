-- name: CreateFeedFollows :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (
        feed_follow_id, feed_id, follower_id
    ) VALUES ($1, $2, $3) RETURNING *
)

SELECT inserted_feed_follow.*
FROM inserted_feed_follow
;

-- name: GetFeedFollowForUser :many
SELECT
    feeds.name AS feed_name,
    users.name AS name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.feed_id
INNER JOIN users ON feed_follows.follower_id = users.user_id
WHERE users.name = $1;
