package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/ygpark2/mboard/services/content/handler"
)

func main() {
	// Create the service
	srv := service.New(
		service.Name("contents"),
	)

	// Register Handler
	srv.Handle(new(handler.Contents))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
