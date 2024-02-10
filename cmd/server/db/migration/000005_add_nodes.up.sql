CREATE TABLE "nodes"
(
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" VARCHAR(255) NOT NULL,
  "ip" VARCHAR(255) NOT NULL,
  "os" VARCHAR(255) NOT NULL,
  "arch" VARCHAR(255) NOT NULL,
  "mac_address" VARCHAR(255) NOT NULL,
  "cpu" VARCHAR(255) NOT NULL,
  "memory" VARCHAR(255) NOT NULL,
  "disk" VARCHAR(255) NOT NULL,
  "cluster_id" UUID NOT NULL,
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  "agent_id" UUID NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT current_timestamp,
  "updated_at" timestamptz NOT NULL DEFAULT current_timestamp
);

ALTER TABLE "nodes" ADD FOREIGN KEY ("cluster_id") REFERENCES "clusters" ("id");

ALTER TABLE "nodes" ADD FOREIGN KEY ("agent_id") REFERENCES "agents" ("id");
ALTER TABLE "nodes" ADD CONSTRAINT "unique_mac_address" UNIQUE ("mac_address");
ALTER TABLE "nodes" ADD CONSTRAINT "unique_agent_id" UNIQUE ("agent_id");
ALTER TABLE "nodes" ADD CONSTRAINT "unique_name_cluster_id" UNIQUE ("name", "cluster_id");

CREATE INDEX idx_nodes_cluster_id ON "nodes" ("cluster_id");

CREATE INDEX idx_nodes_name ON "nodes" ("name");
