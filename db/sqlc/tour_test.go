package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/YounesBouchbouk/natours-api-golang/util"
	"github.com/stretchr/testify/require"
)

func createRandomTour(t *testing.T) Tour {

	location := createRandomLocation(t)
	startlocation := createRandomStartLocation(t)

	arg := CreateTourParams{
		Name:            util.RandomString(20),
		Duration:        util.RandomInt(1, 100),
		Price:           util.RandomInt(1, 100),
		Summary:         util.RandomString(20),
		Imagecover:      util.RandomString(10),
		Images:          util.RandomString(10),
		Difficulty:      util.RandomDiffeculty(),
		SecretTour:      sql.NullBool{Bool: false, Valid: true},
		LocationID:      location.ID,
		StartLocationID: startlocation.ID,
		Description:     util.RandomString(20),
		StartDates:      time.Now().Add(15 * time.Hour * 24 * 15),
		RatingsQuantity: 3,
		MaxGroupSize:    4,
		RatingsAverage:  3,
	}

	tour, err := testQueries.CreateTour(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, tour)
	require.Equal(t, arg.Name, tour.Name)
	require.Equal(t, arg.Duration, tour.Duration)
	require.Equal(t, arg.Difficulty, tour.Difficulty)
	require.Equal(t, arg.RatingsQuantity, tour.RatingsQuantity)
	require.Equal(t, arg.Price, tour.Price)
	require.Equal(t, arg.Summary, tour.Summary)
	require.Equal(t, arg.Description, tour.Description)
	require.Equal(t, arg.Imagecover, tour.Imagecover)
	require.Equal(t, arg.Images, tour.Images)

	return tour

}

func TestCreateTour(t *testing.T) {
	createRandomTour(t)
}

func TestGetTour(t *testing.T) {
	tour1 := createRandomTour(t)
	tour2, err := testQueries.GetTourByID(context.Background(), tour1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, tour2)
	require.Equal(t, tour1.ID, tour2.ID)
}

func TestDeleteTour(t *testing.T) {
	tour := createRandomTour(t)

	err := testQueries.DeleteTour(context.Background(), tour.ID)

	require.NoError(t, err)

	Tour2, err := testQueries.GetTourByID(context.Background(), tour.ID)

	require.Empty(t, Tour2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestGetListOfTour(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTour(t)
	}

	tours, err := testQueries.GetAllTours(context.Background(), 5)

	require.NoError(t, err)
	require.Len(t, tours, 5)

	for _, loc := range tours {
		require.NotEmpty(t, loc)
	}

}
