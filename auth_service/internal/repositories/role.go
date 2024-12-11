package repositories

import (
	"Thirawoot/shopping_cart/internal/domain"
	port_out "Thirawoot/shopping_cart/internal/ports/out"

	"gorm.io/gorm"
)

type UserRoleRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) port_out.UserRoleRepository {
	return &UserRoleRepositoryImpl{db: db}
}

func (r *UserRoleRepositoryImpl) Create(data domain.UserRole) error {
	role := domain.UserRole{Role: data.Role}
	res := r.db.Create(&role)

	return res.Error
}
