package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	// tags "github.com/micro/services/blog/tags/proto"

	"github.com/ygpark2/mboard/shared/config"
	"github.com/ygpark2/mboard/shared/constants"

	"github.com/ygpark2/mboard/service/account/registry"
	"github.com/ygpark2/mboard/service/board/handler"
)

func main() {

	cfg := config.GetConfig()

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)

	// Create the service
	srv := service.New(
		service.Name(constants.BOARD_SERVICE),
		service.Version(config.Version),

		// Adding some optional lifecycle actions
		service.BeforeStart(func() (err error) {
			logger.Debug("called BeforeStart")
			return
		}),
		service.BeforeStop(func() (err error) {
			logger.Debug("called BeforeStop")
			return
		}),
	)

	srv.Init(
	// micro.WrapHandler(handlerWrappers...),
	// micro.WrapSubscriber(subscriberWrappers...),
	)

	// Publisher publish to "mkit.service.emailer"
	publisher := service.NewEvent(constants.EMAILER_SERVICE)
	// greeterSrv Client to call "mkit.service.greeter"
	greeterSrvClient := greeterPB.NewGreeterService(constants.GREETER_SERVICE, srv.Client())

	// Register Handler
	srv.Handle(new(handler.NewBoardHandler))
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
