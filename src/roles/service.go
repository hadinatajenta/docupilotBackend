package roles

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

func NewRoleService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateRole(ctx context.Context, role *Role) error {
	uuid := uuid.New()

	checkExist, err := s.repo.CheckRoleExist(ctx, role.Name)
	if err != nil {
		return err
	}

	if checkExist {
		return fmt.Errorf("role %s already exists", role.Name)
	}

	roleInput := &Role{
		ID:          uuid.String(),
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, roleInput); err != nil {
		return err
	}

	return nil
}

func (s *service) AssignPermissionsToRole(ctx context.Context, roleID string, permissionIDs []string) (RolePermissionRes, error) {
	role, err := s.repo.GetByID(ctx, roleID)
	if err != nil {
		return RolePermissionRes{}, fmt.Errorf("role not found: %w", err)
	}
	if role == nil {
		return RolePermissionRes{}, fmt.Errorf("role not found")
	}

	exists, err := s.repo.PermissionsExist(ctx, permissionIDs)
	if err != nil {
		return RolePermissionRes{}, err
	}
	if !exists {
		return RolePermissionRes{}, fmt.Errorf("one or more permissions not found")
	}

	if err := s.repo.AssignPermissions(ctx, roleID, permissionIDs); err != nil {
		return RolePermissionRes{}, err
	}

	getNameIds, err := s.repo.GetPermissionByIDS(ctx, permissionIDs)
	if err != nil {
		return RolePermissionRes{}, fmt.Errorf("failed to get permission names: %w", err)
	}

	permissions := make([]any, len(permissionIDs))
	for i, id := range permissionIDs {
		permissions[i] = map[string]string{
			"id":   id,
			"name": getNameIds[i].Name,
		}
	}

	res := &RolePermissionRes{
		RoleID:      role.ID,
		RoleName:    role.Name,
		Permissions: permissions,
	}

	return *res, nil
}

func (s *service) GetAllRoles(ctx context.Context) ([]Role, error) {
	roles, err := s.repo.GetAllRoles(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get roles: %w", err)
	}
	return roles, nil
}
