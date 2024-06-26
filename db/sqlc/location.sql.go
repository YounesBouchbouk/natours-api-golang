// Code generated by sqlc. DO NOT EDIT.
// source: location.sql

package db

import (
	"context"
	"database/sql"
)

const createLocation = `-- name: CreateLocation :one
INSERT INTO "location" ("lat", "long", "address", "description", "day", "type") 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, lat, long, address, description, day, type
`

type CreateLocationParams struct {
	Lat         float64        `json:"lat"`
	Long        float64        `json:"long"`
	Address     sql.NullString `json:"address"`
	Description sql.NullString `json:"description"`
	Day         int64          `json:"day"`
	Type        LocationType   `json:"type"`
}

func (q *Queries) CreateLocation(ctx context.Context, arg CreateLocationParams) (Location, error) {
	row := q.db.QueryRowContext(ctx, createLocation,
		arg.Lat,
		arg.Long,
		arg.Address,
		arg.Description,
		arg.Day,
		arg.Type,
	)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Lat,
		&i.Long,
		&i.Address,
		&i.Description,
		&i.Day,
		&i.Type,
	)
	return i, err
}

const deleteLocation = `-- name: DeleteLocation :exec
DELETE FROM "location" WHERE "id" = $1
`

func (q *Queries) DeleteLocation(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteLocation, id)
	return err
}

const getAllLocation = `-- name: GetAllLocation :many
SELECT id, lat, long, address, description, day, type FROM "location" 
limit $1
`

func (q *Queries) GetAllLocation(ctx context.Context, limit int32) ([]Location, error) {
	rows, err := q.db.QueryContext(ctx, getAllLocation, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Location{}
	for rows.Next() {
		var i Location
		if err := rows.Scan(
			&i.ID,
			&i.Lat,
			&i.Long,
			&i.Address,
			&i.Description,
			&i.Day,
			&i.Type,
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

const getLocationByID = `-- name: GetLocationByID :one
SELECT id, lat, long, address, description, day, type FROM "location" WHERE "id" = $1
`

func (q *Queries) GetLocationByID(ctx context.Context, id int64) (Location, error) {
	row := q.db.QueryRowContext(ctx, getLocationByID, id)
	var i Location
	err := row.Scan(
		&i.ID,
		&i.Lat,
		&i.Long,
		&i.Address,
		&i.Description,
		&i.Day,
		&i.Type,
	)
	return i, err
}

const updateLocation = `-- name: UpdateLocation :exec
UPDATE "location"
SET "lat" = $1, "long" = $2, "address" = $3, "description" = $4, "day" = $5, "type" = $6
WHERE "id" = $7
RETURNING id, lat, long, address, description, day, type
`

type UpdateLocationParams struct {
	Lat         float64        `json:"lat"`
	Long        float64        `json:"long"`
	Address     sql.NullString `json:"address"`
	Description sql.NullString `json:"description"`
	Day         int64          `json:"day"`
	Type        LocationType   `json:"type"`
	ID          int64          `json:"id"`
}

func (q *Queries) UpdateLocation(ctx context.Context, arg UpdateLocationParams) error {
	_, err := q.db.ExecContext(ctx, updateLocation,
		arg.Lat,
		arg.Long,
		arg.Address,
		arg.Description,
		arg.Day,
		arg.Type,
		arg.ID,
	)
	return err
}
