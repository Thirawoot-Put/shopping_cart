package port_in

import "Thirawoot/shopping_cart/internal/dto"

type UserRoleService interface {
	Create(data *dto.UserRoleCreate) string
	GetById(id uint) (*dto.UserRoleResponse, error)
}
