-- name: CreateAcademicYear :one
INSERT INTO "academic_year" (
    "name", "start_date", "end_date"
) VALUES(
    $1, $2, $3
) RETURNING *;

-- name: GetAcademicYearByName :one
SELECT * FROM "academic_year"
WHERE "name" = $1 LIMIT 1;

-- name: ListUsersByName :many
SELECT * FROM "user"
ORDER BY "name"
LIMIT $1
OFFSET $2;
