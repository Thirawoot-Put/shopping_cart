package port_out

import "Thirawoot/shopping_cart/internal/domain"

type UserRoleRepository interface {
	Create(data domain.UserRole) error
	GetById(id uint) (*domain.UserRole, error)
}
