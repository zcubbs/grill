CREATE TABLE "sessions"
(
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" UUID NOT NULL,
  "refresh_token" TEXT NOT NULL,
  "user_agent" VARCHAR(255) NOT NULL,
  "client_ip" VARCHAR(255) NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT current_timestamp
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
