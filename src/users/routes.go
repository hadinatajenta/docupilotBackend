package users

import (
	"gobackend/core/middleware"
	"gobackend/src/permission"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *Handler, repo permission.PermissionChecker) {
	usersGroup := r.Group("/users")
	{
		usersGroup.GET("/me",
			middleware.AuthMiddleware(),
			permission.PermissionMiddleware(repo, "view_self_profile"),
			h.GetDetailByFirebaseUID,
		)

		usersGroup.GET("/", middleware.AuthMiddleware(), h.GetUsers)
		usersGroup.POST("/", middleware.AuthMiddleware(), h.CreateUser)
	}
}
