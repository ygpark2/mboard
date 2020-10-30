package main

import (
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/server"
	"github.com/rs/zerolog/log"

	"github.com/xmlking/micro-starter-kit/shared/config"
	"github.com/xmlking/micro-starter-kit/shared/constants"
	logWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/log"
	transWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/transaction"
	validatorWrapper "github.com/xmlking/micro-starter-kit/shared/wrapper/validator"
	"github.com/ygpark2/mboard/service/account/handler"
	"github.com/ygpark2/mboard/service/account/registry"
	"github.com/ygpark2/mboard/service/account/repository"
)

func main() {
	cfg := config.GetConfig()

	// Initialize Features
	var clientWrappers []client.Wrapper
	var handlerWrappers []server.HandlerWrapper
	var subscriberWrappers []server.SubscriberWrapper

	// Wrappers are invoked in the order as they added
	if cfg.Features.Reqlogs.Enabled {
		clientWrappers = append(clientWrappers, logWrapper.NewClientWrapper())
		handlerWrappers = append(handlerWrappers, logWrapper.NewHandlerWrapper())
		subscriberWrappers = append(subscriberWrappers, logWrapper.NewSubscriberWrapper())
	}
	//if cfg.Features.Translogs.Enabled {
	//    topic := cfg.Features.Translogs.Topic
	//    publisher := micro.NewEvent(topic, client.DefaultClient) // service.Client())
	//    handlerWrappers = append(handlerWrappers, transWrapper.NewHandlerWrapper(publisher))
	//    subscriberWrappers = append(subscriberWrappers, transWrapper.NewSubscriberWrapper(publisher))
	//}
	if cfg.Features.Validator.Enabled {
		handlerWrappers = append(handlerWrappers, validatorWrapper.NewHandlerWrapper())
		subscriberWrappers = append(subscriberWrappers, validatorWrapper.NewSubscriberWrapper())
	}

	srv := service.NewService(
		service.Name(constants.ACCOUNT_SERVICE),
		service.Version(config.Version),
		// myMicro.WithTLS(),
		// Wrappers are applied in reverse order so the last is executed first.
		service.WrapClient(clientWrappers...),
		// Adding some optional lifecycle actions
		service.BeforeStart(func() (err error) {
			log.Debug().Msg("called BeforeStart")
			return
		}),
		service.BeforeStop(func() (err error) {
			log.Debug().Msg("called BeforeStop")
			return
		}),
	)

	if cfg.Features.Translogs.Enabled {
		topic := cfg.Features.Translogs.Topic
		publisher := service.NewEvent(topic, service.Client())
		handlerWrappers = append(handlerWrappers, transWrapper.NewHandlerWrapper(publisher))
		subscriberWrappers = append(subscriberWrappers, transWrapper.NewSubscriberWrapper(publisher))
	}

	service.Init(
		service.WrapHandler(handlerWrappers...),
		service.WrapSubscriber(subscriberWrappers...),
	)

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}

	// // Handlers
	userHandler := handler.NewUserHandler(ctn.Resolve("user-repository").(repository.UserRepository), publisher, greeterSrvClient)
	profileHandler := ctn.Resolve("profile-handler").(profilePB.ProfileServiceHandler)

	// Register Handlers
	userPB.RegisterUserServiceHandler(service.Server(), userHandler)
	profilePB.RegisterProfileServiceHandler(service.Server(), profileHandler)

	println(config.GetBuildInfo())

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
