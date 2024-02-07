-- name: CreateCluster :one
INSERT INTO clusters (
  name,
  description,
  is_active
) VALUES (
 $1, $2, $3
) RETURNING *;

-- name: GetCluster :one
SELECT * FROM clusters
WHERE id = $1 LIMIT 1;

-- name: GetClusterByName :one
SELECT * FROM clusters
WHERE name = $1 LIMIT 1;

-- name: GetAllClusters :many
SELECT * FROM clusters;

-- name: GetAllClustersActive :many
SELECT * FROM clusters
WHERE is_active = true;

-- name: UpdateCluster :one
UPDATE clusters
SET
  name = COALESCE(sqlc.narg(name), name),
  description = COALESCE(sqlc.narg(description), description),
  is_active = COALESCE(sqlc.narg(is_active), is_active)
WHERE
  id = sqlc.arg(id)
  RETURNING *;
