package menus

import "github.com/gin-gonic/gin"

func RegisMenuRoutes(rg *gin.RouterGroup, h *Handler) {
	menu := rg.Group("/menu")
	{
		menu.GET("/get-list-menu", h.GetMenus)
	}
}
