package handler

import (
	"context"
	auth "github.com/lbrulet/GoMicroservices/auth/proto/auth"
	"github.com/lbrulet/GoMicroservices/users/proto/users"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
)

type Auth struct{
	UsersService go_micro_srv_users.UsersService
}

// Auth.Login is called by the API as /auth/login with post body {"username": "sankamille", "password": "password123"}
func (e *Auth) Login(ctx context.Context, req *auth.LoginRequest, rsp *auth.Response) error {
	log.Log("Received auth.Call API request")

	if len(req.Username) == 0 {
		return errors.BadRequest("go.micro.api.auth", "no content for username")
	}

	rsp.Data = "got your request " + req.Username
	rsp.Success = true
	return nil
}

// Auth.Register is called by the API as /auth/register with post body {"username": "sankamille", "email": "", "password: "password123"}
func (e *Auth) Register(ctx context.Context, req *auth.RegisterRequest, rsp *auth.Response) error {
	log.Log("Received auth.Call API request")

	if len(req.Username) == 0 || len(req.Password) == 0 || len(req.Email) == 0 {
		return errors.BadRequest("go.micro.api.auth", "content missing")
	}

	user, err := e.UsersService.Create(context.TODO(), &go_micro_srv_users.CreateRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		return errors.BadRequest("go.micro.api.auth", err.Error())
	}

	rsp.Data = "New users registered as " + user.Email + " " + user.Username
	rsp.Success = true
	return nil
}
