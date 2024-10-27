package cqrs

import (
	"scibat/src/shared/domain/event"
)

type Adapter struct {
	//
}

func (a Adapter) InitCommand() event.Command {
	//TODO implement me
	panic("implement me")
}

func New() *Adapter {
	return &Adapter{
		//
	}
}
