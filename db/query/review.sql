-- name: CreateReview :one
INSERT INTO "review" ("id", "review", "rating", "tour", "user") 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetReviewByID :one
SELECT * FROM "review" WHERE "id" = $1;

-- name: UpdateReview :exec
UPDATE "review"
SET "review" = $1, "rating" = $2, "tour" = $3, "user" = $4
WHERE "id" = $5
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM "review" WHERE "id" = $1;


-- name: GetAllReviews :many
SELECT * FROM "review";