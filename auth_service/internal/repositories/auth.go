package repositories

import (
	port_out "Thirawoot/shopping_cart/internal/ports/out"
	"fmt"
)

type AuthRepositoryImpl struct{}

func NewAuthRepository() port_out.AuthRepository {
	return &AuthRepositoryImpl{}
}

func (r *AuthRepositoryImpl) Create() error {
	fmt.Println("==> hello from repositories")
	return nil
}
