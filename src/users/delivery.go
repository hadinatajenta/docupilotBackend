package users

import (
	"gobackend/shared/response"
	"gobackend/shared/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv UService
}

func NewUserHandler(srv UService) *Handler {
	return &Handler{srv}
}

// Sync user data dari Firebase
// func (h *Handler) SyncUser(c *gin.Context) {
// 	uid := c.GetString("firebase_uid")
// 	email := c.GetString("firebase_email")
// 	name := c.GetString("firebase_name")
// 	avatar := c.GetString("firebase_avatar")

// 	user, err := h.srv.SyncFirebaseUser(c, uid, email, name, avatar)
// 	if err != nil {
// 		response.ErrorWithCode(c, http.StatusInternalServerError, "error from server side", err, "500-ISERR")
// 		return
// 	}
// 	response.Success(c, http.StatusCreated, "Success Create user data", user)
// }

func (h *Handler) GetDetailByFirebaseUID(c *gin.Context) {
	firebaseUid := c.GetString("user_id")
	if firebaseUid == "" {
		response.Error(c, http.StatusBadRequest, "tidak menemukan user_id", nil)
		return
	}

	user, err := h.srv.GetByFirebaseUID(c, firebaseUid)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, utils.InternalServerErr, err)
		return
	}
	response.Success(c, http.StatusOK, "Success get user detail", user)
}
