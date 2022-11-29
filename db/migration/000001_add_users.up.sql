CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(20) UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user" ("username");
