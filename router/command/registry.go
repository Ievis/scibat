package router

import (
	"scibat/src/competition/application/command/close"
	"scibat/src/competition/application/command/create"
)

func (r *Router) register() {
	r.add("/foo", close.New())
	r.add("/bar", create.New())
}
