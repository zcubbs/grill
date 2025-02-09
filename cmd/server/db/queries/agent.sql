-- name: CreateAgent :one
INSERT INTO agents
(name,token,is_active)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAgent :one
SELECT * FROM agents WHERE id = $1;

-- name: GetAgentByToken :one
SELECT * FROM agents WHERE token = $1;

-- name: GetAllAgents :many
SELECT * FROM agents
ORDER BY name;

-- name: UpdateAgentLastConnection :one
UPDATE agents
SET
  version = COALESCE(sqlc.narg(version), version),
  last_connection = current_timestamp,
  updated_at = current_timestamp
WHERE id = $1
RETURNING *;

