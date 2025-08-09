package users

import (
	"gobackend/core/middleware"
	"gobackend/src/roles"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup, h *Handler, roleRepo roles.Repository) {

	usersGroup := r.Group("/users")
	{
		usersGroup.GET("/me",
			middleware.AuthMiddleware(),
			middleware.PermissionMiddleware(roleRepo, "view_profile"),
			h.GetDetailByFirebaseUID,
		)
	}
}
