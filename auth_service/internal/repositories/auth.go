package repositories

import (
	port_out "Thirawoot/shopping_cart/internal/ports/out"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) port_out.AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) Create() error {
	fmt.Println("==> hello from repositories")
	return nil
}
