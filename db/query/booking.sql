-- name: CreateBooking :one
INSERT INTO "booking" ("tour", "user", "price", "paid") 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetBookingByTourAndUser :many
SELECT * FROM "booking" 
WHERE "tour" = $1 AND "user" = $2
ORDER BY created_at;

-- name: UpdateBooking :exec
UPDATE "booking"
SET "price" = $1, "paid" = $2
WHERE "tour" = $3 AND "user" = $4
RETURNING *;

-- name: DeleteBooking :exec
DELETE FROM "booking" WHERE "tour" = $1 AND "user" = $2;

-- name: GetAllBooking :many
SELECT * FROM "booking";