package port_out

import "Thirawoot/shopping_cart/internal/dto"

type AuthRepository interface {
	Create(data dto.UserCreate) error
}
