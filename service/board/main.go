package main

import (
	"time"

	"github.com/google/uuid"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/logger"

	// tags "github.com/micro/services/blog/tags/proto"

	"github.com/ygpark2/mboard/shared/config"
	"github.com/ygpark2/mboard/shared/constants"

	"github.com/ygpark2/mboard/service/board/registry"
)

func main() {

	logger.Debug("++++++++++++++++++++++ start NewContainer ++++++++++++++++++++++++++++++")
	// Initialize DI Container
	ctn := registry.NewContainer()

	logger.Debug("++++++++++++++++++++++ start auth start ++++++++++++++++++++++++++++++")
	// setupAuthForService("admin", "micro")
	logger.Debug("++++++++++++++++++++++ end auth start ++++++++++++++++++++++++++++++")

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
	/*
		micro.WrapHandler(handlerWrappers...),
		micro.WrapSubscriber(subscriberWrappers...),
	*/
	)

	// Publisher publish to "mkit.service.emailer"
	// publisher := service.NewEvent(constants.EMAILER_SERVICE)
	// greeterSrv Client to call "mkit.service.greeter"
	// greeterSrvClient := greeterPB.NewGreeterService(constants.GREETER_SERVICE, srv.Client())

	// Register Handler
	srv.Handle(ctn.BoardHandler)
	/*
		srv.Handle(&handler.Posts{
			Tags: tags.NewTagsService("tags", srv.Client()),
		})
	*/

	// srv.Options()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

// setupAuthForService generates auth credentials for the service
func setupAuthForService(accID string, accSecret string) error {
	// opts := auth.DefaultAuth.Options()

	// extract the account creds from options, these can be set by flags
	// accID := ID
	// accSecret := Secret

	// if no credentials were provided, self generate an account
	if len(accID) == 0 || len(accSecret) == 0 {
		opts := []auth.GenerateOption{
			auth.WithType("service"),
			auth.WithScopes("service"),
		}

		acc, err := auth.Generate(uuid.New().String(), opts...)
		if err != nil {
			return err
		}
		if logger.V(logger.DebugLevel, logger.DefaultLogger) {
			logger.Debugf("Auth [%v] Generated an auth account", auth.DefaultAuth.String())
		}

		accID = acc.ID
		accSecret = acc.Secret
	}

	// generate the first token
	token, err := auth.Token(
		auth.WithCredentials(accID, accSecret),
		auth.WithExpiry(time.Minute*10),
	)
	if err != nil {
		return err
	}

	// set the credentials and token in auth options
	auth.DefaultAuth.Init(
		auth.ClientToken(token),
		auth.Credentials(accID, accSecret),
	)
	return nil
}
