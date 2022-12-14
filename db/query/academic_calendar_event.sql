-- name: CreateEvent :one
INSERT INTO "academic_calendar_event" (
    "academic_year_id", "name", "type", "start_date", "end_date"
) VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: CreateEvents :copyfrom
INSERT INTO "academic_calendar_event" (
    "academic_year_id", "name", "type", "start_date", "end_date"
) VALUES($1, $2, $3, $4, $5);

-- name: ListEventsByAcademicYearID :many
SELECT * FROM "academic_calendar_event"
WHERE "academic_year_id" = $1
ORDER BY "start_date";