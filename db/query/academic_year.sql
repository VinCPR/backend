-- name: CreateAcademicYear :one
INSERT INTO "academic_year" (
    "name", "start_date", "end_date"
) VALUES(
    $1, $2, $3
) RETURNING *;

-- name: GetAcademicYearByID :one
SELECT * FROM "academic_year"
WHERE "id" = $1 LIMIT 1;

-- name: GetAcademicYearByName :one
SELECT * FROM "academic_year"
WHERE "name" = $1 LIMIT 1;

-- name: ListAcademicYearByEndDate :many
SELECT * FROM "academic_year"
ORDER BY "end_date" DESC
LIMIT $1
OFFSET $2;