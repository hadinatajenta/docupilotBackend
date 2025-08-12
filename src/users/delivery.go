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
	u, r := h.srv.GetUsers(ctx)
	if r != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get users", r)
		return
	}

	response.Success(ctx, http.StatusOK, "Success get users data", u)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if verrs, ok := err.(validator.ValidationErrors); ok {
			out := make(map[string]string)
			for _, fe := range verrs {
				field := fe.Field()
				switch fe.Tag() {
				case "required":
					out[field] = "is required"
				case "email":
					out[field] = "must be a valid email"
				case "min":
					out[field] = "must be at least " + fe.Param() + " characters"
				default:
					out[field] = "is invalid"
				}
			}
			response.Error(c, http.StatusBadRequest, utils.InvalidRequest, out)
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
