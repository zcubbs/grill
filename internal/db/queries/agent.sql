-- name: CreateAgent :one
INSERT INTO agents
(name,"group",token,scopes,active)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAgent :one
SELECT * FROM agents WHERE id = $1;

-- name: GetAgentByToken :one
SELECT * FROM agents WHERE token = $1;

-- name: GetAllAgents :many
SELECT * FROM agents
ORDER BY name;
