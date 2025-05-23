package port_in

import "Thirawoot/shopping_cart/internal/dto"

type AuthService interface {
	CreateAdmin(data *dto.UserCreate) (*uint, error)
	CreateCustomer(data *dto.UserCreate) (*uint, error)
	FindByUsername(username string) (*dto.UserResponse, error)
	AuthUsername(data *dto.UserLogin) (string, error, int)
}
