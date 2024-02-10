-- name: CreateNode :one
INSERT INTO nodes (
  name,
  cluster_id,
  is_active
) VALUES (
 $1, $2, $3
) RETURNING *;

-- name: GetNode :one
SELECT * FROM nodes
WHERE id = $1 LIMIT 1;

-- name: GetNodeByName :one
SELECT * FROM nodes
WHERE name = $1 LIMIT 1;

-- name: GetAllNodes :many
SELECT * FROM nodes;

-- name: GetAllNodesActive :many
SELECT * FROM nodes
WHERE is_active = true;

-- name: GetNodesByCluster :many
SELECT * FROM nodes
WHERE cluster_id = $1;

-- name: GetNodesByClusterActive :many
SELECT * FROM nodes
WHERE cluster_id = $1
AND is_active = true;

-- name: UpdateNode :one
UPDATE nodes
SET
  name = COALESCE(sqlc.narg(name), name),
  ip = COALESCE(sqlc.narg(ip), ip),
  os = COALESCE(sqlc.narg(os), os),
  arch = COALESCE(sqlc.narg(arch), arch),
  mac_address = COALESCE(sqlc.narg(mac_address), mac_address),
  cpu = COALESCE(sqlc.narg(cpu), cpu),
  memory = COALESCE(sqlc.narg(memory), memory),
  disk = COALESCE(sqlc.narg(disk), disk),
  agent_id = COALESCE(sqlc.narg(agent_id), agent_id),
  cluster_id = COALESCE(sqlc.narg(cluster_id), cluster_id),
  is_active = COALESCE(sqlc.narg(is_active), is_active),
  updated_at = current_timestamp
WHERE
  id = sqlc.arg(id)
  RETURNING *;
