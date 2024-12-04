package domain

import "time"

type UserRole struct {
	ID   uint   `gorm:"primaryKey;autoincrement"`
	Role string `gorm:"type:vachar(255);unique;not null"`
}

type User struct {
	ID       uint   `gorm:"primaryKey;autoincrement"`
	Username string `gorm:"type:vachar(255);unique;not null"`
	Password string `gorm:"type:vachar(255);not null"`

	Role UserRole

	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null"`

	IsDelete bool `gorm:"default:false"`
}
