package frog_blossom_db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestInitSetupConfigTx(t *testing.T) {
	// Arrange
	store := NewStore(testDB)

	newUser := createRandomUser(t)
	newUser2 := createRandomUser(t)

	newPages := createRandomPage(t)
	newPages2 := createRandomPage(t)

	newPosts := createRandomPosts(t)
	newPosts2 := createRandomPosts(t)

	n := 5

	errs := make(chan error)
	results := make(chan InitSetupConfigTxResult)

	// Act
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
						MetaTitle: sql.NullString{
							String: "Sample Meta Title 1", Valid: true,
						},
						MetaDescription: sql.NullString{
							String: "Sample Meta Description 1", Valid: true,
						},
						MetaRobots: sql.NullString{
							String: "noindex, nofollow", Valid: true,
						},
						MetaOgImage: sql.NullString{
							String: "https://example.com/image1.jpg", Valid: true,
						},
						Locale: sql.NullString{
							String: "en_US", Valid: true,
						},
						PageAmount:   100,
						SiteLanguage: sql.NullString{String: "en", Valid: true},
						MetaKey:      "sample_key_1",
						MetaValue:    "sample_value_1",
					},
					{
						PageID: sql.NullInt64{
							Int64: newPages2.ID, Valid: true,
						},
						PostsID: sql.NullInt64{
							Int64: newPosts2.ID, Valid: true,
						},
						MetaTitle: sql.NullString{
							String: "Sample Meta Title 2", Valid: true,
						},
						MetaDescription: sql.NullString{
							String: "Sample Meta Description 2", Valid: true,
						},
						MetaRobots: sql.NullString{
							String: "index, follow", Valid: true,
						},
						MetaOgImage: sql.NullString{
							String: "https://example.com/image2.jpg", Valid: true,
						},
						Locale: sql.NullString{
							String: "fr_FR", Valid: true,
						},
						PageAmount:   200,
						SiteLanguage: sql.NullString{String: "fr", Valid: true},
						MetaKey:      "sample_key_2",
						MetaValue:    "sample_value_2",
					},
				},
			})

			errs <- err
			results <- result
		}()
	}

	// Assert
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

		metas := result.Metas
		require.NotEmpty(t, metas)

		for _, meta := range metas {
			storeMeta, err := store.GetMeta(context.Background(), meta.ID)
			require.NoError(t, err)
			require.NotEmpty(t, storeMeta)
			require.Equal(t, meta.ID, storeMeta.ID)
			require.Equal(t, meta.PageID, storeMeta.PageID)
			require.Equal(t, meta.PostsID, storeMeta.PostsID)
			require.Equal(t, meta.MetaTitle, storeMeta.MetaTitle)
			require.Equal(t, meta.MetaDescription, storeMeta.MetaDescription)
			require.Equal(t, meta.MetaRobots, storeMeta.MetaRobots)
			require.Equal(t, meta.MetaOgImage, storeMeta.MetaOgImage)
			require.Equal(t, meta.Locale, storeMeta.Locale)
			require.Equal(t, meta.PageAmount, storeMeta.PageAmount)
			require.Equal(t, meta.SiteLanguage, storeMeta.SiteLanguage)
			require.Equal(t, meta.MetaKey, storeMeta.MetaKey)
			require.Equal(t, meta.MetaValue, storeMeta.MetaValue)
		}
	}
}

