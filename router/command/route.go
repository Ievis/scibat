package router

import "scibat/src/common/domain/event"

type Route struct {
	Path    string
	Command event.Command
}

func New(path string, command event.Command) *Route {
	return &Route{
		Path:    path,
		Command: command,
	}
}
