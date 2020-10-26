package micro

import (
	"github.com/micro/go-micro"
	gc "github.com/micro/micro/v3/service/client/grpc"
	"github.com/micro/micro/v3/service/options"
	gs "github.com/micro/micro/v3/service/server/grpc"
	"github.com/rs/zerolog/log"

	"github.com/ygpark2/mboard/shared/config"
)

func WithTLS() options.Option {
	if config.IsSecure() {
		if tlsConf, err := config.CreateServerCerts(); err != nil {
			log.Panic().Err(err).Msg("unable to load certs")
		} else {
			log.Info().Msg("TLS Enabled")
			return func(o *micro.Options) {
				o.Client.Init(
					gc.AuthTLS(tlsConf),
				)
				o.Server.Init(
					gs.AuthTLS(tlsConf),
				)
			}
		}
	}
	return func(o *options.Options) {} // noops
}
