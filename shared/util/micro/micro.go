package micro

import (
    "github.com/micro/go-micro/v3"
    gc "github.com/micro/go-micro/v3/client/grpc"
    gs "github.com/micro/go-micro/v3/server/grpc"
    "github.com/rs/zerolog/log"

    "github.com/ygpark2/mboard/shared/config"
)


func WithTLS() micro.Option {
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
    return func(o *micro.Options) {} // noops
}

