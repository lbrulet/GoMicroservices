package handler

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/lbrulet/GoMicroservices/users-service/proto/users"
	"github.com/micro/go-micro/util/log"

	api "github.com/micro/go-micro/api/proto"
)

type Users struct{
	UsersService go_micro_srv_users.UsersService
	Logger *logrus.Entry
}

// Users.Call is called by the API as /users/call with post body {"name": "foo"}
func (e *Users) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Users.Call request")

	rsp.StatusCode = 200
	rsp.Body = string("JE PARLE CHINOIS")

	return nil
}
