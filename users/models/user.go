package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username   string
	Email      string
	Password   string
	IsAdmin    bool
	IsVerified bool
}
