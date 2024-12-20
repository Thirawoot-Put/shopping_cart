package services

import (
	"Thirawoot/shopping_cart/internal/domain"
	"Thirawoot/shopping_cart/internal/dto"
	port_in "Thirawoot/shopping_cart/internal/ports/in"
	port_out "Thirawoot/shopping_cart/internal/ports/out"
	jwtservice "Thirawoot/shopping_cart/internal/shared/jwtService"
	"Thirawoot/shopping_cart/internal/shared/status"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repo port_out.AuthRepository
}

func NewAuthService(r port_out.AuthRepository) port_in.AuthService {
	return &AuthServiceImpl{repo: r}
}

func (s *AuthServiceImpl) CreateAdmin(data *dto.UserCreate) (*uint, error) {
	if data.Username == "" || data.Password == "" {
		return nil, fmt.Errorf("username and password are required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate hashpassword")
	}

	newUser := domain.User{Username: data.Username, Password: string(hash), RoleId: 16}

	user, err := s.repo.Create(&newUser)
	if err != nil {
		return nil, fmt.Errorf("Failed to write user: %w", err)
	}

	return &user.ID, nil
}

func (s *AuthServiceImpl) CreateCustomer(data *dto.UserCreate) (*uint, error) {
	if data.Username == "" || data.Password == "" {
		return nil, fmt.Errorf("username and password are required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate hashpassword")
	}

	newUser := domain.User{Username: data.Username, Password: string(hash), RoleId: 17}

	user, err := s.repo.Create(&newUser)
	if err != nil {
		return nil, fmt.Errorf("Failed to write user: %w", err)
	}

	return &user.ID, nil
}

func (s *AuthServiceImpl) FindByUsername(username string) (*dto.UserResponse, error) {
	if username == "" {
		return nil, fmt.Errorf("username are required")
	}

	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{ID: user.ID, Username: user.Username}, nil
}

func (s *AuthServiceImpl) AuthUsername(data *dto.UserLogin) (string, error, uint) {
	if data.Username == "" || data.Password == "" {
		return "", fmt.Errorf("username and password are required"), status.BadRequest
	}

	user, err := s.repo.FindByUsername(data.Username)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err), status.NotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return "", fmt.Errorf("password match error: %w", err), status.BadRequest
	}

	tokenString, err := jwtservice.GenerateToken(user.ID, user.UserRole.Role)
	if err != nil {
		return "", err, status.BadRequest
	}

	return tokenString, nil, status.OK
}
