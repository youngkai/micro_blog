package main

import (
	"blog/posts/handler"
	posts "blog/posts/proto/posts"
	"blog/posts/subscriber"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.posts"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	posts.RegisterPostsHandler(service.Server(), &handler.Posts{
		Store:  service.Options().Store,
		Client: service.Client(),
	})

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.posts", service.Server(), new(subscriber.Post))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
