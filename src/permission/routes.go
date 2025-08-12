package permission

import (
	"gobackend/core/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPermissionRoutes(router *gin.RouterGroup, handler *Handler, repo PermissionChecker) {
	permission := router.Group("/permission")
	{
		permission.POST("/create", PermissionMiddleware(repo, "create_permission"), handler.CreatePermission)
		permission.GET("/all",
			middleware.AuthMiddleware(),
			// PermissionMiddleware(repo, "view_permissions"),
			handler.GetAllPermissions,
		)
		permission.GET("/:role_id",
			middleware.AuthMiddleware(),
			handler.GetPermissionByRoleId,
		)
		// permission.GET("/list", handler.ListPermissions)
		// permission.GET("/:id", handler.GetPermissionByID)
		// permission.PUT("/:id", handler.UpdatePermission)
		// permission.DELETE("/:id", handler.DeletePermission)
	}
}
