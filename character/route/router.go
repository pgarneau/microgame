package route

import (
	"github.com/gorilla/mux"
	"github.com/pgarneau/microgame/character/route/middleware/logger"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range AllRoutes {
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
