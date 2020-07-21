package main

import (
	"blog/comment/handler"
	comment "blog/comment/proto/comment"
	"blog/comment/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.comment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	comment.RegisterCommentHandler(service.Server(), new(handler.Comment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.comment", service.Server(), new(subscriber.Comment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
