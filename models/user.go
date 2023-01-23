package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model

	ID         int
	ExternalID string `gorm:"unique;not null;"`

	Username string `gorm:"unique;not null;"`
	Email    string `gorm:"unique;not null;"`
	Password string `gorm:"not null;"`

	Role string `gorm:"not null;size:255;default:'default'"`

	IsActive bool `gorm:"not null;default:true"`
	IsAdmin  bool `gorm:"not null;default:false"`
	IsStaff  bool `gorm:"not null;default:false"`
}
