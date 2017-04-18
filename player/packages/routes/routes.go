package routes

import (
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

var AllRoutes = Routes{
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
