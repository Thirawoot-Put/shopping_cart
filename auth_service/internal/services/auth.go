package services

import (
	port_in "Thirawoot/shopping_cart/internal/ports/in"
	port_out "Thirawoot/shopping_cart/internal/ports/out"
	"fmt"
)

type AuthServiceImpl struct {
	repo port_out.AuthRepository
}

func NewAuthService(r port_out.AuthRepository) port_in.AuthService {
	return &AuthServiceImpl{repo: r}
}

func (s *AuthServiceImpl) Create() {
	fmt.Println("--> hello from services")
}
