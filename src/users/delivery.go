package users

import (
	"errors"
	"gobackend/shared/response"
	"gobackend/shared/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	srv UService
}

func NewUserHandler(srv UService) *Handler {
	return &Handler{srv}
}

func (h *Handler) SyncUser(c *gin.Context) {
	uid := c.GetString("firebase_uid")
	email := c.GetString("firebase_email")
	name := c.GetString("firebase_name")
	avatar := c.GetString("firebase_avatar")

	user, err := h.srv.SyncFirebaseUser(c, uid, email, name, avatar)
	if err != nil {
		response.ErrorWithCode(c, http.StatusInternalServerError, "error from server side", err, "500-ISERR")
		return
	}

	response.Success(c, http.StatusCreated, "Success Create user data", user)
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request", err)
		return
	}

	token, err := h.srv.Login(c, req.Email, req.Password)
	if err != nil {
		if err.Error() == "invalid credentials" {
			response.Error(c, http.StatusUnauthorized, "invalid credentials", err)
			return
		}
		response.ErrorWithCode(c, http.StatusInternalServerError, "terjadi kesalahan dari sisi server", err, "SVRERR500")
		return
	}

	response.Success(c, http.StatusOK, "Login success", token)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var req TokenReq

	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			msg := ve[0].Field() + utils.IsRequired
			response.Error(c, http.StatusBadRequest, utils.InvalidRequest, msg)
			return
		}
		response.Error(c, http.StatusBadRequest, utils.InvalidRequest, err)
		return
	}

	newToken, err := h.srv.RefreshToken(c, req.RefreshToken)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, utils.InternalServerErr, err)
		return
	}
	response.Success(c, http.StatusOK, "refresh success", newToken)
}

func (h *Handler) Logout(c *gin.Context) {
	var req TokenReq

	if err := c.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			msg := ve[0].Field() + utils.IsRequired
			response.Error(c, http.StatusBadRequest, utils.InvalidRequest, msg)
			return
		}
		response.Error(c, http.StatusBadRequest, utils.InvalidRequest, err)
		return
	}

	if err := h.srv.Logout(c, req.RefreshToken); err != nil {
		response.Error(c, http.StatusInternalServerError, utils.InternalServerErr, err)
		return
	}

	response.Success(c, http.StatusOK, "Logout Success", nil)

}
