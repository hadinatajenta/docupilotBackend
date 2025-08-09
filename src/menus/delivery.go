package menus

import (
	"gobackend/shared/response"
	"gobackend/shared/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv Service
}

func NewMenusHandler(srv Service) *Handler {
	return &Handler{srv}
}

func (h *Handler) GetMenus(ctx *gin.Context) {
	userRoles := ctx.GetHeader("User-roles")

	menus, err := h.srv.GetMenuByRole(ctx, userRoles)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, utils.InternalServerErr, err)
		return
	}

	response.Success(ctx, http.StatusOK, "berhasil mendapatkan menu", menus)
}
