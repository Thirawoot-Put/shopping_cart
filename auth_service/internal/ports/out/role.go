package port_out

import "Thirawoot/shopping_cart/internal/domain"

type UserRoleRepository interface {
	Create(data domain.UserRole) error
	FindRole(id uint) (*domain.UserRole, error)
	FindRoles() (*[]domain.UserRole, error, int64)
	DeleteRole(id uint) error
}
