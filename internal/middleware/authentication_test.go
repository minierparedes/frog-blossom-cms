package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/reflection/frog-blossom-cms/api"
	"github.com/reflection/frog-blossom-cms/token"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func addAuthorization(t *testing.T, req *http.Request, tokenMaker token.Maker, authorizationType string, username string, duration time.Duration) {
	testToken, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, testToken)
	req.Header.Set(AuthorizationHeaderKey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name      string
		setupAuth func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		checkResp func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, req, tokenMaker, AuthorizationTypeBearer, "user", time.Minute)
			},
			checkResp: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := api.NewTestServer(t, nil)

			authPath := "/auth"
			server.Router.GET(
				authPath,
				AuthenticationMiddleware(server.TokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResp(t, recorder)
		})
	}
}
