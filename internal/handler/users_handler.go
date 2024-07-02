package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"net/http"
)

// createUsersRequest represents the request payload for creating a user
// @Description Request parameters for creating a user
type createUsersRequest struct {
	Username    string         `json:"username" binding:"required"`
	Email       string         `json:"email" binding:"required"`
	Password    string         `json:"password" binding:"required"`
	Role        string         `json:"role" binding:"required"`
	FirstName   string         `json:"first_name" binding:"required"`
	LastName    string         `json:"last_name" binding:"required"`
	UserUrl     sql.NullString `json:"user_url" binding:"required"`
	Description sql.NullString `json:"description" binding:"required"`
}

// CreateUsersHandler handles the request to create a user
// @Summary Create a user
// @Description Create a new user with the provided parameters
// @Tags users
// @Accept json
// @Produce json
// @Param createUsersRequest body createUsersRequest true "Create User Request"
// @Success 200 {object} db.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
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
			UserUrl:     req.UserUrl,
			Description: req.Description,
		}

		user, err := store.CreateUsers(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

// getUsersRequest represents the request URI parameters for getting a user
// @Description Request parameters for getting a user
type getUsersRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetUsersHandler handles the request to get a user by ID
// @Summary Get a user
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} db.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [get]
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

// listUsersRequest represents the query parameters for the ListUsersHandler
// @Description Request parameters for listing users
type listUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// ListUsersHandler handles the request to list users
// @Summary List users
// @Description List users with pagination
// @Tags users
// @Accept json
// @Produce json
// @Param page_id query int true "Page ID"
// @Param page_size query int true "Page Size"
// @Success 200 {array} db.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [get]
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

// updateUserRequest represents the body for updating a user
// @Description Request parameters for updating a user
type updateUserRequest struct {
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Role        string         `json:"role"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	UserUrl     sql.NullString `json:"user_url"`
	Description sql.NullString `json:"description"`
}

// UpdateUserHandler handles the request to update a user
// @Summary Update a user
// @Description Update a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param body body updateUserRequest true "User update request"
// @Success 200 {object} db.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [put]
func UpdateUserHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req updateUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		var uri struct {
			ID int64 `uri:"id" binding:"required,min=1"`
		}
		if err := ctx.ShouldBindUri(&uri); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		}

		users, err := store.GetUsers(ctx, uri.ID)
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
		}

		user, err := store.UpdateUsers(ctx, args)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

// softDeleteUserRequest represents the URI parameter for the SoftDeleteUsersHandler
// @Description Request parameter for soft deleting a user
type softDeleteUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// SoftDeleteUsersHandler handles the request to soft delete a user
// @Summary Soft delete a user
// @Description Soft delete a user by setting the IsDeleted flag to true
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} db.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id}/soft_delete [delete]
func SoftDeleteUsersHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req softDeleteUserRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := store.GetUsers(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}
			return
		}

		if user.ID != req.ID {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user's ID does not match the provided user ID"})
			return
		}

		arg := db.SoftDeleteUsersParams{
			ID: user.ID,
			IsDeleted: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		}

		result, err := store.SoftDeleteUsers(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
