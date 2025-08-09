package menus

import (
	"context"
	"errors"
)

type menuService struct {
	repo Repository
}

func NewMenuService(repo Repository) Service {
	return &menuService{repo}
}

func (s *menuService) GetMenuByRole(ctx context.Context, role string) ([]Menus, error) {
	if role == "" {
		return nil, errors.New("role kosong")
	}

	menus, err := s.repo.GetMenuByRole(ctx, role)
	if err != nil {
		return nil, err
	}

	return menus, nil
}
