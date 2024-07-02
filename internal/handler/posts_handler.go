package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"net/http"
)

// createPostsTxRequest represents the request body for creating a post
// @Description Request parameters for creating a post
type createPostsTxRequest struct {
	UserId   int64                 `json:"user_id" binding:"required"`
	Username string                `json:"username" binding:"required"`
	Posts    db.CreatePostsParams  `json:"posts" binding:"required"`
	Metas    db.CreateMetaTxParams `json:"meta"`
}

// CreatePostTxHandler handles the request to create a post
// @Summary Create a post
// @Description Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param username path string true "Username"
// @Param createPostsTxRequest body true "Post creation request"
// @Success 200 {object} db.Post
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /posts [post]
func CreatePostTxHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createPostsTxRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := store.GetUsers(ctx, req.UserId)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}

		if user.Username != req.Username {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Post's username does not match the provided user ID"})
			return
		}

		args := req.toDBParams(user.ID, user.Username)

		post, err := store.CreatePostsTx(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, post)
	}
}

// getPostRequest represents the request parameters for getting a post
// @Description Request parameters for getting a post
type getPostRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// @BasePath /api/v1

// GetPostHandler handles the request to get a post by ID
// @Summary Get a post by ID
// @Schemes
// @Description Retrieve a specific post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} db.Post
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /posts/{id} [get]
func GetPostHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req getPostRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		post, err := store.GetPosts(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusOK, post)
	}
}

// listPostsRequest represents the query parameters for listing posts
// @Description Request parameters for listing posts
type listPostsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// ListPostsHandler handles the request to list posts
// @Summary List posts with pagination
// @Description Retrieve a list of posts with pagination support
// @Tags posts
// @Accept json
// @Produce json
// @Param page_id query int true "Page number"
// @Param page_size query int true "Page size (minimum 5, maximum 10)"
// @Success 200 {array} db.Post
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /posts [get]
func ListPostsHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req listPostsRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		args := db.ListPostsParams{
			Limit:  req.PageSize,
			Offset: (req.PageID - 1) * req.PageSize,
		}

		posts, err := store.ListPosts(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, posts)
	}
}

// updatePostsTxRequest represents the query parameters for listing posts
// @Description Request parameters for updating a post
type updatePostsTxRequest struct {
	UserId   int64                 `json:"user_id" binding:"required"`
	Username string                `json:"username" binding:"required"`
	Posts    db.UpdatePostsParams  `json:"posts"`
	Metas    db.UpdateMetaTxParams `json:"meta"`
}

// UpdatePostsTxHandler handles the request to update a posts
// @Summary Update a post
// @Description Update a post with specified parameters
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param user_id body int64 true "User ID"
// @Param username body string true "Username"
// @Param db.UpdatePostsParams posts body true "Updated post parameters"
// @Param db.UpdateMetaTxParams meta body true "Updated meta parameters"
// @Success 201 {object} db.Post
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /posts/{id} [put]
func UpdatePostsTxHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req updatePostsTxRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		var uri struct {
			ID int64 `uri:"id" binding:"required,min=1"`
		}
		if err := ctx.ShouldBindUri(&uri); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := store.GetUsers(ctx, req.UserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		post, err := store.GetPosts(ctx, uri.ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		meta, err := store.GetMetaByPostsIDForUpdate(ctx, sql.NullInt64{Int64: post.ID, Valid: true})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		if user.Username != req.Username {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Post's username does not match the provided user ID"})
		}

		metaPostID := sql.NullInt64{Int64: post.ID, Valid: true}

		if meta.PostsID != metaPostID {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Meta's pageID does not match the provided page ID"})
			return
		}

		args := req.toDBParams(user.ID, user.Username, &post.ID)

		result, err := store.UpdatePostsTx(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusCreated, result)
	}
}

// deletePostsRequest represents the request parameters for deleting a post
// @Description Request parameters for deleting a post
type deletePostsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// DeletePostTxHandler handles the request to delete a posts
// @Summary Delete a post
// @Description Delete a post with the specified ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {boolean} true
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /posts/{id} [delete]
func DeletePostTxHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req deletePostsRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		post, err := store.GetPosts(ctx, req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		meta, err := store.GetMetaByPostsIDForUpdate(ctx, sql.NullInt64{Int64: post.ID, Valid: true})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		if post.ID != req.ID {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Post's ID does not match the provided post ID"})
			return
		}

		metaPostId := sql.NullInt64{Int64: post.ID, Valid: true}

		if meta.PostsID != metaPostId {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Post's ID does not match the provided post ID"})
		}

		args := db.DeletePostTxParams{PostId: &post.ID}

		result, err := store.DeletePostsTx(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

// toDBParams converts a createPageTxRequest instance into a db.CreatePageTxParams structure for db operations
func (req *createPostsTxRequest) toDBParams(userID int64, username string) db.CreatePostTxParams {
	dbMetas := db.CreateMetaParams{
		MetaTitle:       sql.NullString{String: *req.Metas.MetaTitle, Valid: true},
		MetaDescription: sql.NullString{String: *req.Metas.MetaDescription, Valid: true},
		MetaRobots:      sql.NullString{String: *req.Metas.MetaRobots, Valid: true},
		MetaOgImage:     sql.NullString{String: *req.Metas.MetaOgImage, Valid: true},
		Locale:          sql.NullString{String: *req.Metas.Locale, Valid: true},
		PageAmount:      req.Metas.PageAmount,
		SiteLanguage:    sql.NullString{String: *req.Metas.SiteLanguage, Valid: true},
		MetaKey:         req.Metas.MetaKey,
		MetaValue:       req.Metas.MetaValue,
	}
	return db.CreatePostTxParams{
		UserId:   userID,
		Username: username,
		Posts:    &req.Posts,
		Metas:    dbMetas,
	}
}

// toDBParams converts a updatePagesTxRequest instance into a db.UpdatePageTxParams structure for db operations
func (req *updatePostsTxRequest) toDBParams(userID int64, username string, postID *int64) db.UpdatePostTxParams {

	dbMetas := db.UpdateMetaParams{
		ID:              req.Metas.ID,
		PostsID:         sql.NullInt64{Int64: getInt64(req.Metas.PostsID), Valid: true},
		MetaTitle:       sql.NullString{String: getStr(req.Metas.MetaTitle), Valid: true},
		MetaDescription: sql.NullString{String: getStr(req.Metas.MetaDescription), Valid: true},
		MetaRobots:      sql.NullString{String: getStr(req.Metas.MetaRobots), Valid: true},
		MetaOgImage:     sql.NullString{String: getStr(req.Metas.MetaOgImage), Valid: true},
		Locale:          sql.NullString{String: getStr(req.Metas.Locale), Valid: true},
		PageAmount:      req.Metas.PageAmount,
		SiteLanguage:    sql.NullString{String: getStr(req.Metas.SiteLanguage), Valid: true},
		MetaKey:         req.Metas.MetaKey,
		MetaValue:       req.Metas.MetaValue,
	}
	return db.UpdatePostTxParams{
		UserId:   userID,
		Username: username,
		PostId:   postID,
		Posts:    &req.Posts,
		Metas:    dbMetas,
	}
}
