package main

import (
	"github.com/sirupsen/logrus"
	"github.com/lbrulet/GoMicroservices/users-service/proto/users"
	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"github.com/lbrulet/GoMicroservices/users-gateway/handler"

	users "github.com/lbrulet/GoMicroservices/users-gateway/proto/users"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.users"),
		micro.Version("0.1"),
	)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{
		"micro-service": "users-gateway",
	})

	// Initialise service
	service.Init()

	// Register Handler
	_ = users.RegisterUsersHandler(service.Server(), &handler.Users{
		UsersService: go_micro_srv_users.NewUsersService("go.micro.srv.users-service", service.Client()),
		Logger: logger,
	})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
