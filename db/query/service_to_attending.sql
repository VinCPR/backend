-- name: CreateServiceToAttending :one
INSERT INTO "service_to_attending" (
    service_id,
    attending_id
) VALUES (
    $1 , $2 
) RETURNING *;

<<<<<<< HEAD
-- TODO: change list by service id -> many
-- attending_id -> many

=======
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
-- name: GetServiceToAttendingByServiceID :one
SELECT * FROM "service_to_attending"
WHERE service_id = $1 LIMIT 1;

-- name: GetServiceToAttendingByAttendingID :one
SELECT * FROM "service_to_attending"
WHERE attending_id = $1 LIMIT 1;

-- name: ListServicesToAttendingsByServiceID :many
SELECT * FROM "service_to_attending"
<<<<<<< HEAD
ORDER BY "service_id";

-- name: ListServicesToAttendingsByAttendingID :many
SELECT * FROM "service_to_attending"
ORDER BY "attending_id";

-- name: ListServicesToAttendingsByServiceIDAndAttendingID :many
SELECT * FROM "service_to_attending"
ORDER BY "service_id","attending_id";
=======
ORDER BY "service_id"
LIMIT $1
OFFSET $2;

-- name: ListServicesToAttendingsByAttendingID :many
SELECT * FROM "service_to_attending"
ORDER BY "attending_id"
LIMIT $1
OFFSET $2;

-- name: ListServicesToAttendingsByAll :many
SELECT * FROM "service_to_attending"
ORDER BY "service_id","attending_id"
LIMIT $1
OFFSET $2;
>>>>>>> cfc0062 (add sql and test for hospital, specialty, service and service to attending)
