package port_out

import (
	"Thirawoot/shopping_cart/internal/domain"
)

type AuthRepository interface {
	Create(data *domain.User) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
}
