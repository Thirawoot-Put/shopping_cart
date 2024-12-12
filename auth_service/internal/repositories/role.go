package repositories

import (
	"Thirawoot/shopping_cart/internal/domain"
	port_out "Thirawoot/shopping_cart/internal/ports/out"
	"errors"

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

func (r *UserRoleRepositoryImpl) GetById(id uint) (*domain.UserRole, error) {
	var role = domain.UserRole{ID: id}

	result := r.db.First(&role)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("ROLE_NOT_FOUND")
	}

	return &role, nil
}

func (r *UserRoleRepositoryImpl) Delete(id uint) {

}
