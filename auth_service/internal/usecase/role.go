package usecase

import (
	"Thirawoot/shopping_cart/internal/dto"
	port_in "Thirawoot/shopping_cart/internal/ports/in"
	"Thirawoot/shopping_cart/internal/shared/mapper"
	"Thirawoot/shopping_cart/internal/shared/status"
)

type UserRoleUseCase interface {
	CreateRole(data *dto.UserRoleCreate) mapper.ResBody
	FindRole(id uint) mapper.ResBody
	FindRoles() mapper.ResBody
	DeleteRole(id uint) mapper.ResBody
}

type UserRoleUseCaseImpl struct {
	service port_in.UserRoleService
}

func NewUserRoleUseCasee(s port_in.UserRoleService) UserRoleUseCase {
	return &UserRoleUseCaseImpl{service: s}
}

func (s *UserRoleUseCaseImpl) CreateRole(data *dto.UserRoleCreate) mapper.ResBody {
	res := s.service.Create(data)

	return mapper.ResBody{Code: status.Created, Message: "success", Data: res}
}

func (s *UserRoleUseCaseImpl) FindRole(id uint) mapper.ResBody {
	res, err := s.service.FindRole(id)
	if err != nil {
		return mapper.ResBody{Code: status.NotFound, Message: "error", Data: err.Error()}
	}

	return mapper.ResBody{Code: status.OK, Message: "success", Data: res}
}

func (s UserRoleUseCaseImpl) FindRoles() mapper.ResBody {
	res, err := s.service.FindRoles()
	if err != nil {
		return mapper.ResBody{Code: status.InternalError, Message: "error", Data: err.Error()}
	}
	if len(*res) == 0 {
		return mapper.ResBody{Code: status.NotFound, Message: "not found", Data: *res}
	}

	return mapper.ResBody{Code: status.OK, Message: "success", Data: *res}
}

func (s *UserRoleUseCaseImpl) DeleteRole(id uint) mapper.ResBody {
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
