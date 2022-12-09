-- name: CreatePeriod :one
INSERT INTO "period" (
  academic_year_id,
  name,
  start_date,
  end_date
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPeriodByID :one
SELECT * FROM "period"
WHERE "id" = $1 LIMIT 1;

-- name: ListPeriodsByStartDate :many
SELECT * FROM "period"
WHERE "academic_year_id" = $1
ORDER BY "start_date";