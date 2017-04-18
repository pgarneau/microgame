package router

import (
	"github.com/gorilla/mux"
	"github.com/pgarneau/microgame/character/packages/logger"
	"github.com/pgarneau/microgame/character/packages/routes"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.AllRoutes {
		var handler http.Handler
		handler = route.HandleFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
