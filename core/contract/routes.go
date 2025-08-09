package contract

import (
	"gobackend/app"
	"gobackend/src/auth"
	"gobackend/src/menus"
	"gobackend/src/roles"
	"gobackend/src/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, deps *app.Dependencies) {
	routes := rg.Group("/v1")
	{
		// package roles
		roleRepo := roles.NewRoleRepository(deps.DB)

		// package users
		userRepo := users.NewUserRepository(deps.DB)
		userService := users.NewUserService(userRepo, deps.DB)
		userHandler := users.NewUserHandler(userService)
		users.RegisterUserRoutes(routes, userHandler, roleRepo)

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
