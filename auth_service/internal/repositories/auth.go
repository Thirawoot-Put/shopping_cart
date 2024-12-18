package repositories

import (
	"Thirawoot/shopping_cart/internal/domain"
	port_out "Thirawoot/shopping_cart/internal/ports/out"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) port_out.AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) Create(data *domain.User) (*domain.User, error) {
	user := domain.User{Username: data.Username, Password: data.Password, RoleId: data.RoleId}
	result := r.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *AuthRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	result := r.db.Preload("UserRole").Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
