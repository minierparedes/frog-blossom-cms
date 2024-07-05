package api

import (
	"github.com/gin-gonic/gin"
	"github.com/reflection/frog-blossom-cms/config"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/reflection/frog-blossom-cms/util"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := config.Config{
		TokenSystemmetricKey: util.RandomString(32),
		AccessTokenDuration:  time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
