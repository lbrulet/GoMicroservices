package postgreSQL

import (
	"github.com/jinzhu/gorm"
	"github.com/lbrulet/GoMicroservices/users-service/database"
	"github.com/lbrulet/GoMicroservices/users-service/models"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type repository struct {
	Database *gorm.DB
}

// PostgresConnexion is used to connect to PgSQL and return a DB struct
func PostgresConnexion() (*gorm.DB, error) {
	// db, err := gorm.Open("postgres", "host=postgreSQL port=5432 user=admin dbname=api_micro password=admin123 sslmode=disable")
	db, err := gorm.Open("sqlite3", ":memory:")
	return db, err
}

// NewPostgresRepository used to create a new interface Repository
func NewPostgresRepository(db *gorm.DB) database.Repository {
	return &repository{
		Database: db,
	}
}

// CountByUsernameAndEmail is used to count by username or email
func (r *repository) CountByUsernameAndEmail(username, email string) (int) {
	count := 0
	r.Database.Table("users").Where("username = ?", username).Or("email = ?", email).Count(&count)
	return count
}

// GetByUsername is used to get a user by username
func (r *repository) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	err := r.Database.Where("username = ?", username).First(&user).Error
	return user, err
}

// CreateUser is used to create a user
func (r *repository) CreateUser(user models.User) {
	r.Database.Create(&user)
}

func (r *repository) ResetDatabase() {
	r.Database.DropTableIfExists("users")
}

func (r *repository) MigrateDatabase() {
	r.Database.AutoMigrate(&models.User{})
}