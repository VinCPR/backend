-- name: CreateGroupToBlock :one
INSERT INTO "group_to_block" (
    academic_year_id,
    group_id,
    block_id
) VALUES (
    $1 , $2, $3
) RETURNING *;

-- name: GetGroupToBlockByAcademicYearID :many
SELECT * FROM "group_to_block"
WHERE academic_year_id = $1;

-- name: GetGroupToBlockByBlockID :many
SELECT * FROM "group_to_block"
WHERE block_id = $1;

-- name: GetGroupToBlockByGroupID :many
SELECT * FROM "group_to_block"
WHERE group_id = $1;