package port_out

import (
	"Thirawoot/shopping_cart/internal/domain"
)

type AuthRepository interface {
	Create(data *domain.User) (*domain.User, error)
}
