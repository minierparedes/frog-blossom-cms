package frog_blossom_db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestInitSetupConfigTx(t *testing.T) {
	store := NewStore(testDB)

	newUser := createRandomUser(t)
	newUser2 := createRandomUser(t)

	newPages := createRandomPage(t)
	newPosts := createRandomPosts(t)

	n := 5

	errs := make(chan error)
	results := make(chan InitSetupConfigTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.InitSetupConfigTx(context.Background(), InitSetupConfigTxParams{
				UserId:   newUser.ID,
				Username: newUser.Username,
				Email:    newUser.Email,
				UserURl:  newUser.UserUrl.String,
				InitialPages: []CreatePagesParams{
					{
						Domain:         newPages.Domain,
						AuthorID:       newUser.ID,
						PageAuthor:     newUser.Username,
						Title:          newPages.Title,
						Url:            newPages.Url,
						MenuOrder:      newPages.MenuOrder,
						ComponentType:  newPages.ComponentType,
						ComponentValue: newPages.ComponentValue,
						PageIdentifier: newPages.PageIdentifier,
						OptionID:       newPages.OptionID,
						OptionName:     newPages.OptionName,
						OptionValue:    newPages.OptionValue,
						OptionRequired: false,
					},
					{
						Domain:         newPages.Domain,
						AuthorID:       newUser2.ID,
						PageAuthor:     newUser2.Username,
						Title:          newPages.Title,
						Url:            newPages.Url,
						MenuOrder:      newPages.MenuOrder,
						ComponentType:  newPages.ComponentType,
						ComponentValue: newPages.ComponentValue,
						PageIdentifier: newPages.PageIdentifier,
						OptionID:       newPages.OptionID,
						OptionName:     newPages.OptionName,
						OptionValue:    newPages.OptionValue,
						OptionRequired: false,
					},
				},

				InitialPosts: []CreatePostsParams{
					{
						Title:        newPosts.Title,
						Content:      newPosts.Content,
						AuthorID:     newUser.ID,
						Url:          newPosts.Url,
						Status:       newPosts.Status,
						PublishedAt:  time.Time{},
						PostAuthor:   newUser.Username,
						PostMimeType: newPosts.PostMimeType,
						PublishedBy:  newUser.Username,
						UpdatedBy:    newUser.Username,
					}, {
						Title:        newPosts.Title,
						Content:      newPosts.Content,
						AuthorID:     newUser.ID,
						Url:          newPosts.Url,
						Status:       newPosts.Status,
						PublishedAt:  time.Time{},
						PostAuthor:   newUser.Username,
						PostMimeType: newPosts.PostMimeType,
						PublishedBy:  newUser.Username,
						UpdatedBy:    newUser.Username,
					},
				},
				InitialMeta: []CreateMetaParams{
					{
						PageID: sql.NullInt64{
							Int64: newPages.ID, Valid: true,
						},
						PostsID: sql.NullInt64{
							Int64: newPosts.ID, Valid: true,
						},
						MetaTitle:       sql.NullString{},
						MetaDescription: sql.NullString{},
						MetaRobots:      sql.NullString{},
						MetaOgImage:     sql.NullString{},
						Locale:          sql.NullString{},
						PageAmount:      0,
						SiteLanguage:    sql.NullString{},
						MetaKey:         "",
						MetaValue:       "",
					}, {
						PageID: sql.NullInt64{
							Int64: newPages.ID, Valid: true,
						},
						PostsID: sql.NullInt64{
							Int64: newPosts.ID, Valid: true,
						},
						MetaTitle:       sql.NullString{},
						MetaDescription: sql.NullString{},
						MetaRobots:      sql.NullString{},
						MetaOgImage:     sql.NullString{},
						Locale:          sql.NullString{},
						PageAmount:      0,
						SiteLanguage:    sql.NullString{},
						MetaKey:         "",
						MetaValue:       "",
					},
				},
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		user := result.User
		require.NotEmpty(t, user)
		require.Equal(t, newUser.ID, user.ID)
		require.Equal(t, newUser.Username, user.Username)
		require.Equal(t, newUser.Email, user.Email)
		require.Equal(t, newUser.UserUrl, user.UserUrl)

		_, err = store.GetUsers(context.Background(), user.ID)
		require.NoError(t, err)

		pages := result.Pages
		require.NotEmpty(t, pages)

		for _, page := range pages {
			storePage, err := store.GetPages(context.Background(), page.ID)
			require.NoError(t, err)
			require.NotEmpty(t, storePage)
			require.Equal(t, page.ID, storePage.ID)
			require.Equal(t, page.Domain, storePage.Domain)
			require.Equal(t, page.AuthorID, storePage.AuthorID)
			require.Equal(t, page.PageIdentifier, storePage.PageIdentifier)
			require.Equal(t, page.Title, storePage.Title)
			require.Equal(t, page.Url, storePage.Url)
			require.Equal(t, page.MenuOrder, storePage.MenuOrder)
			require.Equal(t, page.ComponentType, storePage.ComponentType)
			require.Equal(t, page.ComponentValue, storePage.ComponentValue)
			require.Equal(t, page.OptionID, storePage.OptionID)
			require.Equal(t, page.OptionName, storePage.OptionName)
			require.Equal(t, page.OptionValue, storePage.OptionValue)
			require.Equal(t, page.OptionRequired, storePage.OptionRequired)

		}
	}
}
