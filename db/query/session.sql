-- name: CreateSession :one
INSERT INTO "session" (
  id,
  user_email,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSessionByID :one
SELECT * FROM "session"
WHERE id = $1 LIMIT 1;
