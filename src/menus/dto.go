package menus

type Menus struct {
	ID       int8   `json:"id"`
	MenuName string `json:"text" gorm:"column:text"`
	Icon     string `json:"icon"`
	Path     string `json:"path"`
}

type UserRoles struct {
	UserRoles string `json:"user_roles" binding:"required"`
}
