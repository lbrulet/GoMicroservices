package database

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lbrulet/GoMicroservices/users-service/models"
)

// Repository is used to create an interface between all sort of database
type Repository interface {
	CountByUsernameAndEmail(username, email string) (int, error)
	GetByUsername(username string) (models.User, error)
	CreateUser(user models.User) error
}