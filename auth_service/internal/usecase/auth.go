package usecase

import (
	"Thirawoot/shopping_cart/internal/dto"
	port_in "Thirawoot/shopping_cart/internal/ports/in"
	"Thirawoot/shopping_cart/internal/shared/mapper"
	"Thirawoot/shopping_cart/internal/shared/status"
)

type AuthUseCase interface {
	CreateAdmin(data *dto.UserCreate) mapper.ResBody
	CreateCustomer(data *dto.UserCreate) mapper.ResBody
}

type AuthUseCaseImpl struct {
	service port_in.AuthService
}

func NewAuthUseCase(s port_in.AuthService) AuthUseCase {
	return &AuthUseCaseImpl{service: s}
}

func (s *AuthUseCaseImpl) CreateAdmin(data *dto.UserCreate) mapper.ResBody {
	foundUser, _ := s.service.FindByUsername(data.Username)
	if foundUser != nil {
		return mapper.ResBody{Code: status.Conflict, Message: "error", Data: "user already exist"}
	}

	id, err := s.service.CreateAdmin(data)
	if err != nil {
		return mapper.ResBody{Code: status.BadRequest, Message: "error", Data: err.Error()}
	}

	return mapper.ResBody{Code: status.Created, Message: "success", Data: id}
}

func (s *AuthUseCaseImpl) CreateCustomer(data *dto.UserCreate) mapper.ResBody {
	foundUser, _ := s.service.FindByUsername(data.Username)
	if foundUser != nil {
		return mapper.ResBody{Code: status.Conflict, Message: "error", Data: "user already exist"}
	}

	id, err := s.service.CreateCustomer(data)
	if err != nil {
		return mapper.ResBody{Code: status.BadRequest, Message: "error", Data: err.Error()}
	}

	return mapper.ResBody{Code: status.Created, Message: "success", Data: id}
}
