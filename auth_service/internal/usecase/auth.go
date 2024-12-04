package usecase

import (
	port_in "Thirawoot/shopping_cart/internal/ports/in"
	"fmt"
)

type AuthUseCase struct {
	service port_in.AuthService
}

func NewAuthUseCase(s port_in.AuthService) *AuthUseCase {
	return &AuthUseCase{service: s}
}

func (s *AuthUseCase) CreateUser() {
	fmt.Println("-=-> hello from usecase")
	s.service.Create()
}
