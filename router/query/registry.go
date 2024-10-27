package router

import (
	"scibat/src/competition/application/query"
)

func (r *Router) register() {
	r.add("/foo", query.New())
}
