-- name: CreateFeedFollows :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (
        feed_follow_id, feed_id, follower_id
    ) VALUES ($1, $2, $3) RETURNING *
)

SELECT inserted_feed_follow.*
FROM inserted_feed_follow
;
