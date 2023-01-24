-- name: CreateServiceToAttending :one
INSERT INTO "service_to_attending" (
    service_id,
    attending_id
) VALUES (
    $1 , $2 
) RETURNING *;

-- name: GetServiceToAttendingByAttendingID :many
SELECT * FROM "service_to_attending"
WHERE attending_id = $1;

-- name: GetServiceToAttendingByServiceID :many
SELECT * FROM "service_to_attending"
WHERE service_id = $1;

-- name: ListServicesToAttendingsByServiceID :many
SELECT * FROM "service_to_attending"
ORDER BY "service_id";

-- name: ListServicesToAttendingsByAttendingID :many
SELECT * FROM "service_to_attending"
ORDER BY "attending_id";

-- name: ListServicesToAttendingsByServiceIDAndAttendingID :many
SELECT * FROM "service_to_attending"
ORDER BY "service_id","attending_id";
