package menus

import "context"

type Repository interface {
	GetMenuByRole(ctx context.Context, role string) ([]Menus, error)
}

type Service interface {
	GetMenuByRole(ctx context.Context, role string) ([]Menus, error)
}
