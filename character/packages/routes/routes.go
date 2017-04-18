package routes

import (
	"github.com/pgarneau/microgame/character/packages/handlers"
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
		"/characters",
		handlers.ListCharacters,
	},
	Route{
		"Player",
		"GET",
		"/characters/{name}",
		handlers.GetCharacter,
	},
}
