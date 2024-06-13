package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog_blossom_db/db/sqlc"
)

type createPagesTxRequest struct {
	UserId   int64                  `json:"user_id" binding:"required"`
	Username string                 `json:"username" binding:"required"`
	PageId   *int64                 `json:"page_id"`
	PostId   *int64                 `json:"post_id"`
	Pages    []db.CreatePagesParams `json:"pages" binding:"required"`
	Posts    []db.CreatePostsParams `json:"posts"`
	Metas    []createMetaParams     `json:"meta"`
}

type createMetaParams struct {
	PostsID         *int64  `json:"posts_id"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaRobots      *string `json:"meta_robots"`
	MetaOgImage     *string `json:"meta_og_image"`
	Locale          *string `json:"locale"`
	PageAmount      int64   `json:"page_amount"`
	SiteLanguage    *string `json:"site_language"`
	MetaKey         string  `json:"meta_key"`
	MetaValue       string  `json:"meta_value"`
}

func CreatePagesTxHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createPagesTxRequest
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

type getPageRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func GetPageHandler(store *db.Store) gin.HandlerFunc {
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
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		ctx.JSON(http.StatusOK, page)
	}
}

type listPagesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func ListPagesHandler(store *db.Store) gin.HandlerFunc {
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
		}
		ctx.JSON(http.StatusOK, pages)
	}
}

type updatePagesRequest struct {
	ID             int64  `json:"id"`
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

func UpdatePagesHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req updatePagesRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		var pageID = req.ID

		pages, err := store.GetPages(ctx, pageID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
		}

		user, err := store.GetUsers(ctx, req.AuthorID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
		}

		args := db.UpdatePagesParams{
			ID:             pages.ID,
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

		page, err := store.UpdatePages(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusCreated, page)
	}
}

type updatePagesTxRequest struct {
	UserId     int64                  `json:"user_id" binding:"required"`
	Username   string                 `json:"username" binding:"required"`
	PageId     *int64                 `json:"page_id"`
	PostId     *int64                 `json:"post_id"`
	MetaPageID *int64                 `json:"meta_page_id"`
	MetaPostID *int64                 `json:"meta_post_id"`
	Pages      []db.UpdatePagesParams `json:"pages" binding:"required"`
	Posts      []db.UpdatePostsParams `json:"posts"`
	Metas      []updateMetaParams     `json:"meta"`
}

type updateMetaParams struct {
	PageID          *int64  `json:"page_id"`
	PostsID         *int64  `json:"posts_id"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	MetaRobots      *string `json:"meta_robots"`
	MetaOgImage     *string `json:"meta_og_image"`
	Locale          *string `json:"locale"`
	PageAmount      int64   `json:"page_amount"`
	SiteLanguage    *string `json:"site_language"`
	MetaKey         string  `json:"meta_key"`
	MetaValue       string  `json:"meta_value"`
}

func UpdatePagesTxHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req updatePagesTxRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := store.GetUsers(ctx, req.UserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		page, err := store.GetPages(ctx, *req.PageId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		meta, err := store.GetMetaByPageIDForUpdate(ctx, sql.NullInt64{Int64: *req.MetaPageID, Valid: true})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		if user.Username != req.Username {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page's username does not match the provided user ID"})
			return
		}

		// Create pointers to pass to toDBParams
		pageID := &page.ID
		metaPageID := &meta.PageID

		// TODO: run app
		args := req.toDBParams(user.ID, user.Username, pageID, *metaPageID)

		result, err := store.UpdatePageTx(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusCreated, result)
	}
}

func (req *createPagesTxRequest) toDBParams(userID int64, username string) db.CreateContentTxParams {
	var metas []db.CreateMetaParams
	for _, m := range req.Metas {
		metas = append(metas, db.CreateMetaParams{
			PostsID:         sql.NullInt64{Int64: *m.PostsID, Valid: m.PostsID != nil},
			MetaTitle:       sql.NullString{String: *m.MetaTitle, Valid: true},
			MetaDescription: sql.NullString{String: *m.MetaDescription, Valid: true},
			MetaRobots:      sql.NullString{String: *m.MetaRobots, Valid: true},
			MetaOgImage:     sql.NullString{String: *m.MetaOgImage, Valid: true},
			Locale:          sql.NullString{String: *m.Locale, Valid: true},
			PageAmount:      m.PageAmount,
			SiteLanguage:    sql.NullString{String: *m.SiteLanguage, Valid: true},
			MetaKey:         m.MetaKey,
			MetaValue:       m.MetaValue,
		})
	}
	return db.CreateContentTxParams{
		UserId:   userID,
		Username: username,
		PageId:   req.PageId,
		PostId:   req.PostId,
		Pages:    req.Pages,
		Posts:    req.Posts,
		Metas:    metas,
	}
}

func (req *updatePagesTxRequest) toDBParams(userID int64, username string, pageID *int64, metaPageID sql.NullInt64) db.UpdateContentTxParams {
	var metas []db.UpdateMetaParams
	for _, m := range req.Metas {
		metas = append(metas, db.UpdateMetaParams{
			PageID:          sql.NullInt64{Int64: getInt64(m.PageID), Valid: m.PageID != nil},
			PostsID:         sql.NullInt64{Int64: getInt64(m.PostsID), Valid: m.PostsID != nil},
			MetaTitle:       sql.NullString{String: getStr(m.MetaTitle), Valid: true},
			MetaDescription: sql.NullString{String: getStr(m.MetaDescription), Valid: true},
			MetaRobots:      sql.NullString{String: getStr(m.MetaRobots), Valid: true},
			MetaOgImage:     sql.NullString{String: getStr(m.MetaOgImage), Valid: true},
			Locale:          sql.NullString{String: getStr(m.Locale), Valid: true},
			PageAmount:      m.PageAmount,
			SiteLanguage:    sql.NullString{String: getStr(m.SiteLanguage), Valid: true},
			MetaKey:         m.MetaKey,
			MetaValue:       m.MetaValue,
		})
	}
	return db.UpdateContentTxParams{
		UserId:     userID,
		Username:   username,
		PageId:     pageID,
		PostId:     req.PostId,
		MetaPageID: nullInt64ToInt64Pointer(metaPageID),
		MetaPostID: req.MetaPostID,
		Pages:      req.Pages,
		Posts:      req.Posts,
		Metas:      metas,
	}
}

func nullInt64ToInt64Pointer(nullInt sql.NullInt64) *int64 {
	if !nullInt.Valid {
		return nil
	}
	return &nullInt.Int64
}

func getInt64(ptr *int64) int64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

func getStr(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
