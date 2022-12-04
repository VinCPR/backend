-- name: CreateAcademicCalendarEvent :one
INSERT INTO "academic_calendar_event" (
    "academic_year_id", "name", "type", "start_date", "end_date"
) VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: CreateAcademicCalendarEvents :copyfrom
INSERT INTO "academic_calendar_event" (
    "academic_year_id", "name", "type", "start_date", "end_date"
) VALUES($1, $2, $3, $4, $5);

