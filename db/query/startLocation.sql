-- startLocation.sql

-- name: CreateStartLocation :one
INSERT INTO "startLocation" ("id", "lat", "long", "address", "description", "type") 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "id", "lat", "long", "address", "description", "type";

-- name: GetStartLocationByID :one
SELECT * FROM "startLocation" WHERE "id" = $1;

-- name: UpdateStartLocation :exec
UPDATE "startLocation"
SET "lat" = $1, "long" = $2, "address" = $3, "description" = $4, "type" = $5
WHERE "id" = $6;

-- name: DeleteStartLocation :exec
DELETE FROM "startLocation" WHERE "id" = $1;
