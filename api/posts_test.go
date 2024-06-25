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
	"time"
)

func TestGetPostHandler(t *testing.T) {
	// Arrange
	post := newPost()

	controller := gomock.NewController(t)
	defer controller.Finish()

	store := mockdb.NewMockStore(controller)

	store.EXPECT().
		GetPosts(gomock.Any(), gomock.Eq(post.ID)).
		Times(1).
		Return(post, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	// Act
	url := fmt.Sprintf("/api/v1/posts/%d", post.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	// Assert
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchPost(t, recorder.Body, post)
}

func newPost() db.Post {
	return db.Post{
		ID:           550,
		Title:        "Lorem ipsum dolor sit amet",
		Content:      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		AuthorID:     975,
		Url:          "https://example.com",
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Time{},
		Status:       "admin",
		PublishedAt:  time.Now().UTC(),
		EditedAt:     time.Time{},
		PostAuthor:   "lofice",
		PostMimeType: "text/plain",
		PublishedBy:  "lofice",
		UpdatedBy:    "lofice",
	}
}

func requireBodyMatchPost(t *testing.T, body *bytes.Buffer, post db.Post) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var getPost db.Post
	err = json.Unmarshal(data, &getPost)
	require.NoError(t, err)
	require.Equal(t, post.ID, getPost.ID)
	require.Equal(t, post.Title, getPost.Title)
	require.Equal(t, post.Content, getPost.Content)
	require.Equal(t, post.AuthorID, getPost.AuthorID)
	require.Equal(t, post.Url, getPost.Url)
	require.WithinDuration(t, post.UpdatedAt, getPost.UpdatedAt, time.Millisecond)
	require.Equal(t, post.Status, getPost.Status)
	require.WithinDuration(t, post.PublishedAt, getPost.PublishedAt, time.Millisecond)
	require.WithinDuration(t, post.EditedAt, getPost.EditedAt, time.Millisecond)
	require.Equal(t, post.PostAuthor, getPost.PostAuthor)
	require.Equal(t, post.PostMimeType, getPost.PostMimeType)
	require.Equal(t, post.PublishedBy, getPost.PublishedBy)
	require.Equal(t, post.UpdatedBy, getPost.UpdatedBy)
}
