package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/reflection/frog-blossom-cms/internal/roles"
	"net/http"
)

func Authorization(requiredAction roles.Action) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")

		// check if the role has the required action permission
		if !hasPermission(role, requiredAction) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "you don't have permission to access this resource"})
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
