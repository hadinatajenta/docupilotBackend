package permission

import (
	"context"
	"fmt"
	"gobackend/shared/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type service struct {
	repo PermissionChecker
}

func NewPermissionService(repo PermissionChecker) Service {
	return &service{repo: repo}
}

func PermissionMiddleware(roleRepo PermissionChecker, requiredPerm string) gin.HandlerFunc {
	return func(c *gin.Context) {

		if roleRepo == nil {
			response.Error(c, http.StatusInternalServerError, "server configuration error", nil)
			c.Abort()
			return
		}
		userID := c.GetString("user_id")
		fmt.Println("User ID from context:", userID)
		if userID == "" {
			response.Error(c, http.StatusUnauthorized, "unauthorized", nil)
			c.Abort()
			return
		}

		perms, err := roleRepo.GetPermissionsByUserID(c, userID)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "server error", err)
			c.Abort()
			return
		}

		hasPermission := false
		for _, p := range perms {
			if p.Name == requiredPerm {
				hasPermission = true
				break
			}
		}
		if !hasPermission {
			response.Error(c, http.StatusForbidden, "forbidden", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

func (s *service) CreatePermission(ctx context.Context, permission *Permission) error {
	uuid := uuid.New()
	if permission.Name == "" {
		return fmt.Errorf("permission name cannot be empty")
	}
	if permission.Description == "" {
		return fmt.Errorf("permission description cannot be empty")
	}

	insert := &Permission{
		ID:          uuid.String(),
		Name:        permission.Name,
		Description: permission.Description,
		CreatedAt:   time.Now(),
	}
	if err := s.repo.CreatePermission(ctx, insert); err != nil {
		return fmt.Errorf("failed to create permission: %w", err)
	}
	return nil
}
