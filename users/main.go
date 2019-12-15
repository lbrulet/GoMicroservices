package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/lbrulet/GoMicroservices/users/handler"

	users "github.com/lbrulet/GoMicroservices/users/proto/users"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.users"),
		micro.Version("0.1"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	users.RegisterUsersHandler(service.Server(), new(handler.Users))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
