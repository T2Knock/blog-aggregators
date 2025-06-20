-- name: CreateUser :one
INSERT INTO users (user_id, user_name)
VALUES (
    $1,
    $2
)
RETURNING *;


-- name: GetUser :one
SELECT * FROM users WHERE user_name=$1 LIMIT 1;
