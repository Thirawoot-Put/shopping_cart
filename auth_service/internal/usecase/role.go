package usecase

import (
	"Thirawoot/shopping_cart/internal/dto"
	port_in "Thirawoot/shopping_cart/internal/ports/in"
)

type UserRoleUseCase struct {
	service port_in.UserRoleService
}

func NewUserRoleUseCasee(s port_in.UserRoleService) *UserRoleUseCase {
	return &UserRoleUseCase{service: s}
}

func (s *UserRoleUseCase) CreateRole(data *dto.UserRoleCreate) string {
	res := s.service.Create(data)
	return res
}
