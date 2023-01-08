-- name: CreateHospital :one
INSERT INTO "hospital" (
    name,
    description,
    address
) VALUES (
    $1 , $2 , $3
) RETURNING *;

-- name: GetHospitalByID :one
SELECT * FROM "hospital"
WHERE "id" = $1 LIMIT 1;

-- name: GetHospitalByName :one
SELECT * FROM "hospital"
WHERE name = $1 LIMIT 1;

-- name: ListHospitalsByName :many
SELECT * FROM "hospital"
ORDER BY "name"
LIMIT $1
OFFSET $2;