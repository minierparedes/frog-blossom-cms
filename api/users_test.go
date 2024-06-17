package api

import (
	"database/sql"
	"fmt"
	mockdb "github.com/reflection/frog-blossom-cms/db/mock"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/reflection/frog-blossom-cms/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsersHandler(t *testing.T) {
	// Arrange
	user := newUser()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	// build stubs
	store.EXPECT().
		GetUsers(gomock.Any(), gomock.Eq(user.ID)).
		Times(1).
		Return(user, nil)

	// start test server and send request
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	// Act
	url := fmt.Sprintf("/api/v1/users/%d", user.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	// Assert
	require.Equal(t, http.StatusOK, recorder.Code)
}

func newUser() db.User {
	return db.User{
		ID:        5,
		Username:  utils.RandomUsername(),
		Email:     "cshopcott6@friendfeed.com",
		Password:  "pP7<8jrQbwS",
		Role:      "user",
		FirstName: "Cointon",
		LastName:  "Shopcott",
		UserUrl: sql.NullString{
			String: "https://robohash.org/nihildelectussed.png?size=50x50&set=set1", Valid: true,
		},
		Description: sql.NullString{
			String: "Morbi porttitor lorem id ligula. Suspendisse ornare consequat lectus. In est risus, auctor sed, tristique in, tempus sit amet, sem. Fusce consequat. Nulla nisl. Nunc nisl.", Valid: true,
		},
	}
}
