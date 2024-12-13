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

func (r *UserRoleRepositoryImpl) FindRole(id uint) (*domain.UserRole, error) {
	var role = domain.UserRole{ID: id}

	result := r.db.First(&role)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("ROLE_NOT_FOUND")
	}

	return &role, nil
}

func (r *UserRoleRepositoryImpl) FindRoles() ([]domain.UserRole, error, int64) {
	var roles []domain.UserRole

	result := r.db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error, result.RowsAffected
	}

	return roles, nil, result.RowsAffected
}

func (r *UserRoleRepositoryImpl) DeleteRole(id uint) error {
	var role = domain.UserRole{ID: id}

	res := r.db.Delete(&role)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
