package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/YounesBouchbouk/natours-api-golang/util"
	"github.com/stretchr/testify/require"
)

func createRandomStartLocation(t *testing.T) StartLocation {

	arg := CreateStartLocationParams{
		Lat:         1.432,
		Long:        -0.123,
		Type:        LocationType(util.RandomLocationType()),
		Description: sql.NullString{String: util.RandomString(20), Valid: true},
		Address:     util.RandomString(20),
	}

	location, err := testQueries.CreateStartLocation(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, location)
	require.Equal(t, arg.Lat, location.Lat)
	require.Equal(t, arg.Long, location.Long)
	require.Equal(t, arg.Address, location.Address)
	require.Equal(t, arg.Type, location.Type)
	require.Equal(t, arg.Description, location.Description)

	return location

}

func TestCreateStartLocation(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomStartLocation(t)
	}
}

func TestGetStartLocation(t *testing.T) {
	location1 := createRandomStartLocation(t)
	location2, err := testQueries.GetStartLocationByID(context.Background(), location1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, location1)
	require.Equal(t, location1.ID, location2.ID)
}

func TestDeleteStartLocation(t *testing.T) {
	location := createRandomLocation(t)

	err := testQueries.DeleteStartLocation(context.Background(), location.ID)

	require.NoError(t, err)

	location2, err := testQueries.GetStartLocationByID(context.Background(), location.ID)

	require.Empty(t, location2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}
