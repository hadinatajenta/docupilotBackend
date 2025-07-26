package middleware

import (
	"gobackend/shared/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":      false,
				"message":     "missing authorization token",
				"http_status": http.StatusUnauthorized,
			})
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":      false,
				"message":     "invalid token format",
				"http_status": http.StatusUnauthorized,
			})
			return
		}

		tokenStr := splitToken[1]

		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":      false,
				"message":     "invalid token",
				"http_status": http.StatusUnauthorized,
			})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}

}
