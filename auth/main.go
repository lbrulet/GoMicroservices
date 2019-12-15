package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/lbrulet/GoMicroservices/auth/handler"
	auth "github.com/lbrulet/GoMicroservices/auth/proto/auth"
	users "github.com/lbrulet/GoMicroservices/users/proto/users"
	"github.com/micro/go-micro"
)

func main() {
	// New Service
	// Create service
	service := micro.NewService(
		micro.Name("go.micro.api.api"),
		micro.Version("0.1"),
	)

	// Init to parse flags
	service.Init()

	// Register Handlers
	auth.RegisterAuthHandler(service.Server(), &handler.Auth{
		UsersService: users.NewUsersService("users", service.Client()),
	})

	// for handler use

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
