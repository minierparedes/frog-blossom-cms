package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
)

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

func CreateUsersHandler(store db.Store) gin.HandlerFunc {
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

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

type getUsersRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func GetUsersHandler(store db.Store) gin.HandlerFunc {
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

type listUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func ListUsersHandler(store db.Store) gin.HandlerFunc {
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

type updateUserRequest struct {
	ID          int64          `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Role        string         `json:"role"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	UserUrl     sql.NullString `json:"user_url"`
	Description sql.NullString `json:"description"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func UpdateUserHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req updateUserRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		users, err := store.GetUsers(ctx, req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		args := db.UpdateUsersParams{
			ID:          users.ID,
			Username:    req.Username,
			Email:       req.Email,
			Password:    req.Password,
			Role:        req.Role,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			UserUrl:     req.UserUrl,
			Description: req.Description,
			UpdatedAt:   time.Time{},
		}

		user, err := store.UpdateUsers(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

type deleteUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func DeleteUsersHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req deleteUserRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := store.GetUsers(ctx, req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		if user.ID != req.ID {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user's ID does not match the provided user ID"})
			return
		}

		arg := db.UpdateUsersParams{
			ID: user.ID,
			IsDeleted: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		}

		result, err := store.UpdateUsers(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result.IsDeleted)
	}
}
