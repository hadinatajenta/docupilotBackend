package users

import (
	"gobackend/core/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *Handler) {

	usersGroup := r.Group("/users")
	{
		usersGroup.GET("/me", middleware.AuthMiddleware(), h.GetDetailByFirebaseUID)
		// usersGroup.POST("/sync", middleware.FirebaseAuthMiddleware(authClient), h.SyncUser)
		// usersGroup.PUT("/profile", middleware.AuthMiddleware(), h.UpdateUserProfile)
	}
}
