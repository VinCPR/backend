-- name: CreateUser :one
INSERT INTO "user" (
    username,
    hashed_password,
    role_id
) VALUES (
    $1 , $2 , $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM "user"
ORDER BY username 
LIMIT $1
OFFSET $2;