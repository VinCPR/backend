CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(20) UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE "role" (
  "id" bigserial PRIMARY KEY,
  "role_name" varchar(20) UNIQUE NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE "permission" (
  "id" bigserial PRIMARY KEY,
  "permission_name" varchar(50) UNIQUE NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL
);

CREATE TABLE "role_permission" (
  "role_id" bigint NOT NULL,
  "permission_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL,
  PRIMARY KEY ("role_id", "permission_id")
);

CREATE INDEX ON "user" ("username");

CREATE INDEX ON "role" ("role_name");

CREATE INDEX ON "permission" ("permission_name");

CREATE INDEX ON "role_permission" ("role_id", "permission_id");

ALTER TABLE "user" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");