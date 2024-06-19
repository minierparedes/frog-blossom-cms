package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	mockdb "github.com/reflection/frog-blossom-cms/db/mock"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPageHandler(t *testing.T) {
	// Arrange
	page := newPage()

	controller := gomock.NewController(t)
	defer controller.Finish()

	store := mockdb.NewMockStore(controller)

	store.EXPECT().
		GetPages(gomock.Any(), gomock.Eq(page.ID)).
		Times(1).
		Return(page, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	// Act
	url := fmt.Sprintf("/api/v1/pages/%d", page.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	// Assert
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchPage(t, recorder.Body, page)
}

func newPage() db.Page {
	return db.Page{
		ID:             5,
		Domain:         "example.com",
		AuthorID:       5,
		PageAuthor:     "sdsdsds",
		Title:          "Homepage",
		Url:            "/home",
		MenuOrder:      1,
		ComponentType:  "Text",
		ComponentValue: "Welcome to our website!",
		PageIdentifier: "home",
		OptionID:       98765,
		OptionName:     "site_title",
		OptionValue:    "My Website",
		OptionRequired: true,
	}
}

func requireBodyMatchPage(t *testing.T, body *bytes.Buffer, page db.Page) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var getPage db.Page
	err = json.Unmarshal(data, &getPage)
	require.NoError(t, err)
	require.Equal(t, page, getPage)
}