func TestCreatePostsTx(t *testing.T) {
	// Arrange
	store := NewStore(testDB)

	newUser := createRandomUser(t)
	newPage := createRandomPage(t)
	newPosts := createRandomPosts(t)

	n := 5

	errs := make(chan error)
	results := make(chan CreateContentTxResult)

	// Act
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreatePostsTx(context.Background(), CreateContentTxParams{
				UserId:   newUser.ID,
				Username: newUser.Username,
				PageId:   &newPage.ID,
				Posts: []CreatePostsParams{
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
				Metas: []CreateMetaParams{
					{
						PageID: sql.NullInt64{
							Int64: newPage.ID, Valid: true,
						},
						PostsID: sql.NullInt64{
							Int64: newPosts.ID, Valid: true,
						},
						MetaTitle: sql.NullString{
							String: "Sample Meta Title 1", Valid: true,
						},
						MetaDescription: sql.NullString{
							String: "Sample Meta Description 1", Valid: true,
						},
						MetaRobots: sql.NullString{
							String: "noindex, nofollow", Valid: true,
						},
						MetaOgImage: sql.NullString{
							String: "https://example.com/image1.jpg", Valid: true,
						},
						Locale: sql.NullString{
							String: "en_US", Valid: true,
						},
						PageAmount:   100,
						SiteLanguage: sql.NullString{String: "en", Valid: true},
						MetaKey:      "sample_key_1",
						MetaValue:    "sample_value_1",
					},
					{
						PageID: sql.NullInt64{
							Int64: newPage.ID, Valid: true,
						},
						PostsID: sql.NullInt64{
							Int64: newPosts.ID, Valid: true,
						},
						MetaTitle: sql.NullString{
							String: "Sample Meta Title 2", Valid: true,
						},
						MetaDescription: sql.NullString{
							String: "Sample Meta Description 2", Valid: true,
						},
						MetaRobots: sql.NullString{
							String: "index, follow", Valid: true,
						},
						MetaOgImage: sql.NullString{
							String: "https://example.com/image2.jpg", Valid: true,
						},
						Locale: sql.NullString{
							String: "fr_FR", Valid: true,
						},
						PageAmount:   200,
						SiteLanguage: sql.NullString{String: "fr", Valid: true},
						MetaKey:      "sample_key_2",
						MetaValue:    "sample_value_2",
					},
				},
			})

			errs <- err
			results <- result
		}()
	}

	// Assert
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		user := result.User
		require.NotEmpty(t, user)
		require.Equal(t, newUser.ID, user.ID)
		require.Equal(t, newUser.Username, user.Username)

		_, err = store.GetUsers(context.Background(), user.ID)
		require.NoError(t, err)

		page := result.PageId
		require.NotEmpty(t, page)
		require.Equal(t, newPage.ID, page.ID)

		_, err = store.GetPages(context.Background(), page.ID)
		require.NoError(t, err)

		posts := result.Posts
		require.NotEmpty(t, posts)

		for _, post := range posts {
			storePosts, err := store.GetPosts(context.Background(), post.ID)
			require.NoError(t, err)
			require.NotEmpty(t, storePosts)
			require.Equal(t, post.ID, storePosts.ID)
			require.Equal(t, post.Title, storePosts.Title)
			require.Equal(t, post.Content, storePosts.Content)
			require.Equal(t, post.AuthorID, storePosts.AuthorID)
			require.Equal(t, post.Url, storePosts.Url)
			require.Equal(t, post.Status, storePosts.Status)
			require.Equal(t, post.PublishedAt, storePosts.PublishedAt)
			require.Equal(t, post.EditedAt, storePosts.EditedAt)
			require.Equal(t, post.PostAuthor, storePosts.PostAuthor)
			require.Equal(t, post.PostMimeType, storePosts.PostMimeType)
			require.Equal(t, post.PublishedBy, storePosts.PublishedBy)
			require.Equal(t, post.UpdatedBy, storePosts.UpdatedBy)
		}

		meta := result.Metas
		require.NotEmpty(t, meta)

		for _, meta := range meta {
			storeMeta, err := store.GetMeta(context.Background(), meta.ID)
			require.NoError(t, err)
			require.NotEmpty(t, storeMeta)
			require.Equal(t, meta.PageID, storeMeta.PageID)
			require.Equal(t, meta.PostsID, storeMeta.PostsID)
			require.Equal(t, meta.MetaTitle, storeMeta.MetaTitle)
			require.Equal(t, meta.MetaDescription, storeMeta.MetaDescription)
			require.Equal(t, meta.MetaRobots, storeMeta.MetaRobots)
			require.Equal(t, meta.MetaOgImage, storeMeta.MetaOgImage)
			require.Equal(t, meta.Locale, storeMeta.Locale)
			require.Equal(t, meta.PageAmount, storeMeta.PageAmount)
			require.Equal(t, meta.SiteLanguage, storeMeta.SiteLanguage)
			require.Equal(t, meta.MetaKey, storeMeta.MetaKey)
			require.Equal(t, meta.MetaValue, storeMeta.MetaValue)
		}
	}
}

