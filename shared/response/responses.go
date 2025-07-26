package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	TotalItems  int `json:"total_items,omitempty"`
	CurrentPage int `json:"current_page,omitempty"`
	PerPage     int `json:"per_page,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// === Success ===
func Success(c *gin.Context, status int, message string, data any) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func SuccessWithMeta(c *gin.Context, status int, message string, data any, meta Meta) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
		"meta":    meta,
	})
}

// === Error ===
func Error(c *gin.Context, status int, message string, err any) {
	payload := gin.H{
		"status":  status,
		"message": message,
	}
	if err != nil {
		payload["error"] = err
	}
	c.JSON(status, payload)
}

func ErrorWithCode(c *gin.Context, status int, message string, err error, errorCode string) {
	payload := gin.H{
		"status":     status,
		"message":    message,
		"error_code": errorCode,
	}
	if err != nil {
		payload["error"] = err.Error()
	}
	c.JSON(status, payload)
}

func ValidationFailed(c *gin.Context, message string, errors []ValidationError) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": message,
		"errors":  errors,
	})
}
