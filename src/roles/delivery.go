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

func (h *Handler) AssignPermissions(c *gin.Context) {
	roleID := c.Param("id")
	var req assignPermissionsReq

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request", err)
		return
	}

	res, err := h.srv.AssignPermissionsToRole(c, roleID, req.Permissions)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "permissions assigned successfully", res)
}

func (h *Handler) GetAllRoles(ctx *gin.Context) {
	roles, err := h.srv.GetAllRoles(ctx)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to retrieve roles", err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "Roles retrieved successfully", roles)
}
