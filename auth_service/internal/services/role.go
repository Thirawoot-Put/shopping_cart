package services

import (
	"Thirawoot/shopping_cart/internal/domain"
	"Thirawoot/shopping_cart/internal/dto"
	port_in "Thirawoot/shopping_cart/internal/ports/in"
	port_out "Thirawoot/shopping_cart/internal/ports/out"
	"fmt"
)

type UserRoleServiceImpl struct {
	repo port_out.UserRoleRepository
}

func NewUserRoleService(r port_out.UserRoleRepository) port_in.UserRoleService {
	return &UserRoleServiceImpl{repo: r}
}

func (s *UserRoleServiceImpl) Create(data dto.UserRoleCreate) string {
	newRole := domain.UserRole{Role: data.Role}

	err := s.repo.Create(newRole)
	if err != nil {
		fmt.Errorf("Failed to create user role: %w", err)
	}

	return "Create user role success"
}
