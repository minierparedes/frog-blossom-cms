package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
)

type createPageTxRequest struct {
	UserId   int64                 `json:"user_id" binding:"required"`
	Username string                `json:"username" binding:"required"`
	Pages    db.CreatePagesParams  `json:"pages" binding:"required"`
	Metas    db.CreateMetaTxParams `json:"meta"`
}

func CreatePageTxHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createPageTxRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := store.GetUsers(ctx, req.UserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		if user.Username != req.Username {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page's username does not match the provided user ID"})
			return
		}

		args := req.toDBParams(user.ID, user.Username)

		page, err := store.CreatePageTx(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, page)
	}
}

type getPageRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func GetPageHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req getPageRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		page, err := store.GetPages(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusOK, page)
	}
}

type listPagesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func ListPagesHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req listPagesRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		args := db.ListPagesParams{
			Limit:  req.PageSize,
			Offset: (req.PageID - 1) * req.PageSize,
		}

		pages, err := store.ListPages(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, pages)
	}
}

type updatePagesTxRequest struct {
	UserId   int64                 `json:"user_id" binding:"required"`
	Username string                `json:"username" binding:"required"`
	PageId   *int64                `json:"page_id"`
	PostId   *int64                `json:"post_id"`
	Pages    db.UpdatePagesParams  `json:"pages" binding:"required"`
	Posts    db.UpdatePostsParams  `json:"posts"`
	Metas    db.UpdateMetaTxParams `json:"meta"`
}

func UpdatePagesTxHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req updatePagesTxRequest
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

		page, err := store.GetPages(ctx, *req.PageId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		meta, err := store.GetMetaByPageIDForUpdate(ctx, sql.NullInt64{Int64: page.ID, Valid: true})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		if user.Username != req.Username {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page's username does not match the provided user ID"})
			return
		}

		if page.PageAuthor != req.Username {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page's author does not match the provided username"})
			return
		}

		metaPageID := sql.NullInt64{Int64: page.ID, Valid: true}

		if meta.PageID != metaPageID {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Meta's pageID does not match the provided page ID"})
			return
		}

		args := req.toDBParams(user.ID, user.Username, &page.ID)

		result, err := store.UpdatePageTx(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusCreated, result)
	}
}

type deletePageRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func DeletePageTxHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req deletePageRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		page, err := store.GetPages(ctx, req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		meta, err := store.GetMetaByPageIDForUpdate(ctx, sql.NullInt64{Int64: req.ID, Valid: true})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		if page.ID != req.ID {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page's ID does not match the provided page ID"})
			return
		}

		metaPageId := sql.NullInt64{Int64: req.ID, Valid: true}

		if meta.PageID != metaPageId {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page's ID does not match the provided page ID"})
			return
		}

		args := db.DeletePageTxParams{
			PageId: &page.ID,
		}

		result, err := store.DeletePageTx(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)

	}
}

// toDBParams converts a createPageTxRequest instance into a db.CreatePageTxParams structure for db operations
func (req *createPageTxRequest) toDBParams(userID int64, username string) db.CreatePageTxParams {
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
	return db.CreatePageTxParams{
		UserId:   userID,
		Username: username,
		Pages:    &req.Pages,
		Metas:    dbMetas,
	}
}

// toDBParams converts a updatePagesTxRequest instance into a db.UpdatePageTxParams structure for db operations
func (req *updatePagesTxRequest) toDBParams(userID int64, username string, pageID *int64) db.UpdatePageTxParams {

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
	return db.UpdatePageTxParams{
		UserId:   userID,
		Username: username,
		PageId:   pageID,
		Pages:    &req.Pages,
		Metas:    dbMetas,
	}
}

// getInt64 function safely dereferences a pointer int64 to an int64
func getInt64(ptr *int64) int64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// getStr function safely dereferences a pointer string to a string
func getStr(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
