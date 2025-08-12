package users

import (
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

	user, err := h.srv.GetByUserID(c, firebaseUid)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, utils.InternalServerErr, err)
		return
	}
	response.Success(c, http.StatusOK, "Success get user detail", user)
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	p := utils.Parse(ctx)
	data, meta, err := h.srv.GetUsers(ctx, p)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get users", err)
		return
	}

	response.SuccessWithMeta(ctx, http.StatusOK, "Success get users data", data, response.Meta{
		TotalItems:  meta.TotalItems,
		CurrentPage: meta.CurrentPage,
		PerPage:     meta.PerPage,
	})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if verrs, ok := err.(validator.ValidationErrors); ok {
			validationErrors := make([]response.ValidationError, 0, len(verrs))
			for _, fe := range verrs {
				validationErrors = append(validationErrors, response.ValidationError{
					Field:   fe.Field(),
					Message: utils.GetErrorMessage(fe),
				})
			}
			response.ValidationFailed(c, utils.InternalServerErr, validationErrors)
			return
		}

		response.Error(c, http.StatusBadRequest, utils.InvalidRequest, err.Error())
		return
	}

	resp, err := h.srv.CreateUser(c, req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, utils.InternalServerErr, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "success create new user!", resp)

}
