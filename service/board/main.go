package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/ygpark2/mboard/service/board/handler"
	tags "github.com/micro/services/blog/tags/proto"
)

func main() {
	// Create the service
	srv := service.New(
		service.Name("boards"),
	)

	// Register Handler
	srv.Handle(new(handler.BoardHander))
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
