-- name: CreateTour :one
INSERT INTO "tour" ( "name", "duration", "maxGroupSize", "difficulty", "ratingsAverage", "ratingsQuantity", "price", "summary", "description", "imageCover", "images", "startDates", "secret_tour", "startlocationId", "locationId") 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING "id", "name", "duration", "maxGroupSize", "difficulty", "ratingsAverage", "ratingsQuantity", "price", "summary", "description", "imageCover", "images", "startDates", "secret_tour", "startlocationId", "locationId";

-- name: GetTourByID :one
SELECT * FROM "tour" WHERE "id" = $1;

-- name: GetAllTours :one
SELECT * FROM "tour";

-- name: UpdateTour :exec
UPDATE "tour"
SET "name" = $1, "duration" = $2, "maxGroupSize" = $3, "difficulty" = $4, "ratingsAverage" = $5, "ratingsQuantity" = $6, "price" = $7, "summary" = $8, "description" = $9, "imageCover" = $10, "images" = $11, "startDates" = $12, "secret_tour" = $13, "startlocationId" = $14, "locationId" = $15
WHERE "id" = $16
RETURNING *;

-- name: DeleteTour :exec
DELETE FROM "tour" WHERE "id" = $1;
