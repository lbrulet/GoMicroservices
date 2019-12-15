package main

import (
	"github.com/sirupsen/logrus"

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

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{
		"micro-service": "auth",
	})
	// Init to parse flags
	service.Init()

	// Register Handlers
	_ = auth.RegisterAuthHandler(service.Server(), &handler.Auth{
		UsersService: users.NewUsersService("go.micro.srv.users", service.Client()),
		Logger:       logger,
	})

	// for handler use

	// Run server
	if err := service.Run(); err != nil {
		logger.Fatal()
	}
}
