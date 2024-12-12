package domain

import (
	"time"
)

type UserRole struct {
	ID   uint   `gorm:"primaryKey;autoincrement"`
	Role string `gorm:"type:varchar(255);unique;not null"`

	User []User `gorm:"foreignKey:RoleId"`
}

type User struct {
	ID       uint   `gorm:"primaryKey;autoincrement"`
	Username string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`

	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`

	DeletedAt bool `gorm:"default:false"`

	RoleId uint `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
