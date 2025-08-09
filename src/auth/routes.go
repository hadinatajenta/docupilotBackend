package auth

import (
	"gobackend/core/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup, h *Handler) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
		authGroup.POST("/refresh", middleware.AuthMiddleware(), h.RefreshToken)
		authGroup.POST("/logout", middleware.AuthMiddleware(), h.Logout)
	}
}
