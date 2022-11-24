CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(20) NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE "role" (
  "id" bigserial PRIMARY KEY,
  "role_name" varchar(20) NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE "permission" (
  "id" bigserial PRIMARY KEY,
  "permission_name" varchar(50) NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE "role_permission" (
  "role_id" bigint,
  "permission_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL
);

CREATE INDEX ON "user" ("username");

CREATE INDEX ON "role" ("role_name");

CREATE INDEX ON "permission" ("permission_name");

CREATE INDEX ON "role_permission" ("role_id", "permission_id");

ALTER TABLE "user" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");