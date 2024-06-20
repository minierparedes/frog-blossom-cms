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

		args := db.CreateMetaParams{
			PageID:          sql.NullInt64{Int64: page.ID, Valid: true},
			PostsID:         req.PostsID,
			MetaTitle:       req.MetaTitle,
			MetaDescription: req.MetaDescription,
			MetaRobots:      req.MetaRobots,
			MetaOgImage:     req.MetaOgImage,
			Locale:          req.Locale,
			PageAmount:      req.PageAmount,
			SiteLanguage:    req.SiteLanguage,
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
