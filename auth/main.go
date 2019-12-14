package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/lbrulet/GoMicroservices/auth/handler"
	"github.com/lbrulet/GoMicroservices/auth/subscriber"

	auth "github.com/lbrulet/GoMicroservices/auth/proto/auth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	_ = auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Register Struct as Subscriber
	_ = micro.RegisterSubscriber("go.micro.srv.auth", service.Server(), new(subscriber.Auth))

	// Register Function as Subscriber
	_ = micro.RegisterSubscriber("go.micro.srv.auth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