func TestCreatePageTx(t *testing.T) {
	// Arrange
	store := NewStore(testDB)

	newUser := createRandomUser(t)
	newPage := createRandomPage(t)
	newPost := createRandomPosts(t)

	n := 5

	errs := make(chan error)
	results := make(chan CreateContentTxResult)

	// Act
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreatePageTx(context.Background(), CreateContentTxParams{
				UserId:   newUser.ID,
				Username: newUser.Username,
				PostId:   &newPost.ID,
				Pages: []CreatePagesParams{
					{
						Domain:         newPage.Domain,
						AuthorID:       newUser.ID,
						PageAuthor:     newUser.Username,
						Title:          newPage.Title,
						Url:            newPage.Url,
						MenuOrder:      newPage.MenuOrder,
						ComponentType:  newPage.ComponentType,
						ComponentValue: newPage.ComponentValue,
						PageIdentifier: newPage.PageIdentifier,
						OptionID:       newPage.OptionID,
						OptionName:     newPage.OptionName,
						OptionValue:    newPage.OptionValue,
						OptionRequired: false,
					},
					{
						Domain:         newPage.Domain,
						AuthorID:       newUser.ID,
						PageAuthor:     newUser.Username,
						Title:          newPage.Title,
						Url:            newPage.Url,
						MenuOrder:      newPage.MenuOrder,
						ComponentType:  newPage.ComponentType,
						ComponentValue: newPage.ComponentValue,
						PageIdentifier: newPage.PageIdentifier,
						OptionID:       newPage.OptionID,
						OptionName:     newPage.OptionName,
						OptionValue:    newPage.OptionValue,
						OptionRequired: false,
					},
				},
				Metas: []CreateMetaParams{
					{
						PageID: sql.NullInt64{
							Int64: newPage.ID, Valid: true,
						},
						PostsID: sql.NullInt64{
							Int64: newPost.ID, Valid: true,
						},
						MetaTitle: sql.NullString{
							String: "Sample Meta Title 1", Valid: true,
						},
						MetaDescription: sql.NullString{
							String: "Sample Meta Description 1", Valid: true,
						},
						MetaRobots: sql.NullString{
							String: "noindex, nofollow", Valid: true,
						},
						MetaOgImage: sql.NullString{
							String: "https://example.com/image1.jpg", Valid: true,
						},
						Locale: sql.NullString{
							String: "en_US", Valid: true,
						},
						PageAmount:   100,
						SiteLanguage: sql.NullString{String: "en", Valid: true},
						MetaKey:      "sample_key_1",
						MetaValue:    "sample_value_1",
					},
					{
						PageID: sql.NullInt64{
							Int64: newPage.ID, Valid: true,
						},
						PostsID: sql.NullInt64{
							Int64: newPost.ID, Valid: true,
						},
						MetaTitle: sql.NullString{
							String: "Sample Meta Title 2", Valid: true,
						},
						MetaDescription: sql.NullString{
							String: "Sample Meta Description 2", Valid: true,
						},
						MetaRobots: sql.NullString{
							String: "index, follow", Valid: true,
						},
						MetaOgImage: sql.NullString{
							String: "https://example.com/image2.jpg", Valid: true,
						},
						Locale: sql.NullString{
							String: "fr_FR", Valid: true,
						},
						PageAmount:   200,
						SiteLanguage: sql.NullString{String: "fr", Valid: true},
						MetaKey:      "sample_key_2",
						MetaValue:    "sample_value_2",
					},
				},
			})

			errs <- err
			results <- result
		}()
	}

	// Assert
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		user := result.User
		require.NotEmpty(t, user)
		require.Equal(t, newUser.ID, user.ID)
		require.Equal(t, newUser.Username, user.Username)

		_, err = store.GetUsers(context.Background(), newUser.ID)
		require.NoError(t, err)
	}
}

