package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomBooking(t *testing.T) Booking {
	tour := createRandomTour(t)
	user := createRandomUser(t)

	args := CreateBookingParams{
		Tour:  sql.NullInt64{Int64: tour.ID, Valid: true},
		User:  sql.NullInt64{Int64: user.ID, Valid: true},
		Price: 33,
		Paid:  false,
	}

	booking, err := testQueries.CreateBooking(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, booking)

	return booking
}
func TestCreateBooking(t *testing.T) {
	createRandomBooking(t)
}

func TestGetBooking(t *testing.T) {
	newBooking := createRandomBooking(t)
	args := GetBookingByTourAndUserParams{
		Tour: newBooking.Tour,
		User: newBooking.User,
	}

	booking, err := testQueries.GetBookingByTourAndUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, booking)

}
func TestDeleteBookingt(t *testing.T) {
	// createRandomBooking(t)
	newBooking := createRandomBooking(t)

	args := DeleteBookingParams{
		Tour: newBooking.Tour,
		User: newBooking.User,
	}

	err := testQueries.DeleteBooking(context.Background(), args)
	require.NoError(t, err)
}
