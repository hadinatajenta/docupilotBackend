package contract

import (
	"gobackend/app"
	"gobackend/src/menus"
	"gobackend/src/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, deps *app.Dependencies) {
	routes := rg.Group("/v1")
	{
		userRepo := users.NewUserRepository(deps.DB)
		userService := users.NewUserService(userRepo, deps.DB)
		userHandler := users.NewUserHandler(userService)
		users.RegisUserRoute(routes, userHandler)

		menuRepo := menus.NewMenuRepository(deps.DB)
		menuService := menus.NewMenuService(menuRepo)
		menuHandler := menus.NewMenusHandler(menuService)
		menus.RegisMenuRoutes(routes, menuHandler)
	}
}
