-- name: CreateRotationEvent :one
INSERT INTO "clinical_rotation_event" (
    "academic_year_id", "group_id", "service_id", "start_date", "end_date"
) VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: CreateRotationEvents :copyfrom
INSERT INTO "clinical_rotation_event" (
    "academic_year_id", "group_id", "service_id", "start_date", "end_date"
) VALUES($1, $2, $3, $4, $5);

-- name: ListRotationEventsByAcademicYearID :many
SELECT * FROM "clinical_rotation_event"
WHERE "academic_year_id" = $1
ORDER BY "start_date";