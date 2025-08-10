package contract

import (
	"gobackend/app"
	"gobackend/src/auth"
	"gobackend/src/menus"
	"gobackend/src/permission"
	"gobackend/src/roles"
	"gobackend/src/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, deps *app.Dependencies) {
	routes := rg.Group("/v1")
	{
		// package permissions
		permissionRepo := permission.NewPermissionRepository((deps.DB))
		permissionService := permission.NewPermissionService(permissionRepo)
		permissionHandler := permission.NewPermissionHandler(permissionService)
		permission.RegisterPermissionRoutes(routes, permissionHandler, permissionRepo)

		// package roles
		roleRepo := roles.NewRoleRepository(deps.DB)
		roleService := roles.NewRoleService(roleRepo)
		roleHandler := roles.NewRoleHandler(roleService)
		roles.RegisRoleRoutes(routes, roleHandler, permissionRepo)

		// package users
		userRepo := users.NewUserRepository(deps.DB)
		userService := users.NewUserService(userRepo, deps.DB)
		userHandler := users.NewUserHandler(userService)
		users.RegisterUserRoutes(routes, userHandler, permissionRepo)

		// package menus
		menuRepo := menus.NewMenuRepository(deps.DB)
		menuService := menus.NewMenuService(menuRepo)
		menuHandler := menus.NewMenusHandler(menuService)
		menus.RegisMenuRoutes(routes, menuHandler)

		// package auth
		authRepo := auth.NewAuthRepository(deps.DB)
		authService := auth.NewAuthService(authRepo, userRepo, deps.DB)
		authHandler := auth.NewAuthHandler(authService)
		auth.RegisterAuthRoutes(routes, authHandler)
	}

}
