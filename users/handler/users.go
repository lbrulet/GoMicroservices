package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	users "github.com/lbrulet/GoMicroservices/users/proto/users"
)

type Users struct{}

func (e *Users) Create(ctx context.Context, req *users.CreateRequest, rsp *users.User) error {
	rsp.Username = req.Username
	rsp.Password = req.Password
	rsp.Email = req.Email
	rsp.Id = 1
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Users) Call(ctx context.Context, req *users.Request, rsp *users.Response) error {
	log.Log("Received Users.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Users) Stream(ctx context.Context, req *users.StreamingRequest, stream users.Users_StreamStream) error {
	log.Logf("Received Users.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&users.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Users) PingPong(ctx context.Context, stream users.Users_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&users.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
