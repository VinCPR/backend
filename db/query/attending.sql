-- name: CreateAttending :one
INSERT INTO "attending" (
    user_id,
    attending_id,
    first_name ,
    last_name,
    mobile
) VALUES (
    $1 , $2 , $3, $4 , $5
) RETURNING *;

-- name: GetAttendingByID :one
SELECT * FROM "attending"
WHERE "id" = $1 LIMIT 1;

-- name: GetAttendingByUserId :one
SELECT * FROM "attending"
WHERE user_id = $1 LIMIT 1;

-- name: GetAttendingByAttendingId :one
SELECT * FROM "attending"
WHERE attending_id = $1 LIMIT 1;

-- name: ListAttendingsByName :many
SELECT * FROM "attending"
ORDER BY first_name, last_name
LIMIT $1
OFFSET $2;