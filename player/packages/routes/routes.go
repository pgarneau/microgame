package routes

import (
	"github.com/gorilla/mux"
	"github.com/pgarneau/microgame/player/packages/handlers"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"Players",
		"GET",
		"/players",
		handlers.GetPlayers,
	},
	Route{
		"Player",
		"GET",
		"/players/{player}",
		handlers.GetPlayer,
	},
}
