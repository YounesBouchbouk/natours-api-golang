package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	require.NoError(t, err)

	require.NotEmpty(t, hash)
}

func TestComaprepassord(t *testing.T) {
	password := "younes"

	hash, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hash)

	result := CheckPasswordHash(password, hash)

	require.True(t, result)

	result = CheckPasswordHash("alpha", hash)

	require.False(t, result)

}
