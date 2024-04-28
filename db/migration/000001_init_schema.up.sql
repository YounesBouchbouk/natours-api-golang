CREATE TYPE "Role" AS ENUM (
  'user',
  'admin',
  'guide',
  'leadguide'
);

CREATE TYPE "Difficulty" AS ENUM (
  'low',
  'medieum',
  'hard',
  'very_hard'
);

CREATE TYPE "location_type" AS ENUM (
  'point',
  'square',
  'circle'
);

CREATE TABLE "user" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "role" varchar NOT NULL,
  "photo" varchar NOT NULL,
  "password" varchar NOT NULL,
  "active" bool DEFAULT true,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "tour" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL ,
  "duration" bigint NOT NULL ,
  "created_at" timestamp DEFAULT (now()),
  "max_group_size" bigint NOT NULL,
  "difficulty" varchar NOT NULL ,
  "ratings_average" bigint NOT NULL ,
  "ratings_quantity" bigint NOT NULL,
  "price" bigint NOT NULL ,
  "summary" varchar NOT NULL ,
  "description" varchar NOT NULL ,
  "imagecover" varchar NOT NULL,
  "images" varchar NOT NULL,
  "start_dates" date NOT NULL ,
  "secret_tour" bool DEFAULT false,
  "start_location_id" bigint NOT NULL,
  "location_id" bigint NOT NULL
);

CREATE TABLE "review" (
  "id" BIGSERIAL PRIMARY KEY,
  "created_at" timestamp DEFAULT (now()),
  "review" varchar NOT NULL,
  "rating" bigint NOT NULL,
  "tour" bigint,
  "user" bigint
);

CREATE TABLE "booking" (
  "tour" bigint,
  "user" bigint,
  "price" bigint NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "paid" bool NOT NULL
);

CREATE TABLE "startLocation" (
  "id" BIGSERIAL PRIMARY KEY,
  "lat" FLOAT NOT NULL,
  "long" FLOAT NOT NULL,
  "address" varchar NOT NULL,
  "description" varchar,
  "type" location_type DEFAULT 'point'
);

CREATE TABLE "location" (
  "id" BIGSERIAL PRIMARY KEY,
  "lat" FLOAT NOT NULL,
  "long" FLOAT NOT NULL,
  "address" varchar,
  "description" varchar,
  "day" bigint NOT NULL,
  "type" location_type DEFAULT 'point'
);

ALTER TABLE "tour" ADD FOREIGN KEY ("start_location_id") REFERENCES "startLocation" ("id");

ALTER TABLE "tour" ADD FOREIGN KEY ("location_id") REFERENCES "location" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("tour") REFERENCES "tour" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("user") REFERENCES "user" ("id");

ALTER TABLE "booking" ADD FOREIGN KEY ("tour") REFERENCES "tour" ("id");

ALTER TABLE "booking" ADD FOREIGN KEY ("user") REFERENCES "user" ("id");
