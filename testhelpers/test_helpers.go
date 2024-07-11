package testhelpers

import (
	"github.com/reflection/frog-blossom-cms/common"
	"github.com/reflection/frog-blossom-cms/config"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/reflection/frog-blossom-cms/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// NewTestServer helper for handler functions
func NewTestServer(t *testing.T, store db.Store) *common.Server {
	testConfig := config.Config{
		TokenSystemmetricKey: util.RandomString(32),
		AccessTokenDuration:  time.Minute,
	}

	server, err := common.NewTestingServer(testConfig, store)
	require.NoError(t, err)

	return server
}
