package usecase

import (
	"Thirawoot/shopping_cart/internal/dto"
	port_in "Thirawoot/shopping_cart/internal/ports/in"
	"Thirawoot/shopping_cart/internal/shared/mapper"
	"Thirawoot/shopping_cart/internal/shared/status"
)

type UserRoleUseCase struct {
	service port_in.UserRoleService
}

func NewUserRoleUseCasee(s port_in.UserRoleService) *UserRoleUseCase {
	return &UserRoleUseCase{service: s}
}

func (s *UserRoleUseCase) CreateRole(data *dto.UserRoleCreate) mapper.ResBody {
	res := s.service.Create(data)

	return mapper.ResBody{Code: status.Created, Message: "success", Data: res}
}

func (s *UserRoleUseCase) GetRoleById(id uint) mapper.ResBody {
	res, err := s.service.GetById(id)
	if err != nil {
		return mapper.ResBody{Code: status.NotFound, Message: "error", Data: err.Error()}
	}

	return mapper.ResBody{Code: status.OK, Message: "success", Data: res}
}
