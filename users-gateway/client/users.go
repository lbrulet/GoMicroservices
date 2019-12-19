package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	users "path/to/service/proto/users"
)

type usersKey struct {}

// FromContext retrieves the client from the Context
func UsersFromContext(ctx context.Context) (users.UsersService, bool) {
	c, ok := ctx.Value(usersKey{}).(users.UsersService)
	return c, ok
}

// Client returns a wrapper for the UsersClient
func UsersWrapper(service micro.Service) server.HandlerWrapper {
	client := users.NewUsersService("go.micro.srv.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, usersKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
