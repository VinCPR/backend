-- name: CreateGroup :one
INSERT INTO "group" (
    "academic_year_id" ,
    "name"
) VALUES (
    $1 , $2
) RETURNING *;

-- name: GetGroupByID :one
SELECT * FROM "group"
WHERE "id" = $1 LIMIT 1;

-- name: GetGroupByName :one
SELECT * FROM "group"
WHERE "name" = $1 LIMIT 1;

-- name: ListGroupsByName :many
SELECT * FROM "group"
WHERE "academic_year_id" = $1
ORDER BY "name";