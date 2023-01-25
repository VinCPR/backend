-- name: CreateUser :one
INSERT INTO "user" (
    email,
    hashed_password,
    role_name
) VALUES (
    $1 , $2 , $3
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM "user"
WHERE "id" = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM "user"
WHERE email = $1 LIMIT 1;

-- name: ListUsersByID :many
SELECT * FROM "user"
ORDER BY "id"
LIMIT $1
OFFSET $2;