func TestUpdatePostsTx(t *testing.T) {
	// Arrange
	store := NewStore(testDB)

	newUser := createRandomUser(t)
	newMeta := createRandomMeta(t)
	now := time.Now().UTC()

	n := 5

	errs := make(chan error)
	results := make(chan UpdateContentTxResult)

	// Act
	for i := 0; i < n; i++ {
		go func() {

			postMeta, err := store.GetMetaByPostsIDForUpdate(context.Background(), sql.NullInt64{
				Int64: newMeta.PostsID.Int64,
				Valid: true,
			})
			require.NoError(t, err)

			result, err := store.UpdatePostsTx(context.Background(), UpdateContentTxParams{
				UserId:     newUser.ID,
				Username:   newUser.Username,
				PageId:     nil,
				PostId:     &postMeta.PostsID.Int64,
				MetaPageID: nil,
				MetaPostID: &postMeta.PostsID.Int64,
				Pages:      nil,
				Posts: []UpdatePostsParams{
					{
						ID:           postMeta.PostsID.Int64,
						Title:        "Lorem ipsum dolor sit amet",
						Content:      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
						AuthorID:     newUser.ID,
						Url:          "https://example.com",
						UpdatedAt:    now,
						Status:       "admin",
						PublishedAt:  now,
						EditedAt:     now,
						PostAuthor:   newUser.Username,
						PostMimeType: "text/plain",
						PublishedBy:  newUser.Username,
						UpdatedBy:    newUser.Username,
					},
				},
				Metas: []UpdateMetaParams{
					{
						ID:              postMeta.ID,
						PageID:          sql.NullInt64{Int64: 0, Valid: false},
						PostsID:         sql.NullInt64{Int64: postMeta.PostsID.Int64, Valid: true},
						MetaTitle:       sql.NullString{String: "Sample Meta Title", Valid: true},
						MetaDescription: sql.NullString{String: "Sample Meta Description", Valid: true},
						MetaRobots:      sql.NullString{String: "index, follow", Valid: true},
						MetaOgImage:     sql.NullString{String: "https://example.com/image.jpg", Valid: true},
						Locale:          sql.NullString{String: "ja_JP", Valid: true},
						PageAmount:      3,
						SiteLanguage: sql.NullString{
							String: "ja", Valid: true,
						},
						MetaKey:   "_thumbnail_id",
						MetaValue: "12345",
					},
				},
			})

			errs <- err
			results <- result
		}()
	}

	// Assert
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		user := result.User
		require.NotEmpty(t, user)
		require.Equal(t, newUser.ID, user.ID)
		require.Equal(t, newUser.Username, user.Username)

		_, err = store.GetUsers(context.Background(), user.ID)
		require.NoError(t, err)

		posts := result.Posts
		require.NotEmpty(t, posts)

		for _, post := range posts {
			storePosts, err := store.GetPosts(context.Background(), post.ID)
			require.NoError(t, err)
			require.NotEmpty(t, storePosts)
			require.Equal(t, post.ID, storePosts.ID)
			require.Equal(t, post.Title, storePosts.Title)
			require.Equal(t, post.Content, storePosts.Content)
			require.Equal(t, post.AuthorID, storePosts.AuthorID)
			require.Equal(t, post.Url, storePosts.Url)
			require.Equal(t, post.Status, storePosts.Status)
			require.Equal(t, post.PublishedAt, storePosts.PublishedAt)
			require.Equal(t, post.EditedAt, storePosts.EditedAt)
			require.Equal(t, post.PostAuthor, storePosts.PostAuthor)
			require.Equal(t, post.PostMimeType, storePosts.PostMimeType)
			require.Equal(t, post.PublishedBy, storePosts.PublishedBy)
			require.Equal(t, post.UpdatedAt, storePosts.UpdatedAt)
		}

		metas := result.Metas
		require.NotEmpty(t, metas)

		for _, meta := range metas {
			storeMeta, err := store.GetMeta(context.Background(), meta.ID)
			require.NoError(t, err)
			require.NotEmpty(t, storeMeta)
			require.Equal(t, meta.ID, storeMeta.ID)
			require.Equal(t, meta.PostsID, storeMeta.PostsID)
			require.Equal(t, meta.MetaTitle, storeMeta.MetaTitle)
			require.Equal(t, meta.MetaDescription, storeMeta.MetaDescription)
			require.Equal(t, meta.MetaRobots, storeMeta.MetaRobots)
			require.Equal(t, meta.MetaOgImage, storeMeta.MetaOgImage)
			require.Equal(t, meta.Locale, storeMeta.Locale)
			require.Equal(t, meta.PageAmount, storeMeta.PageAmount)
			require.Equal(t, meta.SiteLanguage, storeMeta.SiteLanguage)
			require.Equal(t, meta.MetaKey, storeMeta.MetaKey)
			require.Equal(t, meta.MetaValue, storeMeta.MetaValue)
		}
	}
}

