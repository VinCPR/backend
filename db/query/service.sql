-- name: CreateService :one
INSERT INTO "service" (
   specialty_id,
   hospital_id,
   name,
   description
) VALUES (
   $1 , $2 , $3, $4
) RETURNING *;

-- name: GetServiceByID :one
SELECT * FROM "service"
WHERE "id" = $1 LIMIT 1;

-- name: GetServiceByName :one
SELECT * FROM "service"
WHERE name = $1 LIMIT 1;

 -- name: GetServiceByIndex :one
 SELECT * FROM "service"
 WHERE specialty_id = $1 AND hospital_id = $2 AND name = $3 LIMIT 1;

 -- name: GetServiceByHospitalAndSpecialty :many
 SELECT * FROM "service"
 WHERE specialty_id = $1 AND hospital_id = $2;

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
 
-- name: ListAllServicesBySpecialtyIDAndHospitalID :many
SELECT * FROM "service"
ORDER BY "specialty_id","hospital_id";