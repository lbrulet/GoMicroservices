package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"

	"github.com/lbrulet/GoMicroservices/users-gateway/client"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/go-micro/api/proto"
	users "path/to/service/proto/users"
)

type Users struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Users.Call is called by the API as /users/call with post body {"name": "foo"}
func (e *Users) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Users.Call request")

	// extract the client from the context
	usersClient, ok := client.UsersFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.users.users.call", "users client not found")
	}

	// make request
	response, err := usersClient.Call(ctx, &users.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.users.users.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
