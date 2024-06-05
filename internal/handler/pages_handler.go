package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog_blossom_db/db/sqlc"
)

type createPagesRequest struct {
	Domain         string `json:"domain"`
	AuthorID       int64  `json:"author_id"`
	PageAuthor     string `json:"page_author"`
	Title          string `json:"title"`
	Url            string `json:"url"`
	MenuOrder      int64  `json:"menu_order"`
	ComponentType  string `json:"component_type"`
	ComponentValue string `json:"component_value"`
	PageIdentifier string `json:"page_identifier"`
	OptionID       int64  `json:"option_id"`
	OptionName     string `json:"option_name"`
	OptionValue    string `json:"option_value"`
	OptionRequired bool   `json:"option_required"`
}

func CreatePagesHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createPagesRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		args := db.CreatePagesParams{
			Domain:         req.Domain,
			AuthorID:       req.AuthorID,
			PageAuthor:     req.PageAuthor,
			Title:          req.Title,
			Url:            req.Url,
			MenuOrder:      req.MenuOrder,
			ComponentType:  req.ComponentType,
			ComponentValue: req.ComponentValue,
			PageIdentifier: req.PageIdentifier,
			OptionID:       req.OptionID,
			OptionName:     req.OptionName,
			OptionValue:    req.OptionValue,
			OptionRequired: req.OptionRequired,
		}

		page, err := store.CreatePages(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, page)
	}
}
