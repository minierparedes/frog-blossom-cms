package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"net/http"
)

type createPostsTxRequest struct {
	UserId   int64                 `json:"user_id" binding:"required"`
	Username string                `json:"username" binding:"required"`
	Posts    db.CreatePostsParams  `json:"posts" binding:"required"`
	Metas    db.CreateMetaTxParams `json:"meta"`
}

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

type getPostRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

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

type listPostsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

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

type updatePostsTxRequest struct {
	UserId   int64                 `json:"user_id" binding:"required"`
	Username string                `json:"username" binding:"required"`
	PostId   *int64                `json:"post_id" binding:"required"`
	Posts    db.UpdatePostsParams  `json:"posts"`
	Metas    db.UpdateMetaTxParams `json:"meta"`
}

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

		user, err := store.GetUsers(ctx, uri.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		post, err := store.GetPosts(ctx, *req.PostId)
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
		PageID:          sql.NullInt64{Int64: getInt64(req.Metas.PageID), Valid: true},
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
