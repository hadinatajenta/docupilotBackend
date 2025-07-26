package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func FirebaseAuthMiddleware(firebaseAuth *auth.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := firebaseAuth.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Firebase token"})
			return
		}

		c.Set("firebase_uid", token.UID)

		if email, ok := token.Claims["email"].(string); ok {
			c.Set("firebase_email", email)
		}
		if name, ok := token.Claims["name"].(string); ok {
			c.Set("firebase_name", name)
		}
		if picture, ok := token.Claims["picture"].(string); ok {
			c.Set("firebase_avatar", picture)
		}

		c.Next()
	}
}
