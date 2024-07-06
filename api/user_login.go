package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"time"
)

type loginUsersRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type userResponse struct {
	ID          int64          `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Role        string         `json:"role"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	UserUrl     sql.NullString `json:"user_url"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Role:        user.Role,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		UserUrl:     user.UserUrl,
		Description: user.Description,
		CreatedAt:   user.CreatedAt,
	}
}

func loginUser(store db.Store) gin.HandlerFunc {

}
