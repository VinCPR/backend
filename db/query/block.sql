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

-- name: ListBlocksByAcademicYear :many
SELECT * FROM "block"
WHERE "academic_year_id" = $1
ORDER BY "period";

