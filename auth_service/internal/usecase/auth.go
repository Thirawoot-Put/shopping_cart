package usecase

import (
	"Thirawoot/shopping_cart/internal/dto"
	port_in "Thirawoot/shopping_cart/internal/ports/in"
	"Thirawoot/shopping_cart/internal/shared/mapper"
	"Thirawoot/shopping_cart/internal/shared/status"
)

type AuthUseCase struct {
	service port_in.AuthService
}

func NewAuthUseCase(s port_in.AuthService) *AuthUseCase {
	return &AuthUseCase{service: s}
}

func (s *AuthUseCase) CreateAdmin(data *dto.UserCreate) mapper.ResBody {
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

func (s *AuthUseCase) CreateCustomer(data *dto.UserCreate) mapper.ResBody {
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

func (s *AuthUseCase) LoginUser(data *dto.UserLogin) mapper.ResBody {
	tokenString, err, code := s.service.AuthUsername(data)
	if err != nil {
		return mapper.ResBody{Code: code, Message: "error", Data: err}
	}

	return mapper.ResBody{Code: code, Message: "success", Data: tokenString}
}
