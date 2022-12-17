-- name: CreateSpecialty :one
INSERT INTO "specialty" (
    name,
    description
) VALUES (
    $1 , $2 
) RETURNING *;

-- name: GetSpecialtyByName :one
SELECT * FROM "specialty"
WHERE name = $1 LIMIT 1;

-- name: ListSpecialtiesByName :many
SELECT * FROM "specialty"
ORDER BY name 
LIMIT $1
OFFSET $2;