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
