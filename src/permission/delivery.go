package permission

import (
	"gobackend/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv Service
}

func NewPermissionHandler(srv Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) CreatePermission(ctx *gin.Context) {
	var input Permission
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if err := h.srv.CreatePermission(ctx, &input); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create permission", err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, "Permission created successfully", nil)
}
