package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/YounesBouchbouk/natours-api-golang/util"
	"github.com/stretchr/testify/require"
)

func createRandomLocation(t *testing.T) Location {

	arg := CreateLocationParams{
		Lat:         1.432,
		Long:        -0.123,
		Address:     sql.NullString{String: fmt.Sprintf("Address %s", util.RandomString(10)), Valid: true},
		Type:        LocationType(util.RandomLocationType()),
		Description: sql.NullString{String: util.RandomString(20), Valid: true},
		Day:         3,
	}

	location, err := testQueries.CreateLocation(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, location)
	require.Equal(t, arg.Lat, location.Lat)
	require.Equal(t, arg.Long, location.Long)
	require.Equal(t, arg.Address, location.Address)
	require.Equal(t, arg.Type, location.Type)
	require.Equal(t, arg.Description, location.Description)
	require.Equal(t, arg.Day, location.Day)

	return location

}

func TestCreateLocation(t *testing.T) {
	createRandomLocation(t)
}

func TestGetLocation(t *testing.T) {
	location1 := createRandomLocation(t)
	location2, err := testQueries.GetLocationByID(context.Background(), location1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, location1)
	require.Equal(t, location1.ID, location2.ID)
}

func TestDeleteLocation(t *testing.T) {
	location := createRandomLocation(t)

	err := testQueries.DeleteLocation(context.Background(), location.ID)

	require.NoError(t, err)

	location2, err := testQueries.GetLocationByID(context.Background(), location.ID)

	require.Empty(t, location2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestGetListOfLocation(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomLocation(t)
	}

	locations, err := testQueries.GetAllLocation(context.Background(), 5)

	require.NoError(t, err)
	require.Len(t, locations, 5)

	for _, loc := range locations {
		require.NotEmpty(t, loc)
	}

}
