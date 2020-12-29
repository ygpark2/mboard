package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	// tags "github.com/micro/services/blog/tags/proto"
)

func main() {
	event := InitializeEvent()
	event.Start()

	// Create the service
	srv := service.New(
		service.Name("boards"),
	)

	// Register Handler
	srv.Handle(new(handler.boardHandler))
	/*
		srv.Handle(&handler.Posts{
			Tags: tags.NewTagsService("tags", srv.Client()),
		})
	*/

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
