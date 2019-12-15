package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lbrulet/GoMicroservices/users/models"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=postgreSQL port=5432 user=admin dbname=api_micro password=admin123 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, nil
}