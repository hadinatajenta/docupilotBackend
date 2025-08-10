package roles

import (
	"gobackend/core/middleware"
	"gobackend/src/permission"

	"github.com/gin-gonic/gin"
)

func RegisRoleRoutes(router *gin.RouterGroup, handler *Handler, repo permission.PermissionChecker) {
	role := router.Group("/role")
	{
		role.POST("/create",
			middleware.AuthMiddleware(),
			permission.PermissionMiddleware(repo, "add_roles"),
			handler.CreateRole,
		)
	}
}
