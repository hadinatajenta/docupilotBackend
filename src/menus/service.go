package menus

type menuService struct {
	repo Repository
}

func NewMenuService(repo Repository) *menuService {
	return &menuService{repo}
}
