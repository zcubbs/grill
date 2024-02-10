CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE agents (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(255) UNIQUE NOT NULL,
  "token" TEXT UNIQUE NOT NULL,
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  "version" VARCHAR(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  "last_connection" TIMESTAMPTZ
);
