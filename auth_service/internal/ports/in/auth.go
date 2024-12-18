package port_in

import "Thirawoot/shopping_cart/internal/dto"

type AuthService interface {
	CreateAdmin(data dto.UserCreate) (*uint, error)
	CreateCustomer(data dto.UserCreate) error
}
