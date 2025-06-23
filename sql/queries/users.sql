-- name: CreateUser :one
INSERT INTO users (user_id, name) VALUES ($1, $2) RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;
