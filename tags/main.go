package main

import (
	"blog/tags/handler"
	tags "blog/tags/proto/tags"
	"blog/tags/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Serviceost
	service := micro.NewService(
		micro.Name("go.micro.service.tags"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	tags.RegisterTagsHandler(service.Server(), &handler.Tags{
		Store: service.Options().Store,
	})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.tags", service.Server(), new(subscriber.Tags))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
