package api

import (
	"os"
	"testing"
	"time"

	db "github.com/YounesBouchbouk/natours-api-golang/db/sqlc"
	"github.com/YounesBouchbouk/natours-api-golang/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store *db.Store) *Server {
	config := util.Config{
		JWT_secret_key:      util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(store, &config)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
