package handler

import (
	"context"
	"github.com/sirupsen/logrus"

	auth "github.com/lbrulet/GoMicroservices/auth/proto/auth"
	"github.com/lbrulet/GoMicroservices/users/proto/users"
	"github.com/micro/go-micro/errors"
)

type Auth struct {
	ServiceName  string
	UsersService go_micro_srv_users.UsersService
	Logger       *logrus.Entry
	JwtService   *JwtService
}

// Auth.Login is called by the API as /auth/login with post body {"username": "sankamille", "password": "password123"}
func (e *Auth) Login(ctx context.Context, req *auth.LoginRequest, rsp *auth.LoginResponse) error {
	e.Logger.Info("Received auth.Login")

	if len(req.Username) == 0 || len(req.Password) == 0 {
		e.Logger.Info("Content missing in payload")
		return errors.BadRequest(e.ServiceName, "content missing")
	}

	response, err := e.UsersService.Login(context.TODO(), &go_micro_srv_users.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		e.Logger.Info(err.Error())
		return errors.BadRequest(e.ServiceName, err.Error())
	}

	token, err := e.JwtService.CreateToken(int(response.User.Id))
	if err != nil {
		e.Logger.Info(err.Error())
		return errors.InternalServerError(e.ServiceName, "internal server error")
	}

	e.Logger.Infof("User logged as [%d, %s]", response.User.Id, response.User.Username)

	rsp.Token = token
	rsp.Data = response.User
	return nil
}

// Auth.Register is called by the API as /auth/register with post body {"username": "sankamille", "email": "", "password: "password123"}
func (e *Auth) Register(ctx context.Context, req *auth.RegisterRequest, rsp *auth.RegisterResponse) error {
	e.Logger.Info("Received auth.Register")

	if len(req.Username) == 0 || len(req.Password) == 0 || len(req.Email) == 0 {
		e.Logger.Info("Content missing in payload")
		return errors.BadRequest(e.ServiceName, "content missing")
	}

	response, err := e.UsersService.Create(context.TODO(), &go_micro_srv_users.CreateRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		e.Logger.Info(err.Error())
		return errors.BadRequest(e.ServiceName, err.Error())
	}

	e.Logger.Infof("Received from user service [%s]", response.String())
	rsp.Data = response.User
	return nil
}

func (e *Auth) VerifyToken(ctx context.Context, req *auth.TokenRequest, rsp *auth.ClaimsResponse) error {
	e.Logger.Info("Received auth.VerifyToken")

	id, err := e.JwtService.VerifyToken(req.Token)
	if err != nil {
		return errors.BadRequest(e.ServiceName, "invalid token")
	}
	rsp.Id = int32(id)
	return nil
}
