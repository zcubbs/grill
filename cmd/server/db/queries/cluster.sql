-- name: CreateCluster :one
INSERT INTO clusters (
  name,
  description
) VALUES (
 $1, $2
) RETURNING *;

-- name: GetCluster :one
SELECT * FROM clusters
WHERE id = $1 LIMIT 1;

-- name: GetClusterByName :one
SELECT * FROM clusters
WHERE name = $1 LIMIT 1;

-- name: GetAllClusters :many
SELECT * FROM clusters;

-- name: UpdateCluster :one
UPDATE clusters
SET
  name = COALESCE(sqlc.narg(name), name),
  description = COALESCE(sqlc.narg(description), description)
WHERE
  id = sqlc.arg(id)
  RETURNING *;
