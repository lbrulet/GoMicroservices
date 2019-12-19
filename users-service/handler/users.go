package handler

import (
	"context"
	"github.com/lbrulet/GoMicroservices/users-service/database"
	"github.com/lbrulet/GoMicroservices/users-service/models"
	"github.com/micro/go-micro/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	users "github.com/lbrulet/GoMicroservices/users-service/proto/users"
)

// Users is used to create a new RPC service
type Users struct {
	ServiceName string
	Logger      *logrus.Entry
	Database    database.Repository
}

// Create is called as service by an API and will create a user
func (e *Users) Create(ctx context.Context, req *users.CreateRequest, rsp *users.UserResponse) error {
	e.Logger.Info("Received users-service.Create")

	count, err := e.Database.CountByUsernameAndEmail(req.Username, req.Email)
	if err != nil {
		return errors.InternalServerError(e.ServiceName, ERROR_UNEXPECTED)
	}

	if count != 0 {
		e.Logger.Infof("User [%s, %s] already exist", req.Email, req.Username)
		return errors.Conflict(e.ServiceName, ERROR_USER_ALREADY_EXIST)
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		e.Logger.Infof("User [%s, %s] already exist", req.Email, req.Username)
		return errors.Conflict(e.ServiceName, ERROR_USER_ALREADY_EXIST)
	}

	user := models.User{Username: req.Username, Email: req.Email, Password: string(hashed), IsAdmin: false, IsVerified: false}

	if err := e.Database.CreateUser(user); err != nil {
		e.Logger.Fatalf(ERROR_UNEXPECTED)
		return errors.InternalServerError(e.ServiceName, ERROR_UNEXPECTED)
	}

	e.Logger.Infof("New User created - [%d, %s]", user.ID, user.Username)

	rsp.User = UserToRpc(user)
	return nil
}

// Login is called as service by an API and will log the user and generate credentials
func (e *Users) Login(ctx context.Context, req *users.LoginRequest, rsp *users.UserResponse) error {
	e.Logger.Info("Received users-service.Login")

	user, err := e.Database.GetByUsername(req.Username)
	if err != nil {
		e.Logger.Info("User not found")
		return errors.BadRequest(e.ServiceName, ERROR_USER_USERNAME_OR_PASSWORD_INVALID)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		e.Logger.Info("Password does not match")
		return errors.BadRequest(e.ServiceName, ERROR_USER_USERNAME_OR_PASSWORD_INVALID)
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
