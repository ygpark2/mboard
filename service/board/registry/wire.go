package registry

import (
	"github.com/google/wire"

	configPB "github.com/ygpark2/mboard/shared/proto/config"
)

// Container - provide di Container
type Container struct {
	ctn wire.ProviderSet
}

// NewContainer - create new Container
func NewContainer(cfg configPB.Configuration) (*Container, error) {
	providerSet := wire.NewSet()

	providerSet.add()

	wire.NewSet()

	panic(wire.Build(
		wire.Struct(new(Logger), "*"),
		container.buildBoardRepository,
		NewHttpClient,
		NewConcatService,
	))
	/*
		return &Container{
			ctn: wire.Build(providerSet),
		}, nil
	*/
}
