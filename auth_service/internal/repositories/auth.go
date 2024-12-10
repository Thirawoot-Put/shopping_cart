package repositories

import (
	"Thirawoot/shopping_cart/internal/domain"
	"Thirawoot/shopping_cart/internal/dto"
	port_out "Thirawoot/shopping_cart/internal/ports/out"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) port_out.AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) Create(data dto.UserCreate) error {
	user := domain.User{Username: data.Username, Password: data.Password, RoleId: data.RoleId}
	result := r.db.Create(&user)
	return result.Error
}
