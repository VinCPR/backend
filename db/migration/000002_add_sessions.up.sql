CREATE TABLE "session" (
  "id" uuid PRIMARY KEY,
  "user_email" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "session" ADD FOREIGN KEY ("user_email") REFERENCES "user" ("email");
