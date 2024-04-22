-- name: CreateTour :one
INSERT INTO "tour" ("name", "duration", "ratings_average", "max_group_size", "difficulty", "ratings_quantity", "price", "summary", "description", "imagecover", "images", "start_dates", "secret_tour", "start_location_id", "location_id") 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13 , $14 , $15)
RETURNING *;

-- name: GetTourByID :one
SELECT * FROM "tour" WHERE "id" = $1;

-- name: GetAllTours :many
SELECT * FROM "tour" limit $1;

-- name: UpdateTour :exec
UPDATE "tour"
SET "name" = $1, "duration" = $2, "max_group_size" = $3, "difficulty" = $4, "ratings_average" = $5, "ratings_quantity" = $6, "price" = $7, "summary" = $8, "description" = $9, "imagecover" = $10, "images" = $11, "start_dates" = $12, "secret_tour" = $13, "start_location_id" = $14, "location_id" = $15
WHERE "id" = $16
RETURNING *;

-- name: DeleteTour :exec
DELETE FROM "tour" WHERE "id" = $1;
