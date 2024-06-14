package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
)

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// CreateUsers handler

type createUsersRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	UserUrl     string `json:"user_url" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CreateUsersHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req createUsersRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		args := db.CreateUsersParams{
			Username:    req.Username,
			Email:       req.Email,
			Password:    req.Password,
			Role:        req.Role,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			UserUrl:     sql.NullString{String: req.UserUrl, Valid: true},
			Description: sql.NullString{String: req.Description, Valid: true},
		}

		user, err := store.CreateUsers(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

// GetUsers handler

type getUsersRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func GetUsersHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req getUsersRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := store.GetUsers(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		ctx.JSON(http.StatusOK, user)
	}
}

// ListUsers handler

type listUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func ListUsersHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req listUsersRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		args := db.ListUsersParams{
			Limit:  req.PageSize,
			Offset: (req.PageID - 1) * req.PageSize,
		}

		users, err := store.ListUsers(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}

// UpdateUsers handler

type updateUsersRequest struct {
	ID          int64  `json:"id" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	UserUrl     string `json:"user_url" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsDeleted   bool   `json:"is_deleted"`
}

func checkID(store *db.Store, req updateUsersRequest) bool {
	if req.ID == 0 {
		return false
	}

	user, err := store.GetUsers(context.Background(), req.ID)
	if err != nil {
		return false
	}

	if user.Username == req.Username {
		return true
	} else {
		return false
	}
}

func UpdateUsersHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req updateUsersRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		args := db.UpdateUsersParams{
			ID:          req.ID,
			Username:    req.Username,
			Email:       req.Email,
			Password:    req.Password,
			Role:        req.Role,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			UserUrl:     sql.NullString{String: req.UserUrl, Valid: true},
			Description: sql.NullString{String: req.Description, Valid: true},
			UpdatedAt:   time.Now(),
			IsDeleted:   sql.NullBool{Bool: req.IsDeleted, Valid: true},
		}

		if !checkID(store, req) {
			ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("unauthorized")))
			return
		}

		user, err := store.UpdateUsers(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
