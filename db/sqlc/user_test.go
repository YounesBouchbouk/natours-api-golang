package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/YounesBouchbouk/natours-api-golang/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	password := util.RandomString(10)
	arg := CreateUserParams{
		Email:    util.RandomOwnerEmail(),
		Role:     util.RandomRole(),
		Photo:    "photo.com/img1.png",
		Name:     util.RandomString(6),
		Password: password,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Role, user.Role)
	require.Equal(t, arg.Photo, user.Photo)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Password, user.Password)

	return user

}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomUser(t)
	account2, err := testQueries.GetUserByID(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), account.ID)

	require.NoError(t, err)

	account2, err := testQueries.GetUserByID(context.Background(), account.ID)

	require.Empty(t, account2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}

func TestListUsers(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomUser(t)
	}

	users, err := testQueries.GetAllUsers(context.Background(), 5)

	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}
