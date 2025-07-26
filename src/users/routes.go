package users

import (
	"gobackend/core/firebase"
	"gobackend/core/middleware"

	"github.com/gin-gonic/gin"
)

func RegisUserRoute(r *gin.RouterGroup, h *Handler) {
	authClient := firebase.MustInit()
	auth := r.Group("/auth")
	auth.POST("/sync", middleware.FirebaseAuthMiddleware(authClient), h.SyncUser)
	auth.POST("/login", h.Login)
	auth.POST("/refresh", middleware.AuthMiddleware(), h.RefreshToken)
}
