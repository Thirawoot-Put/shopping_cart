package port_in

import "Thirawoot/shopping_cart/internal/dto"

type UserRoleService interface {
	Create(data *dto.UserRoleCreate) string
	FindRole(id uint) (*dto.UserRoleResponse, error)
	DeleteRole(id uint) (string, error)
}
