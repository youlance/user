CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "date_of_birth" date NOT NULL DEFAULT(now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_followers" (
    "follower_id" varchar,
    "followee_id" varchar,
    PRIMARY KEY("follower_id", "followee_id")
);