package roles

import (
	"context"
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
