-- name: CreateServiceToAttending :one
INSERT INTO "service_to_attending" (
    service_id,
    attending_id
) VALUES (
    $1 , $2 
) RETURNING *;

-- name: ListServicesToAttendingsByServiceID :many
SELECT * FROM "service_to_attending"
ORDER BY "service_id";

-- name: ListServicesToAttendingsByAttendingID :many
SELECT * FROM "service_to_attending"
ORDER BY "attending_id";

-- name: ListServicesToAttendingsByServiceIDAndAttendingID :many
SELECT * FROM "service_to_attending"
ORDER BY "service_id","attending_id";
