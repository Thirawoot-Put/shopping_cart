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

func (s *UserRoleServiceImpl) Create(data *dto.UserRoleCreate) string {
	newRole := domain.UserRole{Role: data.Role}

	err := s.repo.Create(newRole)
	if err != nil {
		err := fmt.Errorf("Failed to create user role: %w", err)
		return err.Error()
	}

	return "Create user role success"
}

func (s *UserRoleServiceImpl) FindRole(id uint) (*dto.UserRoleResponse, error) {
	role, err := s.repo.FindRole(id)
	if err != nil {
		return nil, err
	}

	res := dto.UserRoleResponse{ID: role.ID, Role: role.Role}

	return &res, nil
}

func (s *UserRoleServiceImpl) FindRoles() (*[]dto.UserRoleResponse, error) {
	roles, err := s.repo.FindRoles()
	if err != nil {
		return nil, err
	}

	res := make([]dto.UserRoleResponse, len(roles))
	for idx, role := range roles {
		res[idx] = dto.UserRoleResponse{ID: role.ID, Role: role.Role}
	}

	return &res, nil
}

func (s *UserRoleServiceImpl) DeleteRole(id uint) (string, error) {
	err := s.repo.DeleteRole(id)
	if err != nil {
		return "", err
	}

	return "Delete success", nil
}
