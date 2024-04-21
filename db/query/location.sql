-- name: CreateLocation :one
INSERT INTO "location" ("lat", "long", "address", "description", "day", "type") 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllLocation :many
SELECT * FROM "location";

-- name: GetLocationByID :one
SELECT * FROM "location" WHERE "id" = $1;

-- name: UpdateLocation :exec
UPDATE "location"
SET "lat" = $1, "long" = $2, "address" = $3, "description" = $4, "day" = $5, "type" = $6
WHERE "id" = $7
RETURNING *;

-- name: DeleteLocation :exec
DELETE FROM "location" WHERE "id" = $1;