func TestUpdatePageTx(t *testing.T) {
	// Arrange
	store := NewStore(testDB)

	newUser := createRandomUser(t)
	newMeta := createRandomMeta(t)

	n := 5

	errs := make(chan error)
	results := make(chan UpdateContentTxResult)

	// Act
	for i := 0; i < n; i++ {
		go func() {

			postMeta, err := store.GetMetaByPageIDForUpdate(context.Background(), sql.NullInt64{
				Int64: newMeta.PageID.Int64,
				Valid: true,
			})
			require.NoError(t, err)

			result, err := store.UpdatePageTx(context.Background(), UpdateContentTxParams{
				UserId:     newUser.ID,
				Username:   newUser.Username,
				PageId:     &postMeta.PageID.Int64,
				PostId:     nil,
				MetaPageID: &postMeta.PageID.Int64,
				MetaPostID: nil,
				Pages: []UpdatePagesParams{
					{
						ID:             postMeta.PageID.Int64,
						Domain:         "example.com",
						AuthorID:       newUser.ID,
						PageAuthor:     newUser.Username,
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
					},
				},
				Posts: nil,
				Metas: []UpdateMetaParams{
					{
						ID:              postMeta.ID,
						PageID:          sql.NullInt64{Int64: postMeta.PageID.Int64, Valid: true},
						PostsID:         sql.NullInt64{Int64: 0, Valid: false},
						MetaTitle:       sql.NullString{String: "Sample Meta Title", Valid: true},
						MetaDescription: sql.NullString{String: "Sample Meta Description", Valid: true},
						MetaRobots:      sql.NullString{String: "index, follow", Valid: true},
						MetaOgImage:     sql.NullString{String: "https://example.com/image.jpg", Valid: true},
						Locale:          sql.NullString{String: "ja_JP", Valid: true},
						PageAmount:      3,
						SiteLanguage: sql.NullString{
							String: "ja", Valid: true,
						},
						MetaKey:   "_thumbnail_id",
						MetaValue: "12345",
					},
				},
			})

			errs <- err
			results <- result
		}()
	}

	// Assert
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		user := result.User
		require.NotEmpty(t, user)
		require.Equal(t, newUser.ID, user.ID)
		require.Equal(t, newUser.Username, user.Username)

		_, err = store.GetUsers(context.Background(), user.ID)
		require.NoError(t, err)

		pages := result.Pages
		require.NotEmpty(t, pages)

		for _, page := range pages {
			storePage, err := store.GetPages(context.Background(), page.ID)
			require.NoError(t, err)
			require.NotEmpty(t, storePage)
			require.Equal(t, storePage.ID, page.ID)
			require.Equal(t, storePage.Domain, page.Domain)
			require.Equal(t, storePage.AuthorID, page.AuthorID)
			require.Equal(t, storePage.PageAuthor, page.PageAuthor)
			require.Equal(t, storePage.Title, page.Title)
			require.Equal(t, storePage.Url, page.Url)
			require.Equal(t, storePage.MenuOrder, page.MenuOrder)
			require.Equal(t, storePage.ComponentType, page.ComponentType)
			require.Equal(t, storePage.ComponentValue, page.ComponentValue)
			require.Equal(t, storePage.PageIdentifier, page.PageIdentifier)
			require.Equal(t, storePage.OptionID, page.OptionID)
			require.Equal(t, storePage.OptionName, page.OptionName)
			require.Equal(t, storePage.OptionValue, page.OptionValue)
			require.Equal(t, storePage.OptionRequired, page.OptionRequired)
		}

		metas := result.Metas
		require.NotEmpty(t, metas)

		for _, meta := range metas {
			storeMeta, err := store.GetMeta(context.Background(), meta.ID)
			require.NoError(t, err)
			require.NotEmpty(t, storeMeta)
			require.Equal(t, meta.ID, storeMeta.ID)
			require.Equal(t, meta.PageID, storeMeta.PageID)
			require.Equal(t, meta.MetaTitle, storeMeta.MetaTitle)
			require.Equal(t, meta.MetaDescription, storeMeta.MetaDescription)
			require.Equal(t, meta.MetaRobots, storeMeta.MetaRobots)
			require.Equal(t, meta.MetaOgImage, storeMeta.MetaOgImage)
			require.Equal(t, meta.Locale, storeMeta.Locale)
			require.Equal(t, meta.PageAmount, storeMeta.PageAmount)
			require.Equal(t, meta.SiteLanguage, storeMeta.SiteLanguage)
			require.Equal(t, meta.MetaKey, storeMeta.MetaKey)
			require.Equal(t, meta.MetaValue, storeMeta.MetaValue)
		}
	}
}

