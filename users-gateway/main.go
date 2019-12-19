package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"github.com/lbrulet/GoMicroservices/users-gateway/handler"
	"github.com/lbrulet/GoMicroservices/users-gateway/client"

	users "github.com/lbrulet/GoMicroservices/users-gateway/proto/users"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.users"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Users srv client
		micro.WrapHandler(client.UsersWrapper(service)),
	)

	// Register Handler
	users.RegisterUsersHandler(service.Server(), new(handler.Users))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
