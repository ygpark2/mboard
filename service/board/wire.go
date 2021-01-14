package main

import (
	"github.com/google/wire"

	"github.com/ygpark2/mboard/service/board/repository"
)

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(repository.NewBoardRepository, NewGreeter, NewMessage)
	return Event{}, nil
}