func TestDeletePostsTx(t *testing.T) {
	// Arrange
	store := NewStore(testDB)

	newPost := createRandomPosts(t)

	n := 5

	errs := make(chan error)

	// Act
	for i := 0; i < n; i++ {
		go func() {

			_, err := store.DeletePostsTx(context.Background(), DeleteContentTxParams{
				PageId: nil,
				PostId: &newPost.ID,
			})

			errs <- err
		}()
	}

	// Assert
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		posts, err := store.GetPosts(context.Background(), newPost.ID)
		require.Error(t, err)
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Empty(t, posts)

		meta, err := store.GetMetaByPostsIDForUpdate(context.Background(), sql.NullInt64{
			Int64: newPost.ID,
			Valid: true,
		})
		require.Error(t, err)
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Empty(t, meta)
	}
}

func TestDeletePageTx(t *testing.T) {
	// Arrange
	store := NewStore(testDB)

	newPage := createRandomPage(t)

	n := 5

	errs := make(chan error)

	// Act
	for i := 0; i < n; i++ {
		go func() {

			_, err := store.DeletePageTx(context.Background(), DeleteContentTxParams{
				PageId: &newPage.ID,
				PostId: nil,
			})

			errs <- err
		}()
	}

	// Assert
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		page, err := store.GetPages(context.Background(), newPage.ID)
		require.Error(t, err)
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Empty(t, page)

		meta, err := store.GetMetaByPageIDForUpdate(context.Background(), sql.NullInt64{
			Int64: newPage.ID,
			Valid: true,
		})
		require.Error(t, err)
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Empty(t, meta)
	}
}
