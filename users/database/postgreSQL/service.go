package postgreSQL

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lbrulet/GoMicroservices/users/database"
	"github.com/lbrulet/GoMicroservices/users/models"
)

type repository struct {
	Database *gorm.DB
}

// PostgresConnexion is used to connect to PgSQL and return a DB struct
func PostgresConnexion() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=postgreSQL port=5432 user=admin dbname=api_micro password=admin123 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, nil
}

// NewPostgresRepository used to create a new interface Repository
func NewPostgresRepository(db *gorm.DB) database.Repository {
	return &repository{
		Database: db,
	}
}

// CountByUsernameAndEmail is used to count by username or email
func (r *repository) CountByUsernameAndEmail(username, email string) (int, error) {
	count := 0
	if err := r.Database.Table("users").Where("username = ?", username).Or("email = ?", email).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetByUsername is used to get a user by username
func (r *repository) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	if notFound := r.Database.Where("username = ?", username).First(&user).RecordNotFound(); notFound {
		return models.User{}, errors.New("user does not exist")
	}
	return user, nil
}

// CreateUser is used to create a user
func (r *repository) CreateUser(user models.User) error {
	if err := r.Database.Create(&user).Error; err != nil {
		return err
	}
	return nil
}




