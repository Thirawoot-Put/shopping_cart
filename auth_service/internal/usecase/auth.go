package usecase

import (
	"Thirawoot/shopping_cart/internal/services"
	"fmt"
)

type AuthUseCase struct {
	service *services.AuthServiceImpl
}

func NewAuthUseCase(s *services.AuthServiceImpl) *AuthUseCase {
	return &AuthUseCase{service: s}
}

func (s *AuthUseCase) CreateUser() {
	fmt.Println("-=-> hello from usecase")
	s.service.Create()
}
