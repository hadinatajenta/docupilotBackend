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
			// permission.PermissionMiddleware(repo, "add_roles"),
			handler.CreateRole,
		)

		role.POST("/:id/permissions",
			middleware.AuthMiddleware(),
			// permission.PermissionMiddleware(repo, "assign_permissions"),
			handler.AssignPermissions,
		)
		role.GET("/all",
			middleware.AuthMiddleware(),
			// permission.PermissionMiddleware(repo, "view_roles"),
			handler.GetAllRoles,
		)
		// role.GET("/list", handler.ListRoles)
		// role.GET("/:id", handler.GetRoleByID)
		// role.PUT("/:id", handler.UpdateRole)
		// role.DELETE("/:id", handler.DeleteRole)
	}
}
