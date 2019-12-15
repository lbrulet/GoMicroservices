package handler

import (
	"context"
	"github.com/micro/go-micro/util/log"

	api "github.com/micro/go-micro/api/proto"
)

type Auth struct{}

// Auth.Call is called by the API as /auth/call with post body {"name": "foo"}
func (e *Auth) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Greeter.Hello API request")
	return nil
}
