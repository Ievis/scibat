package router

import (
	"github.com/gorilla/mux"
	queryInterface "scibat/src/common/domain/query"
)

type Router struct {
	RouteList []*Route
}

func Init() *mux.Router {
	var router Router
	router.register()
	r := mux.NewRouter()
	for i := 0; i < len(router.RouteList); i++ {
		r.HandleFunc(router.RouteList[i].Path, handler(router.RouteList[i].Command))
	}
	return r
}

func (r *Router) add(path string, command queryInterface.Command) {
	r.RouteList = append(r.RouteList, New(path, command))
}
