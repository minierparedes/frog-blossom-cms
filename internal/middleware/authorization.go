package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/reflection/frog-blossom-cms/common"
	"github.com/reflection/frog-blossom-cms/internal/roles"
	"github.com/reflection/frog-blossom-cms/token"
	"net/http"
)

func Authorization(requiredAction roles.Action) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.MustGet(AuthorizationPayloadKey).(*token.Payload)

		// check if the role has the required action permission
		if !hasPermission(role.Role, requiredAction) {
			err := errors.New("you don't have permission to access this resource")
			ctx.JSON(http.StatusUnauthorized, common.ErrorResponse(err))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func hasPermission(role string, action roles.Action) bool {
	actions, exists := roles.RolePermissions[role]
	if !exists {
		return false
	}

	for _, act := range actions {
		if act == action {
			return true
		}
	}
	return false
}
