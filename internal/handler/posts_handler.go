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
