package services

import (
	"Thirawoot/shopping_cart/internal/domain"
	"Thirawoot/shopping_cart/internal/dto"
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

func (s *AuthServiceImpl) CreateAdmin(data dto.UserCreate) (*uint, error) {
	if data.Username == "" || data.Password == "" {
		return nil, fmt.Errorf("username and password are required")
	}
	newUser := domain.User{Username: data.Username, Password: data.Password, RoleId: 16}

	user, err := s.repo.Create(&newUser)
	if err != nil {
		return nil, fmt.Errorf("Failed to write user: %w", err)
	}

	return &user.ID, nil
}

func (s *AuthServiceImpl) CreateCustomer(data dto.UserCreate) error {
	fmt.Println("--> hello from services")
	return nil
}
