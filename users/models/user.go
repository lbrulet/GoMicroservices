package models

import (
	"time"
)

// User is a User model
type User struct {
	ID         uint `gorm:"primary_key"`
	Username   string
	Email      string
	Password   string
	IsAdmin    bool
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}
