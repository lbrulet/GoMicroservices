package handler

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/lbrulet/GoMicroservices/users/models"
	"github.com/micro/go-micro/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	users "github.com/lbrulet/GoMicroservices/users/proto/users"
)

type Users struct {
	ServiceName string
	Logger      *logrus.Entry
	Database    *gorm.DB
}

func (e *Users) Create(ctx context.Context, req *users.CreateRequest, rsp *users.UserResponse) error {
	e.Logger.Info("Received users.Create")

	count := 0
	e.Database.Table("users").Where("username = ?", req.Username).Or("email = ?", req.Email).Count(&count)

	if count != 0 {
		e.Logger.Infof("User [%s, %s] already exist", req.Email, req.Username)
		return errors.Conflict(e.ServiceName, "users already exist")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		e.Logger.Infof("User [%s, %s] already exist", req.Email, req.Username)
		return errors.Conflict(e.ServiceName, "users already exist")
	}

	user := models.User{Username: req.Username, Email: req.Email, Password: string(hashed), IsAdmin: false, IsVerified: false}

	if err := e.Database.Create(&user).Error; err != nil {
		e.Logger.Fatalf("Internal error")
		return errors.InternalServerError(e.ServiceName, "users already exist")
	}

	e.Logger.Infof("New User created - [%d, %s]", user.ID, user.Username)

	rsp.User = UserToRpc(user)
	return nil
}

func (e *Users) Login(ctx context.Context, req *users.LoginRequest, rsp *users.UserResponse) error {
	e.Logger.Info("Received users.Login")

	user := models.User{}
	if notFound := e.Database.Where("username = ?", req.Username).First(&user).RecordNotFound(); notFound {
		e.Logger.Info("User not found")
		return errors.BadRequest(e.ServiceName, "username or password incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		e.Logger.Info("Password does not match")
		return errors.BadRequest(e.ServiceName, "username or password incorrect")
	}

	rsp.User = UserToRpc(user)
	return nil
}

func UserToRpc(user models.User) *users.User {
	return &users.User{
		Id:         int64(user.ID),
		Username:   user.Username,
		Email:      user.Email,
		Password:   user.Password,
		IsAdmin:    user.IsAdmin,
		IsVerified: user.IsVerified,
	}
}
