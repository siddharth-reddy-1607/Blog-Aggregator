-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name, apiKey)
VALUES ($1, $2, $3, $4, ENCODE(SHA256(random()::text::bytea),'base64'))
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE apiKey = $1;
