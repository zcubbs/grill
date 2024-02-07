CREATE TABLE "nodes"
(
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(255) NOT NULL,
  "ip" VARCHAR(255) NOT NULL,
  "cpu" VARCHAR(255) NOT NULL,
  "memory" VARCHAR(255) NOT NULL,
  "disk" VARCHAR(255) NOT NULL,
  "cluster_id" UUID NOT NULL,
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" timestamptz NOT NULL DEFAULT current_timestamp,
  "updated_at" timestamptz NOT NULL DEFAULT current_timestamp
);

ALTER TABLE "nodes" ADD FOREIGN KEY ("cluster_id") REFERENCES "clusters" ("id");

CREATE INDEX idx_nodes_name ON "nodes" ("name");
