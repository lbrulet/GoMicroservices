package main

import (
	"github.com/lbrulet/GoMicroservices/users/database"
	"github.com/lbrulet/GoMicroservices/users/handler"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"

	users "github.com/lbrulet/GoMicroservices/users/proto/users"
)

const (
	SERVICENAME = "go.micro.srv.users"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.users"),
		micro.Version("0.1"),
	)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{
		"micro-service": "users",
	})
	// Initialise service
	service.Init()

	dbService, err := database.ConnectDatabase()
	if err != nil {
		logger.Fatal(err)
		return
	}
	// Register Handler
	if err := users.RegisterUsersHandler(service.Server(), &handler.Users{
		Logger:   logger,
		Database: dbService,
		ServiceName: SERVICENAME,
	}); err != nil {
		logger.Fatal(err)
		return
	}

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
		return
	}
}
