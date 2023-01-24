-- name: CreateBlock :one
INSERT INTO "block" (
  academic_year_id,
  name,
  period
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetBlockByID :one
SELECT * FROM "block"
WHERE "id" = $1 LIMIT 1;

-- name: GetBlockByIndex :one
SELECT * FROM "block"
WHERE "academic_year_id" = $1 AND "period" = $2 AND "name" = $3
LIMIT 1;

-- name: ListBlocksByAcademicYear :many
SELECT * FROM "block"
WHERE "academic_year_id" = $1
ORDER BY "period";

-- name: DeleteBlocksByAcademicYear :exec
DELETE FROM "block" WHERE "academic_year_id" = $1;

