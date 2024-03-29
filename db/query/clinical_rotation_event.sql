-- name: CreateRotationEvent :one
INSERT INTO "clinical_rotation_event" (
    "academic_year_id", "group_id", "service_id", "start_date", "end_date"
) VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetRotationEventByID :one
SELECT * FROM "clinical_rotation_event"
WHERE id = $1 LIMIT 1;

-- name: CreateRotationEvents :copyfrom
INSERT INTO "clinical_rotation_event" (
    "academic_year_id", "group_id", "service_id", "start_date", "end_date"
) VALUES($1, $2, $3, $4, $5);

-- name: ListRotationEventsByAcademicYearID :many
SELECT * FROM "clinical_rotation_event"
WHERE "academic_year_id" = $1
ORDER BY "start_date";

-- name: ListRotationEventsByAcademicYearIDAndDay :many
SELECT * FROM "clinical_rotation_event"
WHERE "academic_year_id" = $1 AND "start_date" <= @day::date and @day::date <= "end_date"
ORDER BY "id";

-- name: ListRotationEventsByAcademicYearIDAndGroupID :many
SELECT * FROM "clinical_rotation_event"
WHERE "academic_year_id" = $1 AND "group_id" = ANY((@group_ids)::bigint[])
ORDER BY "start_date";

-- name: ListRotationEventsByAcademicYearIDAndServiceID :many
SELECT * FROM "clinical_rotation_event"
WHERE "academic_year_id" = $1 AND "service_id" = ANY((@service_ids)::bigint[])
ORDER BY "start_date";

-- name: DeleteRotationEventsByAcademicYear :exec
DELETE FROM "clinical_rotation_event" WHERE "academic_year_id" = $1;
