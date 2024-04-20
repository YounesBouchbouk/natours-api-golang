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
  "id" bigint PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "role" varchar NOT NULL,
  "photo" varchar,
  "password" varchar NOT NULL,
  "confirmpassword" varchar NOT NULL,
  "isPsswordChanged" bool DEFAULT false,
  "passwordResetToken" varchar,
  "passwordResetExpires" date,
  "active" bool DEFAULT true,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "tour" (
  "id" bigint PRIMARY KEY,
  "name" varchar NOT NULL,
  "duration" bigint NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "maxGroupSize" bigint NOT NULL,
  "difficulty" varchar NOT NULL,
  "ratingsAverage" bigint NOT NULL,
  "ratingsQuantity" bigint DEFAULT 10,
  "price" bigint NOT NULL,
  "summary" varchar NOT NULL,
  "description" varchar NOT NULL,
  "imageCover" varchar,
  "images" varchar,
  "startDates" date NOT NULL,
  "secret_tour" bool DEFAULT false,
  "startlocationId" bigint,
  "locationId" bigint
);

CREATE TABLE "review" (
  "id" integer PRIMARY KEY,
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
  "paid" bigint NOT NULL
);

CREATE TABLE "startLocation" (
  "id" bigint PRIMARY KEY,
  "lat" bigint NOT NULL,
  "long" bigint NOT NULL,
  "address" varchar NOT NULL,
  "description" varchar,
  "type" location_type DEFAULT 'point'
);

CREATE TABLE "location" (
  "id" bigint PRIMARY KEY,
  "lat" bigint NOT NULL,
  "long" bigint NOT NULL,
  "address" varchar NOT NULL,
  "description" varchar,
  "day" bigint,
  "type" location_type DEFAULT 'point'
);

ALTER TABLE "tour" ADD FOREIGN KEY ("startlocationId") REFERENCES "startLocation" ("id");

ALTER TABLE "tour" ADD FOREIGN KEY ("locationId") REFERENCES "location" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("tour") REFERENCES "tour" ("id");

ALTER TABLE "review" ADD FOREIGN KEY ("user") REFERENCES "user" ("id");

ALTER TABLE "booking" ADD FOREIGN KEY ("tour") REFERENCES "tour" ("id");

ALTER TABLE "booking" ADD FOREIGN KEY ("user") REFERENCES "user" ("id");
