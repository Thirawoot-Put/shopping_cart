package usecase

import (
	"Thirawoot/shopping_cart/internal/dto"
	port_in "Thirawoot/shopping_cart/internal/ports/in"
	"Thirawoot/shopping_cart/internal/shared"
	"Thirawoot/shopping_cart/internal/shared/status"
)

type UserRoleUseCase struct {
	service port_in.UserRoleService
}

func NewUserRoleUseCasee(s port_in.UserRoleService) *UserRoleUseCase {
	return &UserRoleUseCase{service: s}
}

func (s *UserRoleUseCase) CreateRole(data *dto.UserRoleCreate) shared.Response {
	res := s.service.Create(data)

	return shared.Response{Code: status.Created, Message: "success", Data: res}
}
