package middleware

import (
	"gobackend/shared/response"
	"gobackend/src/roles"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(roleRepo roles.Repository, requiredPerm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id")
		if userID == "" {
			response.Error(c, http.StatusUnauthorized, "unauthorized", nil)
			c.Abort()
			return
		}

		perms, err := roleRepo.GetPermissionsByUserID(c, userID)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "server error", err)
			c.Abort()
			return
		}

		hasPermission := false
		for _, p := range perms {
			if p.Name == requiredPerm {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			response.Error(c, http.StatusForbidden, "forbidden", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
