-- name: CreateAttending :one
INSERT INTO "attending" (
    user_id,
    firstName,
    lastName,
    mobile
) VALUES (
    $1 , $2 , $3, $4 
) RETURNING *;

-- name: GetAttendingByUserId :one
SELECT * FROM "attending"
WHERE user_id = $1 LIMIT 1;

-- name: ListAttendingsByName :many
SELECT * FROM "attending"
ORDER BY "firstName", "lastName"
LIMIT $1
OFFSET $2;