CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "email" varchar(50) NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role_name" varchar(20) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "student" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "student_id" varchar(20) NOT NULL,
  "first_name" varchar(100) NOT NULL,
  "last_name" varchar(100) NOT NULL,
  "mobile" varchar(20) NOT NULL,
  "biography" varchar NOT NULL,
  "image" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "attending" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "attending_id" varchar(20) NOT NULL,
  "first_name" varchar(50) NOT NULL,
  "last_name" varchar(50) NOT NULL,
  "mobile" varchar(20) NOT NULL,
  "biography" varchar NOT NULL,
  "image" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "specialty" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "hospital" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "description" varchar NOT NULL,
  "address" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "service" (
  "id" bigserial PRIMARY KEY,
  "specialty_id" bigint NOT NULL,
  "hospital_id" bigint NOT NULL,
  "name" varchar(100) NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "service_to_attending" (
  "id" bigserial PRIMARY KEY,
  "service_id" bigint NOT NULL,
  "attending_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "academic_year" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "academic_calendar_event" (
  "id" bigserial PRIMARY KEY,
  "academic_year_id" bigint NOT NULL,
  "name" varchar(100) NOT NULL,
  "type" varchar(100) NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "period" (
  "id" bigserial PRIMARY KEY,
  "academic_year_id" bigint NOT NULL,
  "name" varchar(100) NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "block" (
  "id" bigserial PRIMARY KEY,
  "academic_year_id" bigint NOT NULL,
  "name" varchar(100) NOT NULL,
  "period" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "group" (
  "id" bigserial PRIMARY KEY,
  "academic_year_id" bigint NOT NULL,
  "name" varchar(100) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "group_to_block" (
  "id" bigserial PRIMARY KEY,
  "academic_year_id" bigint NOT NULL,
  "group_id" bigint NOT NULL,
  "block_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "student_to_group" (
  "id" bigserial PRIMARY KEY,
  "academic_year_id" bigint NOT NULL,
  "student_id" bigint NOT NULL,
  "group_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "clinical_rotation_event" (
  "id" bigserial PRIMARY KEY,
  "academic_year_id" bigint NOT NULL,
  "group_id" bigint NOT NULL,
  "service_id" bigint NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX ON "user" ("email");

CREATE UNIQUE INDEX ON "student" ("user_id");

CREATE UNIQUE INDEX ON "student" ("student_id");

CREATE UNIQUE INDEX ON "attending" ("user_id");

CREATE UNIQUE INDEX ON "attending" ("attending_id");

CREATE INDEX ON "specialty" ("name");

CREATE INDEX ON "hospital" ("name");

CREATE INDEX ON "service" ("name");

CREATE UNIQUE INDEX ON "service" ("specialty_id", "hospital_id", "name");

CREATE UNIQUE INDEX ON "service_to_attending" ("service_id", "attending_id");

CREATE INDEX ON "service_to_attending" ("service_id");

CREATE INDEX ON "service_to_attending" ("attending_id");

CREATE UNIQUE INDEX ON "academic_year" ("name");

CREATE UNIQUE INDEX ON "period" ("academic_year_id", "name");

CREATE UNIQUE INDEX ON "block" ("academic_year_id", "period", "name");

CREATE UNIQUE INDEX ON "group" ("academic_year_id", "name");

CREATE UNIQUE INDEX ON "group_to_block" ("group_id", "block_id");

CREATE INDEX ON "group_to_block" ("group_id");

CREATE INDEX ON "group_to_block" ("block_id");

CREATE UNIQUE INDEX ON "student_to_group" ("student_id", "group_id");

CREATE INDEX ON "student_to_group" ("student_id");

CREATE INDEX ON "student_to_group" ("group_id");

ALTER TABLE "student" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "attending" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "service" ADD FOREIGN KEY ("specialty_id") REFERENCES "specialty" ("id");

ALTER TABLE "service" ADD FOREIGN KEY ("hospital_id") REFERENCES "hospital" ("id");

ALTER TABLE "service_to_attending" ADD FOREIGN KEY ("service_id") REFERENCES "service" ("id");

ALTER TABLE "service_to_attending" ADD FOREIGN KEY ("attending_id") REFERENCES "attending" ("id");

ALTER TABLE "academic_calendar_event" ADD FOREIGN KEY ("academic_year_id") REFERENCES "academic_year" ("id");

ALTER TABLE "period" ADD FOREIGN KEY ("academic_year_id") REFERENCES "academic_year" ("id");

ALTER TABLE "block" ADD FOREIGN KEY ("academic_year_id") REFERENCES "academic_year" ("id");

ALTER TABLE "block" ADD FOREIGN KEY ("period") REFERENCES "period" ("id");

ALTER TABLE "group" ADD FOREIGN KEY ("academic_year_id") REFERENCES "academic_year" ("id");

ALTER TABLE "group_to_block" ADD FOREIGN KEY ("academic_year_id") REFERENCES "academic_year" ("id");

ALTER TABLE "group_to_block" ADD FOREIGN KEY ("group_id") REFERENCES "group" ("id");

ALTER TABLE "group_to_block" ADD FOREIGN KEY ("block_id") REFERENCES "block" ("id");

ALTER TABLE "student_to_group" ADD FOREIGN KEY ("academic_year_id") REFERENCES "academic_year" ("id");

ALTER TABLE "student_to_group" ADD FOREIGN KEY ("student_id") REFERENCES "student" ("id");

ALTER TABLE "student_to_group" ADD FOREIGN KEY ("group_id") REFERENCES "group" ("id");

ALTER TABLE "clinical_rotation_event" ADD FOREIGN KEY ("academic_year_id") REFERENCES "academic_year" ("id");

ALTER TABLE "clinical_rotation_event" ADD FOREIGN KEY ("group_id") REFERENCES "group" ("id");

ALTER TABLE "clinical_rotation_event" ADD FOREIGN KEY ("service_id") REFERENCES "service" ("id");
