package main

import (
	"github.com/lbrulet/GoMicroservices/users/database/postgreSQL"
	"github.com/lbrulet/GoMicroservices/users/handler"
	users "github.com/lbrulet/GoMicroservices/users/proto/users"
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
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

	db, err := postgreSQL.PostgresConnexion()
	if err != nil {
		logger.Fatal(err)
		return
	}
	repository := postgreSQL.NewPostgresRepository(db)

	// Register Handler
	if err := users.RegisterUsersHandler(service.Server(), &handler.Users{
		Logger:   logger,
		Database: repository,
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
