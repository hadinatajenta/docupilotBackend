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

func (s *service) GetAllPermissions(ctx context.Context) (PermissionResponse, error) {
	permissions, err := s.repo.GetAllPermissions(ctx)
	if err != nil {
		return PermissionResponse{}, fmt.Errorf("failed to get permissions: %w", err)
	}

	resp := PermissionResponse{
		Data: GroupPermissionsByFeature(permissions),
	}
	return resp, nil
}

func GroupPermissionsByFeature(perms []Permission) map[string][]Permission {
	result := make(map[string][]Permission)
	for _, p := range perms {
		result[p.Feature] = append(result[p.Feature], p)
	}
	return result
}

func (s *service) GetPermissionByRoleId(ctx context.Context, roleId string) (GetPermissionByUidRes, error) {
	if roleId == "" {
		return GetPermissionByUidRes{}, fmt.Errorf("user_id cannot be null")
	}

	getPermissions, err := s.repo.GetPermissionByRoleId(ctx, roleId)
	if err != nil {
		return GetPermissionByUidRes{}, err
	}

	resp := GetPermissionByUidRes{
		Role:       roleId,
		Permission: getPermissions,
	}

	return resp, nil
}
