package router

import (
	"github.com/gorilla/mux"
	commandRouter "scibat/router/command"
	queryRouter "scibat/router/query"
)

func Init() *mux.Router {
	var r *mux.Router
	r = commandRouter.Init()
	r = queryRouter.Init()
	return r
}
