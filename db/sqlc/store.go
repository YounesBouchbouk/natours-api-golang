package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store struct {
	*Queries
	db *sql.DB
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TourTransactionResult struct {
	Name            string        `json:"name"`
	Duration        int64         `json:"duration"`
	RatingsAverage  int64         `json:"ratings_average"`
	MaxGroupSize    int64         `json:"max_group_size"`
	Difficulty      string        `json:"difficulty"`
	RatingsQuantity int64         `json:"ratings_quantity"`
	Price           int64         `json:"price"`
	Summary         string        `json:"summary"`
	Description     string        `json:"description"`
	ImageCover      string        `json:"imagecover"`
	Images          string        `json:"images"`
	StartDates      time.Time     `json:"start-dates"`
	SecretTour      bool          `json:"secret_tour"`
	StartLocation   StartLocation `json:"start_location"`
	Location        Location      `json:"location"`
}

type StartlocationTrParams struct {
	Lat         float64        `json:"lat"`
	Long        float64        `json:"long"`
	Address     string         `json:"address"`
	Description sql.NullString `json:"description"`
	Type        LocationType   `json:"type"`
}

type LocationTrParams struct {
	Lat         float64        `json:"lat"`
	Long        float64        `json:"long"`
	Address     sql.NullString `json:"address"`
	Description sql.NullString `json:"description"`
	Day         int64          `json:"day"`
	Type        LocationType   `json:"type"`
}

type TourTrParams struct {
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

// type CreateTourTransactionParams struct {
// 	tour              tourParams          `json:"tour"`
// 	location          locationParams      `json:"location"`
// 	startLocationArgs startlocationParams `json:"startlocation"`
// }

func (store *Store) CreateNewTourTransaction(
	ctx context.Context,
	locationArgs LocationTrParams,
	startLocationArgs StartlocationTrParams,
	tourArgs TourTrParams,
) (TourTransactionResult, error) {

	var tour_result TourTransactionResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		start_Location, err := q.CreateStartLocation(ctx, CreateStartLocationParams(startLocationArgs))

		if err != nil {
			return err
		}

		location, err := q.CreateLocation(ctx, CreateLocationParams(locationArgs))

		if err != nil {
			return err
		}

		tourArgs.StartLocationID = start_Location.ID
		tourArgs.LocationID = location.ID

		created_tour, err := q.CreateTour(ctx, CreateTourParams(tourArgs))

		if err != nil {
			return err
		}

		tour_result = TourTransactionResult{
			Name:            created_tour.Name,
			Duration:        created_tour.Duration,
			RatingsAverage:  created_tour.RatingsAverage,
			Price:           created_tour.Price,
			Description:     created_tour.Description,
			Difficulty:      created_tour.Difficulty,
			MaxGroupSize:    created_tour.MaxGroupSize,
			Location:        location,
			StartDates:      created_tour.StartDates,
			ImageCover:      created_tour.Imagecover,
			RatingsQuantity: created_tour.RatingsQuantity,
			Summary:         created_tour.Summary,
			Images:          created_tour.Images,
			SecretTour:      created_tour.SecretTour.Bool,
			StartLocation:   start_Location,
		}

		return nil

	})

	return tour_result, err
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
