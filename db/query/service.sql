-- name: CreateService :one
INSERT INTO "service" (
   specialty_id,
   hospital_id,
   name,
   description
) VALUES (
   $1 , $2 , $3, $4
) RETURNING *;
 
-- name: GetServiceByName :one
SELECT * FROM "service"
WHERE name = $1 LIMIT 1;
 
-- name: ListServicesBySpecialtyID :many
SELECT * FROM "service"
ORDER BY "specialty_id"
LIMIT $1
OFFSET $2;
 
-- name: ListServicesByHospitalID :many
SELECT * FROM "service"
ORDER BY "hospital_id"
LIMIT $1
OFFSET $2;
 
-- name: ListServicesBySpecialtyIDAndHospitalID :many
SELECT * FROM "service"
ORDER BY "specialty_id","hospital_id"
LIMIT $1
OFFSET $2;
 
