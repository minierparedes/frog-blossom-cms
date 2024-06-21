package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"net/http"
)

type createMetaRequest struct {
	PageID          sql.NullInt64  `json:"page_id"`
	PostsID         sql.NullInt64  `json:"posts_id"`
	MetaTitle       sql.NullString `json:"meta_title"`
	MetaDescription sql.NullString `json:"meta_description"`
	MetaRobots      sql.NullString `json:"meta_robots"`
	MetaOgImage     sql.NullString `json:"meta_og_image"`
	Locale          sql.NullString `json:"locale"`
	PageAmount      int64          `json:"page_amount"`
	SiteLanguage    sql.NullString `json:"site_language"`
	MetaKey         string         `json:"meta_key"`
	MetaValue       string         `json:"meta_value"`
}

func CreateMetaHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createMetaRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		page, err := store.GetPages(ctx, req.PageID.Int64)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		post, err := store.GetPosts(ctx, req.PostsID.Int64)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		args := db.CreateMetaParams{
			PageID:          sql.NullInt64{Int64: page.ID, Valid: true},
			PostsID:         sql.NullInt64{Int64: post.ID, Valid: true},
			MetaTitle:       sql.NullString{String: req.MetaTitle.String, Valid: true},
			MetaDescription: sql.NullString{String: req.MetaDescription.String, Valid: true},
			MetaRobots:      sql.NullString{String: req.MetaRobots.String, Valid: true},
			MetaOgImage:     sql.NullString{String: req.MetaOgImage.String, Valid: true},
			Locale:          sql.NullString{String: req.Locale.String, Valid: true},
			PageAmount:      req.PageAmount,
			SiteLanguage:    sql.NullString{String: req.SiteLanguage.String, Valid: true},
			MetaKey:         req.MetaKey,
			MetaValue:       req.MetaValue,
		}

		meta, err := store.CreateMeta(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, meta)
	}
}
