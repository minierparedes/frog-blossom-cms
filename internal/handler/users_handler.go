package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog_blossom_db/db/sqlc"
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

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

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
	Limit  int32 `form:"limit" binding:"required,min=1"`
	Offset int32 `form:"offset" binding:"required,min=0"`
}

func ListUsersHandler(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req listUsersRequest
		if err := ctx.ShouldBindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		args := db.ListUsersParams{
			Limit:  req.Limit,
			Offset: req.Offset,
		}

		users, err := store.ListUsers(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}

// Delete user//
// type deleteUsersRequest struct {
// 	ID int64 `uri:"id" binding:"required,min=1"`
// }

// func DeleteUsersHandler(store *db.Store) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var req deleteUsersRequest
// 		if err := ctx.ShouldBindUri(&req); err != nil {
// 			ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 			return
// 		}

// 		err := store.DeleteUsers(ctx, req.ID)
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 			return
// 		}
// 		ctx.Status(http.StatusNoContent)
// 	}
// }

// update user //

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

		user, err := store.UpdateUsers(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
