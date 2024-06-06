package handler

import (
	"database/sql"
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

		user, err := store.GetUsers(ctx, req.AuthorID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		if user.Username != req.PageAuthor {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page author does not match the provided author ID"})
			return
		}

		args := db.CreatePagesParams{
			Domain:         req.Domain,
			AuthorID:       user.ID,
			PageAuthor:     user.Username,
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

type getPagesRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func GetPagesHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req getPagesRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		page, err := store.GetPages(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		ctx.JSON(http.StatusOK, page)
	}
}
