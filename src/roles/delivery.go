package roles

import (
	"gobackend/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv Service
}

func NewRoleHandler(srv Service) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) CreateRole(ctx *gin.Context) {
	var input Role
	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid input", err.Error())
		return
	}

	if err := h.srv.CreateRole(ctx, &input); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create role", err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, " Role created successfully", nil)
}
