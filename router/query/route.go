package router

import (
	"scibat/src/common/domain/query"
)

type Route struct {
	Path    string
	Command query.Command
}

func New(path string, command query.Command) *Route {
	return &Route{
		Path:    path,
		Command: command,
	}
}
