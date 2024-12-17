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

func (s *UserRoleUseCase) FindRole(id uint) mapper.ResBody {
	res, err := s.service.FindRole(id)
	if err != nil {
		return mapper.ResBody{Code: status.NotFound, Message: "error", Data: err.Error()}
	}

	return mapper.ResBody{Code: status.OK, Message: "success", Data: res}
}

func (s UserRoleUseCase) FindRoles() mapper.ResBody {
	res, err := s.service.FindRoles()
	if err != nil {
		return mapper.ResBody{Code: status.InternalError, Message: "error", Data: err.Error()}
	}
	if len(*res) == 0 {
		return mapper.ResBody{Code: status.NotFound, Message: "not found", Data: *res}
	}

	return mapper.ResBody{Code: status.OK, Message: "success", Data: *res}
}

func (s *UserRoleUseCase) DeleteRole(id uint) mapper.ResBody {
	_, err := s.service.FindRole(id)
	if err != nil {
		return mapper.ResBody{Code: status.NotFound, Message: "error", Data: err.Error()}
	}

	data, err := s.service.DeleteRole(id)
	if err != nil {
		return mapper.ResBody{Code: status.BadRequest, Message: "error", Data: err.Error()}
	}

	return mapper.ResBody{Code: status.OK, Message: "success", Data: data}
}
