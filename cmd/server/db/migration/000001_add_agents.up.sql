CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE agents (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(255) UNIQUE NOT NULL,
  "group" VARCHAR(255) NOT NULL,
  "token" TEXT UNIQUE NOT NULL,
  "scopes" TEXT NOT NULL,
  "active" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  "last_connection" TIMESTAMPTZ
);
