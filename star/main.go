package main

import (
	"blog/star/handler"
	star "blog/star/proto/star"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.star"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	star.RegisterStarHandler(service.Server(), new(handler.Star))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
