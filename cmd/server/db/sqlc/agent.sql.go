// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: agent.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createAgent = `-- name: CreateAgent :one
INSERT INTO agents
(name,token,version,is_active)
VALUES ($1, $2, $3, $4)
RETURNING id, name, token, is_active, version, created_at, updated_at, last_connection
`

type CreateAgentParams struct {
	Name     string `json:"name"`
	Token    string `json:"token"`
	Version  string `json:"version"`
	IsActive bool   `json:"is_active"`
}

func (q *Queries) CreateAgent(ctx context.Context, arg CreateAgentParams) (Agent, error) {
	row := q.db.QueryRow(ctx, createAgent,
		arg.Name,
		arg.Token,
		arg.Version,
		arg.IsActive,
	)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.IsActive,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastConnection,
	)
	return i, err
}

const getAgent = `-- name: GetAgent :one
SELECT id, name, token, is_active, version, created_at, updated_at, last_connection FROM agents WHERE id = $1
`

func (q *Queries) GetAgent(ctx context.Context, id uuid.UUID) (Agent, error) {
	row := q.db.QueryRow(ctx, getAgent, id)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.IsActive,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastConnection,
	)
	return i, err
}

const getAgentByToken = `-- name: GetAgentByToken :one
SELECT id, name, token, is_active, version, created_at, updated_at, last_connection FROM agents WHERE token = $1
`

func (q *Queries) GetAgentByToken(ctx context.Context, token string) (Agent, error) {
	row := q.db.QueryRow(ctx, getAgentByToken, token)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.IsActive,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastConnection,
	)
	return i, err
}

const getAllAgents = `-- name: GetAllAgents :many
SELECT id, name, token, is_active, version, created_at, updated_at, last_connection FROM agents
ORDER BY name
`

func (q *Queries) GetAllAgents(ctx context.Context) ([]Agent, error) {
	rows, err := q.db.Query(ctx, getAllAgents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Agent{}
	for rows.Next() {
		var i Agent
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Token,
			&i.IsActive,
			&i.Version,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LastConnection,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAgentLastConnection = `-- name: UpdateAgentLastConnection :one
UPDATE agents
SET
  version = $2,
  last_connection = current_timestamp,
  updated_at = current_timestamp
WHERE id = $1
RETURNING id, name, token, is_active, version, created_at, updated_at, last_connection
`

type UpdateAgentLastConnectionParams struct {
	ID      uuid.UUID `json:"id"`
	Version string    `json:"version"`
}

func (q *Queries) UpdateAgentLastConnection(ctx context.Context, arg UpdateAgentLastConnectionParams) (Agent, error) {
	row := q.db.QueryRow(ctx, updateAgentLastConnection, arg.ID, arg.Version)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.IsActive,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastConnection,
	)
	return i, err
}
