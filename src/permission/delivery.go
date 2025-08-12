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

func (h *Handler) GetAllPermissions(ctx *gin.Context) {
	permissions, err := h.srv.GetAllPermissions(ctx)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to retrieve permissions", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "Permissions retrieved successfully", permissions)
}

func (h *Handler) GetPermissionByRoleId(ctx *gin.Context) {
	role_id := ctx.Param("role_id")

	data, err := h.srv.GetPermissionByRoleId(ctx, role_id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to retrieve permissions", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success get Permission by role_id", data)
}
