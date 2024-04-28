-- name: CreateUser :one
INSERT INTO "user" ("name", "email", "role","photo", "password") 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM "user" WHERE "id" = $1;

-- name: UpdateUser :exec
UPDATE "user"
SET "name" = $2 , "email" = $3, "role" = $4
WHERE "id" = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user" WHERE "id" = $1;

-- name: GetAllUsers :many
SELECT "name", "email", "role","photo","active","created_at" FROM "user"
limit $1;

-- name: GetUserByEmail :one
SELECT "name", "email", "role" , "password" FROM "user" WHERE "email" = $1;