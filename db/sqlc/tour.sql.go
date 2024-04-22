// Code generated by sqlc. DO NOT EDIT.
// source: tour.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createTour = `-- name: CreateTour :one
INSERT INTO "tour" ("name", "duration", "ratings_average", "max_group_size", "difficulty", "ratings_quantity", "price", "summary", "description", "imagecover", "images", "start_dates", "secret_tour", "start_location_id", "location_id") 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13 , $14 , $15)
RETURNING id, name, duration, created_at, max_group_size, difficulty, ratings_average, ratings_quantity, price, summary, description, imagecover, images, start_dates, secret_tour, start_location_id, location_id
`

type CreateTourParams struct {
	Name            string       `json:"name"`
	Duration        int64        `json:"duration"`
	RatingsAverage  int64        `json:"ratings_average"`
	MaxGroupSize    int64        `json:"max_group_size"`
	Difficulty      string       `json:"difficulty"`
	RatingsQuantity int64        `json:"ratings_quantity"`
	Price           int64        `json:"price"`
	Summary         string       `json:"summary"`
	Description     string       `json:"description"`
	Imagecover      string       `json:"imagecover"`
	Images          string       `json:"images"`
	StartDates      time.Time    `json:"start_dates"`
	SecretTour      sql.NullBool `json:"secret_tour"`
	StartLocationID int64        `json:"start_location_id"`
	LocationID      int64        `json:"location_id"`
}

func (q *Queries) CreateTour(ctx context.Context, arg CreateTourParams) (Tour, error) {
	row := q.db.QueryRowContext(ctx, createTour,
		arg.Name,
		arg.Duration,
		arg.RatingsAverage,
		arg.MaxGroupSize,
		arg.Difficulty,
		arg.RatingsQuantity,
		arg.Price,
		arg.Summary,
		arg.Description,
		arg.Imagecover,
		arg.Images,
		arg.StartDates,
		arg.SecretTour,
		arg.StartLocationID,
		arg.LocationID,
	)
	var i Tour
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Duration,
		&i.CreatedAt,
		&i.MaxGroupSize,
		&i.Difficulty,
		&i.RatingsAverage,
		&i.RatingsQuantity,
		&i.Price,
		&i.Summary,
		&i.Description,
		&i.Imagecover,
		&i.Images,
		&i.StartDates,
		&i.SecretTour,
		&i.StartLocationID,
		&i.LocationID,
	)
	return i, err
}

const deleteTour = `-- name: DeleteTour :exec
DELETE FROM "tour" WHERE "id" = $1
`

func (q *Queries) DeleteTour(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTour, id)
	return err
}

const getAllTours = `-- name: GetAllTours :many
SELECT id, name, duration, created_at, max_group_size, difficulty, ratings_average, ratings_quantity, price, summary, description, imagecover, images, start_dates, secret_tour, start_location_id, location_id FROM "tour" limit $1
`

func (q *Queries) GetAllTours(ctx context.Context, limit int32) ([]Tour, error) {
	rows, err := q.db.QueryContext(ctx, getAllTours, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tour
	for rows.Next() {
		var i Tour
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Duration,
			&i.CreatedAt,
			&i.MaxGroupSize,
			&i.Difficulty,
			&i.RatingsAverage,
			&i.RatingsQuantity,
			&i.Price,
			&i.Summary,
			&i.Description,
			&i.Imagecover,
			&i.Images,
			&i.StartDates,
			&i.SecretTour,
			&i.StartLocationID,
			&i.LocationID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTourByID = `-- name: GetTourByID :one
SELECT id, name, duration, created_at, max_group_size, difficulty, ratings_average, ratings_quantity, price, summary, description, imagecover, images, start_dates, secret_tour, start_location_id, location_id FROM "tour" WHERE "id" = $1
`

func (q *Queries) GetTourByID(ctx context.Context, id int64) (Tour, error) {
	row := q.db.QueryRowContext(ctx, getTourByID, id)
	var i Tour
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Duration,
		&i.CreatedAt,
		&i.MaxGroupSize,
		&i.Difficulty,
		&i.RatingsAverage,
		&i.RatingsQuantity,
		&i.Price,
		&i.Summary,
		&i.Description,
		&i.Imagecover,
		&i.Images,
		&i.StartDates,
		&i.SecretTour,
		&i.StartLocationID,
		&i.LocationID,
	)
	return i, err
}

const updateTour = `-- name: UpdateTour :exec
UPDATE "tour"
SET "name" = $1, "duration" = $2, "max_group_size" = $3, "difficulty" = $4, "ratings_average" = $5, "ratings_quantity" = $6, "price" = $7, "summary" = $8, "description" = $9, "imagecover" = $10, "images" = $11, "start_dates" = $12, "secret_tour" = $13, "start_location_id" = $14, "location_id" = $15
WHERE "id" = $16
RETURNING id, name, duration, created_at, max_group_size, difficulty, ratings_average, ratings_quantity, price, summary, description, imagecover, images, start_dates, secret_tour, start_location_id, location_id
`

type UpdateTourParams struct {
	Name            string       `json:"name"`
	Duration        int64        `json:"duration"`
	MaxGroupSize    int64        `json:"max_group_size"`
	Difficulty      string       `json:"difficulty"`
	RatingsAverage  int64        `json:"ratings_average"`
	RatingsQuantity int64        `json:"ratings_quantity"`
	Price           int64        `json:"price"`
	Summary         string       `json:"summary"`
	Description     string       `json:"description"`
	Imagecover      string       `json:"imagecover"`
	Images          string       `json:"images"`
	StartDates      time.Time    `json:"start_dates"`
	SecretTour      sql.NullBool `json:"secret_tour"`
	StartLocationID int64        `json:"start_location_id"`
	LocationID      int64        `json:"location_id"`
	ID              int64        `json:"id"`
}

func (q *Queries) UpdateTour(ctx context.Context, arg UpdateTourParams) error {
	_, err := q.db.ExecContext(ctx, updateTour,
		arg.Name,
		arg.Duration,
		arg.MaxGroupSize,
		arg.Difficulty,
		arg.RatingsAverage,
		arg.RatingsQuantity,
		arg.Price,
		arg.Summary,
		arg.Description,
		arg.Imagecover,
		arg.Images,
		arg.StartDates,
		arg.SecretTour,
		arg.StartLocationID,
		arg.LocationID,
		arg.ID,
	)
	return err
}
